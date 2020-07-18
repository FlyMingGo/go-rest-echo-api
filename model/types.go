package model

type UserInfo struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
	City string `json:"city"`
}
