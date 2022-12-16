package menucake

import (
	"fmt"

	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/interfaces"
	"github.com/hendralatumeten/cakestore/src/libs"
)

type menucake_service struct {
	repo interfaces.MenuCakeRepo
}

func NewService(reps interfaces.MenuCakeRepo) *menucake_service {
	return &menucake_service{reps}
}

func (r *menucake_service) List() (*models.MenuCakeAll, error) {
	data, err := r.repo.ListCake()

	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *menucake_service) Detail(id int) (*models.MenuCakeAll, error) {
	data, err := r.repo.DetailCake(id)

	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *menucake_service) Add(data *models.MenuCake) *libs.Response {
	data, err := r.repo.Save(data)
	if data != nil {
		fmt.Println("Add Data Berhasil")
	}
	if err != nil {
		panic(err)
	}

	return libs.Respone("data berhasil ditambhkan", 200, false)
}
