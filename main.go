package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/stomatolog1/praktologia/internal/handlers"
	"github.com/stomatolog1/praktologia/internal/repository"
	"github.com/stomatolog1/praktologia/internal/servise"
)

func main() {
	// Получить текущую директорию
	wd, _ := os.Getwd()
	staticDir := filepath.Join(wd, "static")

	// Инициализация in-memory репозиториев (без БД)
	adminRepo := repository.NewAdminAkkRepository(nil)
	customerRepo := repository.NewCustomerRepository(nil)
	equipmentRepo := repository.NewEquipmentRepository(nil)
	executorRepo := repository.NewExecutorRepository(nil)
	projectRepo := repository.NewProjectRepository(nil)
	sotrudnikRepo := repository.NewSotrudnikRepository(nil)
	workspaceRepo := repository.NewWorkSpaceRepository(nil)

	// Инициализация сервисов
	adminService := servise.NewAdminAkkService(adminRepo)
	customerService := servise.NewCustomerService(customerRepo)
	equipmentService := servise.NewEquipmentService(equipmentRepo)
	executorService := servise.NewExecutorService(executorRepo)
	projectService := servise.NewProjectService(projectRepo)
	sotrudnikService := servise.NewSotrudnikService(sotrudnikRepo)
	workspaceService := servise.NewWorkSpaceService(workspaceRepo)

	// Инициализация Gin
	router := gin.Default()

	// Инициализация handlers и регистрация маршрутов
	adminHandler := handlers.NewAdminHandler(adminService)
	adminHandler.RegisterRoutes(router)

	customerHandler := handlers.NewCustomerHandler(customerService)
	customerHandler.RegisterRoutes(router)

	equipmentHandler := handlers.NewEquipmentHandler(equipmentService)
	equipmentHandler.RegisterRoutes(router)

	executorHandler := handlers.NewExecutorHandler(executorService)
	executorHandler.RegisterRoutes(router)

	projectHandler := handlers.NewProjectHandler(projectService)
	projectHandler.RegisterRoutes(router)

	sotrudnikHandler := handlers.NewSotrudnikHandler(sotrudnikService)
	sotrudnikHandler.RegisterRoutes(router)

	workspaceHandler := handlers.NewWorkSpaceHandler(workspaceService)
	workspaceHandler.RegisterRoutes(router)

	// Обслуживание статических файлов в конце
	router.Static("/static", staticDir)
	router.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(staticDir, "index.html"))
	})

	// Запуск сервера
	router.Run(":8080")
}



