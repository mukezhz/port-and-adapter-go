package persistent

import (
	"fmt"
	"myapp/internal/domain/models"
)

type MockRepository struct {
	processes []models.Process
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		processes: []models.Process{
			{Name: "mock"},
		},
	}
}

func (repo MockRepository) Find(data string) bool {
	return repo.processes[0].Name == data
}

func (repo MockRepository) Create(data string) bool {
	fmt.Println("mock create")
	return true
}
