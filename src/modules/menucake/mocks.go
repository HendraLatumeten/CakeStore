package menucake

import (
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/stretchr/testify/mock"
)

type RepoMenuMock struct {
	mock mock.Mock
}

func (m *RepoMenuMock) ListCake() (*models.MenuCakeAll, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.MenuCakeAll), nil
}

func (m *RepoMenuMock) DetailCake(id int) (*models.MenuCakeAll, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.MenuCakeAll), nil
}
func (m *RepoMenuMock) SaveCake(data *models.MenuCake) (*models.MenuCake, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.MenuCake), nil
}
func (m *RepoMenuMock) DeleteCake(id int) (*models.MenuCake, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.MenuCake), nil
}

func (m *RepoMenuMock) UpdateCake(id string, data *models.MenuCake) (*models.MenuCake, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.MenuCake), nil
}
