package model

type Executor struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"Name"`
	SecondName string `json:"SecondName"`
	Surname string `json:"Surname,omitempty"`
	TypeDesign string `json:"TypeDesign"`
	Tax float64 `json:"Tax"`
	PayTime string `json:"payTime"`
	Cost float64 `json:"Cost"`
}
type RequsetExecutor struct{
	Name string `json:"Name"`
	SecondName string `json:"SecondName"`
	Surname string `json:"Surname"`
	TypeDesign string `json:"TypeDesign"`
	Tax float64 `json:"Tax"`
	PayTime string `json:"payTime"`
	Cost float64 `json:"Cost"`
}