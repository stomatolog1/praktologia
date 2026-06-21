package model


type WorkSpace struct{
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"Name"`
	Discription string `json:"Discription"`
	Projects []Project
}

type RequestWorkSpace struct{
	Name string `json:"Name"`
	Discription string `json:"Discription"`
}