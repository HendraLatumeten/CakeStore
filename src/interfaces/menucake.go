package interfaces

import (
	"github.com/hendralatumeten/cakestore/src/database/models"
)

type MenuCakeRepo interface {
	FindAll() (*models.MenuCake, error)
	// Save(data *models.MenuCake) (*models.MenuCake, error)
}

type MenuCakeService interface {
	GetAll() (*models.MenuCake, error)
	// Add(data *models.MenuCake) *libs.Response
}
