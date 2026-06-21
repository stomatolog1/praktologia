package model

type Equipment struct {
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"Name"`
	Description string `json:"Description"`
	TypeOperating string `json:"TypeOperating"`
	RentalCost float64 `json:"RentaCost"`
	PayTime string `json:"PayTime"`
	Cost float64 `json:"Cost"`
}
type RequsetEquipment struct{
	Name string `json:"Name"`
	Description string `json:"Description"`
	TypeOperating string `json:"TypeOperating"`
	RentalCost float64 `json:"RentaCost"`
	PayTime string `json:"PayTime"`
	Cost float64 `json:"Cost"`
}