package menucake

import (
	"fmt"

	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/interfaces"
)

type menucake_service struct {
	repo interfaces.MenuCakeRepo
}

func NewService(reps interfaces.MenuCakeRepo) *menucake_service {
	return &menucake_service{reps}
}

func (r *menucake_service) GetAll() (*models.MenuCake, error) {
	data, err := r.repo.FindAll()
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
