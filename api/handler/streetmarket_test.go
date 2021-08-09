package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/juliomaia/street-market-api/entity"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/juliomaia/street-market-api/usecase/book/mock"
	"github.com/stretchr/testify/assert"
)

func Test_createStreetMarket(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeStreetMarketHandlers(r, *n, service)
	path, err := r.GetRoute("createStreetMarket").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/streetmarket", path)

	service.EXPECT().
		CreateStreetMarket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	h := createStreetMarket(service)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
  "title": "I Am Ozzy",
  "author": "Ozzy Osbourne",
  "pages": 294,
  "quantity":1


}`)
	resp, _ := http.Post(ts.URL+"/v1/streetmarket", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var sm *entity.StreetMarket
	json.NewDecoder(resp.Body).Decode(&sm)
	assert.Equal(t, "Feira de Quarta", sm.NomeFeira)
}
