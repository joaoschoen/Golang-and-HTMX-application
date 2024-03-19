package model

// TYPES

type User struct {
	Id    string `form:"id"`
	Name  string `form:"name"`
	Value string `form:"value"`
}
