package model

type NMA struct {
    ProjectName      string      `json:"projectName"`
    DevelopmentTime  string      `json:"developmentTime"`
    TotalCost        float64     `json:"totalCost"`
    Resources        []Resource  `json:"resources"`
}

type Resource struct {
    Name      string  `json:"name"`
    Cost      float64 `json:"cost"`
    CostType  string  `json:"costType"`
}
