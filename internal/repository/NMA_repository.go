package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
)

type NMARepository struct {
	projectRepo   *ProjectRepository
	sotrudnikRepo *SotrudnikRepository
	equipmentRepo *EquipmentRepository
}

func NewNMARepository(projectRepo *ProjectRepository, sotrudnikRepo *SotrudnikRepository, equipmentRepo *EquipmentRepository) *NMARepository {
	return &NMARepository{
		projectRepo:   projectRepo,
		sotrudnikRepo: sotrudnikRepo,
		equipmentRepo: equipmentRepo,
	}
}

func (r *NMARepository) GetNMAData(projectID uint) (*model.NMA, error) {
    // В реальном приложении здесь будет логика получения данных из БД
    // и расчет стоимости нематериального актива.
    // Для примера просто создадим моковые данные.

    project, err := r.projectRepo.GetProjectByID(projectID)
    if err != nil {
        return nil, err
    }

    nma := &model.NMA{
        ProjectName:     project.Name,
        DevelopmentTime: project.Deadline,
    }

    var totalCost float64
    var resources []model.Resource

    for _, worker := range project.Workers {
        resource := model.Resource{
            Name:     worker.Name + " " + worker.SecondName,
            Cost:     worker.PayMount,
            CostType: "Зарплата",
        }
        resources = append(resources, resource)
        totalCost += worker.PayMount
    }

    for _, equipment := range project.Equipments {
        resource := model.Resource{
            Name:     equipment.Name,
            Cost:     equipment.Cost,
            CostType: "Оборудование",
        }
        resources = append(resources, resource)
        totalCost += equipment.Cost
    }

    nma.TotalCost = totalCost
    nma.Resources = resources

    return nma, nil
}
