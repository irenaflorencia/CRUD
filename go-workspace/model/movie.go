package model

import "gorm.io/gorm"

type Movie struct{
	gorm.Model
	Title 		string 'gorm:"size:255;not null"'
	Slug 		string 'gorm:"size:255;unique:true;not null"'
	Description	string 'gorm:"not null"'
	Duration 	string 'gorm:"size:5;not null"'
	Image 		string 'gorm:"size:255;not null"'
}


