package model

type Todo struct {
	Id     int    `form:"id"`
	Title  string `form:"title"`
	Status string `form:"status"`
}
