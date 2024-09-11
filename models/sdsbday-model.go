package models

import (
	"context"
	"log"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/configs"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

func Fetch_sdsbdayHome() (helpers.Response, error) {
	var obj entities.Model_sdsbday
	var arraobj []entities.Model_sdsbday
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idsdsb4dday , datesdsb4dday, 
			prize1_sdsb4dday , prize2_sdsb4dday, prize3_sdsb4dday, 
			create_sdsb4dday, COALESCE(createdate_sdsb4dday,now()), update_sdsb4dday, COALESCE(updatedate_sdsb4dday,now())  
			FROM ` + configs.DB_tbl_trx_sdsb4d_day + ` 
			ORDER BY datesdsb4dday DESC LIMIT 365 
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idsdsb4dday_db                                                                             int
			datesdsb4dday_db, prize1_sdsb4dday_db, prize2_sdsb4dday_db, prize3_sdsb4dday_db            string
			create_sdsb4dday_db, createdate_sdsb4dday_db, update_sdsb4dday_db, updatedate_sdsb4dday_db string
		)

		err = row.Scan(
			&idsdsb4dday_db, &datesdsb4dday_db, &prize1_sdsb4dday_db, &prize2_sdsb4dday_db, &prize3_sdsb4dday_db,
			&create_sdsb4dday_db, &createdate_sdsb4dday_db, &update_sdsb4dday_db, &updatedate_sdsb4dday_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		date_day, _ := goment.New(datesdsb4dday_db)
		create_day, _ := goment.New(createdate_sdsb4dday_db)
		update_day, _ := goment.New(updatedate_sdsb4dday_db)
		if create_sdsb4dday_db != "" {
			create = create_sdsb4dday_db + ", " + create_day.Format("YYYY-MM-DD HH:mm:ss")
		}
		if update_sdsb4dday_db != "" {
			update = update_sdsb4dday_db + ", " + update_day.Format("YYYY-MM-DD HH:mm:ss")
		}
		obj.Sdsbday_id = idsdsb4dday_db
		obj.Sdsbday_date = date_day.Format("YYYY-MM-DD")
		obj.Sdsbday_prize1 = prize1_sdsb4dday_db
		obj.Sdsbday_prize2 = prize2_sdsb4dday_db
		obj.Sdsbday_prize3 = prize3_sdsb4dday_db
		obj.Sdsbday_create = create
		obj.Sdsbday_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_sdsbdayHome(admin, tanggal, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(configs.DB_tbl_trx_sdsb4d_day, "datesdsb4dday", tanggal)
		if !flag {
			sql_insert := `
				insert into
				` + configs.DB_tbl_trx_sdsb4d_day + ` (
					idsdsb4dday , datesdsb4dday, create_sdsb4dday, createdate_sdsb4dday
				) values (
					$1 ,$2, $3, $4
				)
			`
			stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
			helpers.ErrorCheck(e_insert)
			defer stmt_insert.Close()
			field_column := configs.DB_tbl_trx_sdsb4d_day + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			res_newrecord, e_newrecord := stmt_insert.ExecContext(
				ctx,
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
				tanggal,
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			helpers.ErrorCheck(e_newrecord)
			insert, e := res_newrecord.RowsAffected()
			helpers.ErrorCheck(e)
			if insert > 0 {
				flag = true
				msg = "Succes"
				log.Println("Data Berhasil di save")
			}
		} else {
			msg = "Duplicate Entry"
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func Save_sdsbdayGenerator(admin, field, prize, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "Edit" {
		sql_update := `
				UPDATE 
				` + configs.DB_tbl_trx_sdsb4d_day + `  
				SET ` + field + ` =$1,  
				update_sdsb4dday=$2, updatedate_sdsb4dday=$3
				WHERE idsdsb4dday=$4 
			`
		stmt_record, e := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(e)
		rec_record, e_record := stmt_record.ExecContext(
			ctx,
			prize,
			admin,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			idrecord)
		helpers.ErrorCheck(e_record)
		update_record, e_record := rec_record.RowsAffected()
		helpers.ErrorCheck(e_record)

		defer stmt_record.Close()
		if update_record > 0 {
			flag = true
			msg = "Succes"
			log.Printf("Update SDSB4D-DAY Success : %d\n", idrecord)
		} else {
			log.Println("Update SDSB4D-DAY failed")
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
func Save_Generator(admin string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	for i := 0; i <= 299; i++ {
		tglnow2, _ := goment.New()
		tanggal := tglnow2.Add(i, "days").Format("YYYY-MM-DD")
		flag = CheckDB(configs.DB_tbl_trx_sdsb4d_day, "datesdsb4dday", tanggal)
		if !flag {
			sql_insert := `
			insert into
			` + configs.DB_tbl_trx_sdsb4d_day + ` (
				idsdsb4dday , datesdsb4dday, prize1_sdsb4dday, prize2_sdsb4dday, prize3_sdsb4dday, 
				create_sdsb4dday, createdate_sdsb4dday
			) values (
				$1 ,$2, $3, $4, $5,
				$6, $7
			)
		`
			stmt_insert, e_insert := con.PrepareContext(ctx, sql_insert)
			helpers.ErrorCheck(e_insert)
			defer stmt_insert.Close()

			prize_1 := helpers.GenerateNumber(4)
			prize_2 := helpers.GenerateNumber(4)
			prize_3 := helpers.GenerateNumber(4)
			field_column := configs.DB_tbl_trx_sdsb4d_day + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			res_newrecord, e_newrecord := stmt_insert.ExecContext(
				ctx,
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
				tanggal,
				prize_1,
				prize_2,
				prize_3,
				admin,
				tglnow.Format("YYYY-MM-DD HH:mm:ss"))
			helpers.ErrorCheck(e_newrecord)
			insert, e := res_newrecord.RowsAffected()
			helpers.ErrorCheck(e)
			if insert > 0 {
				flag = true
				msg = "Succes"
				log.Println("Data Berhasil di save")
			}
		} else {
			break
		}
	}

	if flag {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = msg
		res.Record = nil
		res.Time = time.Since(render_page).String()
	}

	return res, nil
}
