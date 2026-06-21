package model

type AdminAkk struct{
	ID uint `json:"ID" gorm::"primarykey"`
	Login string `json:"Login"`
	HashPass string `json:"Login"`
}

type RequestAdminAkk struct{
	Login string `json:"Login"`
	HashPass string `json:"Login"`
}