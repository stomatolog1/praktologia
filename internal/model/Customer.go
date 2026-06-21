package model

type Customer struct{
	ID uint `json:"id" gorm:"primaryKey"`
	InnCustomer int `json:"InnCustomer"`
	TypeCustomer string `json:"TypeCustomer"`
	Firm string `json:"Firm"`
	Name string `json:"Name"`
	SecondName string `json:"SecondName"`
	Surname string `json:"Surname,omitempty"`
	Email string `json:"Email"`
	Phone int `json:"Phone"`
}
type RequestCustomer struct{
	InnCustomer int `json:"InnCustomer"`
	TypeCustomer string `json:"TypeCustomer"`
	Firm string `json:"Firm"`
	Name string `json:"Name"`
	SecondName string `json:"SecondName"`
	Surname string `json:"Surname"`
	Email string `json:"Email"`
	Phone int `json:"Phone"`
}