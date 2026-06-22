package model

type AdminAkk struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	WorkSpace []int  `json:"WorkSpace,omitempty"`
	Login     string `json:"Login"`
	HashPass  string `json:"-"`
}

type RequestAdminAkk struct {
	Login    string `json:"Login"`
	HashPass string `json:"HashPass"`
}

type RequestAdminLogin struct {
	Login    string `json:"Login"`
	HashPass string `json:"HashPass"`
}

type AdminResponse struct {
	ID    uint   `json:"id"`
	Login string `json:"Login"`
}