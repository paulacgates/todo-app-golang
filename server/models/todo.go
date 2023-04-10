package models

type Todo struct {
	Id    int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

var Todos = []Todo{}
