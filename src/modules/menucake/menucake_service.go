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
	data, err := r.repo.SaveCake(data)
	if data != nil {
		fmt.Println("Add Data Berhasil")
	}
	if err != nil {
		panic(err)
	}

	return libs.Respone("data berhasil ditambhkan", 200, false)
}

func (r *menucake_service) Delete(id int) (*models.MenuCake, error) {
	data, err := r.repo.DeleteCake(id)

	if data != nil {
		fmt.Println("Data Berhasil Terhapus")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *menucake_service) Update(id string, data *models.MenuCake) (*models.MenuCake, error) {
	data, err := r.repo.UpdateCake(id, data)

	if data != nil {
		fmt.Println("Data Berhasil Terupdate")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
