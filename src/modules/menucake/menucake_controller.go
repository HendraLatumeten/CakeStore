package menucake

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/interfaces"
	"github.com/hendralatumeten/cakestore/src/libs"
)

type menucake_ctrl struct {
	svc interfaces.MenuCakeService
}

func NewCtrl(ctrl interfaces.MenuCakeService) *menucake_ctrl {
	return &menucake_ctrl{ctrl}
}

func (c *menucake_ctrl) ListAll(w http.ResponseWriter, r *http.Request) {
	c.svc.List().Send(w)
}

func (c *menucake_ctrl) DetailAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}
	c.svc.Detail(id).Send(w)

}

func (c *menucake_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	var data models.MenuCake

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	c.svc.Add(&data)

}

func (c *menucake_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	c.svc.Delete(id).Send(w)
}

func (c *menucake_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	var data models.MenuCake
	id := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	c.svc.Update(id, &data).Send(w)

}

func (c *menucake_ctrl) Test(w http.ResponseWriter, r *http.Request) {

	libs.Respone("Hello Cakestore", 200, true).Send(w)

}
