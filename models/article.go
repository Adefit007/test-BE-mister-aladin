package models

import "time"

type Article struct {
	ID        	int    		`json:"id"`
	Author    	string 		`json:"author" gorm:"type: text"`
	Title     	string 		`json:"title" gorm:"type: text"`
	Body     	string 		`json:"body" gorm:"type: text"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}