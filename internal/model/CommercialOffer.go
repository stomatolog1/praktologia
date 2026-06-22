package model

type CommercialOffer struct {
    CustomerName         string    `json:"customerName"`
    CustomerPosition     string    `json:"customerPosition"`
    CustomerEmail        string    `json:"customerEmail"`
    CompanyName          string    `json:"companyName"`
    Date                 string    `json:"date"`
    ProjectName          string    `json:"projectName"`
    TechnicalTask        string    `json:"technicalTask"`
    TotalCost            float64   `json:"totalCost"`
    Services             []Service `json:"services"`
    CompanyDirectorName  string    `json:"companyDirectorName"`
}

type Service struct {
    Name         string  `json:"name"`
    Hours        int     `json:"hours"`
    Interval     string  `json:"interval"`
    Cost         float64 `json:"cost"`
}
