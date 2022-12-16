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

func NewCtrl(reps interfaces.MenuCakeService) *menucake_ctrl {
	return &menucake_ctrl{reps}
}

func (re *menucake_ctrl) ListAll(w http.ResponseWriter, r *http.Request) {
	data, err := re.svc.List()
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(&data, 200, false).Send(w)
}

func (re *menucake_ctrl) DetailAll(w http.ResponseWriter, r *http.Request) {
	a := mux.Vars(r)["id"]
	id, err := strconv.Atoi(a)
	data, err := re.svc.Detail(id)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(&data, 200, false).Send(w)
}

func (re *menucake_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.MenuCake

	err := json.NewDecoder(r.Body).Decode(&data.Title)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	re.svc.Add(&data).Send(w)

}
