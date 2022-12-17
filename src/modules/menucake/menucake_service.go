package menucake

import (
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/interfaces"
	"github.com/hendralatumeten/cakestore/src/libs"
)

type menucake_service struct {
	repo interfaces.MenuCakeRepo
}

func NewService(repo interfaces.MenuCakeRepo) *menucake_service {
	return &menucake_service{repo}
}

func (s *menucake_service) List() *libs.Response {
	data, err := s.repo.ListCake()

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (s *menucake_service) Detail(id int) *libs.Response {
	data, err := s.repo.DetailCake(id)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (s *menucake_service) Add(data *models.MenuCake) *libs.Response {
	data, err := s.repo.SaveCake(data)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(data, 200, false)
}

func (s *menucake_service) Delete(id int) *libs.Response {
	data, err := s.repo.DeleteCake(id)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (s *menucake_service) Update(id string, data *models.MenuCake) *libs.Response {
	data, err := s.repo.UpdateCake(id, data)

	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	return libs.Respone(data, 200, false)
}
