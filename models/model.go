package models

type Car struct {
	ID           int    ` json:"id"`
	CustomerName string `json:"customer"`
	CarName      string `json:"name"`
	Status       string ` json:"status"`
}
