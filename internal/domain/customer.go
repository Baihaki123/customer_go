package domain

import "time"

type Customer struct {
	ID            int          `json:"cst_id" gorm:"primaryKey;column:cst_id;autoIncrement"`
	NationalityID int          `json:"nationality_id" gorm:"column:nationality_id"`
	Name          string       `json:"cst_name" gorm:"column:cst_name;size:50;not null"`
	DOB           time.Time    `json:"cst_dob" gorm:"column:cst_dob;type:date"`
	PhoneNum      string       `json:"cst_phoneNum" gorm:"column:cst_phoneNum;size:20"`
	Email         string       `json:"cst_email" gorm:"column:cst_email;size:50"`
	Nationality   *Nationality `json:"nationality,omitempty" gorm:"foreignKey:NationalityID;references:ID"`
	Families      []FamilyList `json:"family_list,omitempty" gorm:"foreignKey:CustomerID;references:ID"`
}
