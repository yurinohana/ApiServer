package models

type User struct {
	ID   int    `dynamo:"id"`
	Name string `dynamo:"name"`
}
