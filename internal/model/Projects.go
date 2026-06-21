package model

type Project struct{
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"Name"`
	Discription string `json:"Discription"`
	Deadline string `json:"Deadline"` 
	Price float64 `json:"Price"` 
	Workers []Sotrudnik `json:"Workers" gorm:"-"`
	Equipments []Equipment `json:"Equipments" gorm:"-"`
}

type RequestProject struct{
	Name string `json:"Name"`
	Discription string `json:"Discription"`
	Deadline string `json:"Deadline"` 
	Price float64 `json:"Price"` 
	Workers []Sotrudnik `json:"Workers" gorm:"-"`
	Equipments []Equipment `json:"Equipments" gorm:"-"`
}