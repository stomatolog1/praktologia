package model


type WorkSpace struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	AdminID     uint      `json:"admin_id gorm:"foreignKey:AdminID""`
	Name        string    `json:"name"`
	Discription string    `json:"description,omitempty"`
	Projects    []Project `json:"projects" gorm:"foreignKey:WorkSpaceID"`
}

type RequestWorkSpace struct {
	AdminID     uint   `json:"admin_id"`
	Name        string `json:"name"`
	Discription string `json:"description,omitempty"`
}