package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string `json:"title" from:"title"`
	Author        string `json:"author" from:"author"`
	PublisherYear int    `json:"publisherYear" from:"publisherYear"`
	TotalPage     int    `json:"totalPage" from:"totalPage"`
}
