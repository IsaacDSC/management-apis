package model

import "gorm.io/gorm"

type API struct {
	gorm.Model
	Name   string
	Method string
	Url    string
	Header string
	Body   string
}

// Name (TEXT)
//Method (TEXT)
//Url (TEXT)
//Header (TEXT com quebras de linhas /n sendo um novo header)
// Body (TEXT json em string)
