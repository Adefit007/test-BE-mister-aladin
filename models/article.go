package models

import "time"

type Article struct {
	ID        	int    		`json:"id"`
	Author    	string 		`json:"author" gorm:"type: varchar(255)"`
	Title     	string 		`json:"title" gorm:"type: varchar(255)"`
	Body     	string 		`json:"body" gorm:"type: varchar(255)"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}