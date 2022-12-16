package interfaces

import (
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/libs"
)

type MenuCakeRepo interface {
	ListCake() (*models.MenuCakeAll, error)
	DetailCake(id int) (*models.MenuCakeAll, error)
	Save(data *models.MenuCake) (*models.MenuCake, error)
}

type MenuCakeService interface {
	List() (*models.MenuCakeAll, error)
	Detail(id int) (*models.MenuCakeAll, error)
	Add(data *models.MenuCake) *libs.Response
}
