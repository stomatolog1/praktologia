package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
)

type CommercialOfferRepository struct {
	projectRepo  *ProjectRepository
	customerRepo *CustomerRepository
}

func NewCommercialOfferRepository(projectRepo *ProjectRepository, customerRepo *CustomerRepository) *CommercialOfferRepository {
	return &CommercialOfferRepository{
		projectRepo:  projectRepo,
		customerRepo: customerRepo,
	}
}

func (r *CommercialOfferRepository) GetCommercialOfferData(projectID uint, customerID uint) (*model.CommercialOffer, error) {
    // В реальном приложении здесь будет логика получения данных из БД
    // и формирование коммерческого предложения.
    // Для примера просто создадим моковые данные.

    project, err := r.projectRepo.GetProjectByID(projectID)
    if err != nil {
        return nil, err
    }

    customer, err := r.customerRepo.GetCustomerByID(customerID)
    if err != nil {
        return nil, err
    }

    offer := &model.CommercialOffer{
        CustomerName:     customer.Name,
        CustomerEmail:    customer.Email,
        CompanyName:      customer.Firm,
        Date:             project.Deadline, // Пример
        ProjectName:      project.Name,
        TechnicalTask:    "Пример технического задания", // Пример
        CompanyDirectorName: customer.Firm, // Пример
    }

    var totalCost float64
    var services []model.Service

    // Пример услуг
    service1 := model.Service{
        Name:     "Разработка дизайна",
        Hours:    40,
        Interval: "2 недели",
        Cost:     100000,
    }
    services = append(services, service1)
    totalCost += service1.Cost

    service2 := model.Service{
        Name:     "Разработка серверной части",
        Hours:    80,
        Interval: "4 недели",
        Cost:     200000,
    }
    services = append(services, service2)
    totalCost += service2.Cost

    offer.Services = services
    offer.TotalCost = totalCost

    return offer, nil
}
