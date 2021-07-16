package models

type Demo struct {  //表对应的结构体
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
