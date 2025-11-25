package domain

import "time"

type FamilyList struct {
	ID         int       `json:"fl_id" gorm:"primaryKey;column:fl_id;autoIncrement"`
	CustomerID int       `json:"cst_id" gorm:"column:cst_id"`
	Relation   string    `json:"fl_relation" gorm:"column:fl_relation;size:50"`
	Name       string    `json:"fl_name" gorm:"column:fl_name;size:50"`
	DOB        time.Time `json:"fl_dob" gorm:"column:fl_dob;type:date"`
}
