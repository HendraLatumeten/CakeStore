package menucake

import (
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gorilla/mux"
	"github.com/hendralatumeten/cakestore/src/database/models"
	"github.com/hendralatumeten/cakestore/src/libs"
	"github.com/stretchr/testify/mock"
)

var dataMock = models.MenuCakeAll{
	{ID: 2, Title: "test", Description: "okeebanget", Rating: 5, Image: "test.jpg"},
	{ID: 3, Title: "test", Description: "okeebanget", Rating: 5, Image: "test.jpg"},
}

var repos = RepoMenuMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

func TestListAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	//repo
	repos.mock.On("ListCake").Return(&dataMock, nil)

	//router
	mux.HandleFunc("/test/menucake", ctrl.ListAll)
	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/menucake", nil))

	//response models
	var menu models.MenuCake
	response := libs.Response{
		Data: &menu,
	}

	// if err := json.NewDecoder(w.Body); err != nil {
	// 	t.Fatal(err)
	// }

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, response.IsError, "isError must be false")
}
