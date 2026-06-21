package model

type AdminAkk struct{
	ID uint `json:"id" gorm:"primaryKey"`
	WorkSpace []int `json:"WorkSpace,omitempty"`
	Login string `json:"Login"`
	HashPass string `json:"HashPass"`
}

type RequestAdminAkk struct{
	Login string `json:"Login"`
	HashPass string `json:"HashPass"`
}