package domain

type Nationality struct {
	ID   int    `json:"nationality_id" gorm:"primaryKey;column:nationality_id;autoIncrement"`
	Name string `json:"nationality_name" gorm:"column:nationality_name;size:50"`
	Code string `json:"nationality_code" gorm:"column:nationality_code;size:2"`
}
