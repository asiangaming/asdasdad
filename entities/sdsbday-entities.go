package entities

type Model_sdsbday struct {
	Sdsbday_id     int    `json:"sdsbday_id"`
	Sdsbday_date   string `json:"sdsbday_date"`
	Sdsbday_prize1 string `json:"sdsbday_prize1"`
	Sdsbday_prize2 string `json:"sdsbday_prize2"`
	Sdsbday_prize3 string `json:"sdsbday_prize3"`
	Sdsbday_create string `json:"sdsbday_create"`
	Sdsbday_update string `json:"sdsbday_update"`
}
type Controller_sdsbdaysave struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tanggal  string `json:"tanggal" validate:"required"`
}
type Controller_sdsbdayprize1save struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tipe     string `json:"tipe" validate:"required"`
	Prize    string `json:"prize" validate:"required"`
}

type Responseredis_sdsbday struct {
	Sdsbday_id     int    `json:"sdsbday_id"`
	Sdsbday_date   string `json:"sdsbday_date"`
	Sdsbday_prize1 string `json:"sdsbday_prize1"`
	Sdsbday_prize2 string `json:"sdsbday_prize2"`
	Sdsbday_prize3 string `json:"sdsbday_prize3"`
	Sdsbday_create string `json:"sdsbday_create"`
	Sdsbday_update string `json:"sdsbday_update"`
}
