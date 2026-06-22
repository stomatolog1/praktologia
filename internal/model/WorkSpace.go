package model

type WorkspaceResource struct {
	Type       string  `json:"type"`
	ResourceID uint    `json:"resourceId"`
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	UnitCost   float64 `json:"unitCost"`
	LineCost   float64 `json:"lineCost"`
}

type WorkSpace struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	AdminID     uint                `json:"admin_id"`
	Name        string              `json:"name"`
	Discription string              `json:"description,omitempty"`
	TotalCost   float64             `json:"totalCost"`
	Resources   []WorkspaceResource `json:"resources,omitempty" gorm:"-"`
	Projects    []Project           `json:"projects" gorm:"foreignKey:WorkSpaceID"`
}

type RequestWorkSpace struct {
	AdminID     uint                `json:"admin_id"`
	Name        string              `json:"name"`
	Discription string              `json:"description,omitempty"`
	TotalCost   float64             `json:"totalCost"`
	Resources   []WorkspaceResource `json:"resources,omitempty"`
}
