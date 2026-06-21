package model

type Sotrudnik struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"Name"`
	SecondName string `json:"SecondName"`
	Surname string `json:"Surname,omitempty"`
	Position string `json:"Position"`
	PayMount float64 `json:"PayMount"`
	Tax float64 `json:"Tax"`
}

type RequestSotrudnik struct{
	Name string `json:"Name"`
	SecondName string `json:"SecondName"`
	Surname string `json:"Surname"`
	Position string `json:"Position"`
	PayMount float64 `json:"PayMount"`
	Tax float64 `json:"Tax"`
}