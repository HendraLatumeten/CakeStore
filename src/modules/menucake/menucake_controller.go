package menucake

import (
	"net/http"

	"github.com/hendralatumeten/cakestore/src/interfaces"
	"github.com/hendralatumeten/cakestore/src/libs"
)

type menucake_ctrl struct {
	svc interfaces.MenuCakeService
}

func NewCtrl(reps interfaces.MenuCakeService) *menucake_ctrl {
	return &menucake_ctrl{reps}
}

func (re *menucake_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := re.svc.GetAll()
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(data, 200, false).Send(w)
}
