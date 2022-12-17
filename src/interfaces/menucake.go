package interfaces

import (
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/libs"
)

type MenuCakeRepo interface {
	ListCake() (*models.MenuCakeAll, error)
	DetailCake(id int) (*models.MenuCakeAll, error)
	SaveCake(data *models.MenuCake) (*models.MenuCake, error)
	DeleteCake(id int) (*models.MenuCake, error)
	UpdateCake(id string, data *models.MenuCake) (*models.MenuCake, error)
}

type MenuCakeService interface {
	List() *libs.Response
	Detail(id int) *libs.Response
	Add(data *models.MenuCake) *libs.Response
	Delete(id int) *libs.Response
	Update(id string, data *models.MenuCake) *libs.Response
}
