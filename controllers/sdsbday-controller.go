package controllers

import (
	"log"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/entities"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/helpers"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Sdsbdayhome(c *fiber.Ctx) error {
	field_redis := "LISTSDSBDAY_SDSB4D"

	var obj entities.Responseredis_sdsbday
	var arraobj []entities.Responseredis_sdsbday
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		sdsbday_id, _ := jsonparser.GetInt(value, "sdsbday_id")
		sdsbday_date, _ := jsonparser.GetString(value, "sdsbday_date")
		sdsbday_prize1, _ := jsonparser.GetString(value, "sdsbday_prize1")
		sdsbday_prize2, _ := jsonparser.GetString(value, "sdsbday_prize2")
		sdsbday_prize3, _ := jsonparser.GetString(value, "sdsbday_prize3")
		sdsbday_create, _ := jsonparser.GetString(value, "sdsbday_create")
		sdsbday_update, _ := jsonparser.GetString(value, "sdsbday_update")

		obj.Sdsbday_id = int(sdsbday_id)
		obj.Sdsbday_date = sdsbday_date
		obj.Sdsbday_prize1 = sdsbday_prize1
		obj.Sdsbday_prize2 = sdsbday_prize2
		obj.Sdsbday_prize3 = sdsbday_prize3
		obj.Sdsbday_create = sdsbday_create
		obj.Sdsbday_update = sdsbday_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_sdsbdayHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 0)
		log.Println("SDSBDAY MYSQL")
		return c.JSON(result)
	} else {
		log.Println("SDSBDAY CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func SdsbdaySave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_sdsbdaysave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_sdsbdayHome(
		client_admin,
		client.Tanggal, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSDSBDAY_SDSB4D"
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete BACKEND LISTSDSBDAY_SDSB4D : %d", val_master)
	field_redis_api := "SDSB4D_LISTSDSBDAY_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBDAY_SDSB4D : %d", val_api)
	return c.JSON(result)
}
func SdsbdayGeneratorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_sdsbdayprize1save)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	field := ""
	switch client.Tipe {
	case "prize1":
		field = "prize1_sdsb4dday"
	case "prize2":
		field = "prize2_sdsb4dday"
	case "prize3":
		field = "prize3_sdsb4dday"
	}

	result, err := models.Save_sdsbdayGenerator(
		client_admin,
		field, client.Prize, client.Sdata, client.Idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSDSBDAY_SDSB4D"
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete BACKEND LISTSDSBDAY_SDSB4D : %d", val_master)
	field_redis_api := "SDSB4D_LISTSDSBDAY_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBDAY_SDSB4D : %d", val_api)
	return c.JSON(result)
}
func SdsbdayGeneratorNumber(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_Generator(client_admin)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSDSBDAY_SDSB4D"
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete BACKEND LISTSDSBDAY_SDSB4D : %d", val_master)
	field_redis_api := "SDSB4D_LISTSDSBDAY_API"
	val_api := helpers.DeleteRedis(field_redis_api)
	log.Printf("Redis Delete API LISTSDSBDAY_SDSB4D : %d", val_api)
	return c.JSON(result)
}
