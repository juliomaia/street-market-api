package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/juliomaia/street-market-api/api/presenter"
	"github.com/juliomaia/street-market-api/entity"
	"github.com/juliomaia/street-market-api/usecase/streetmarket"
)

func listStreetMarkets(service streetmarket.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading street markets"
		var data []*entity.StreetMarket
		var err error

		distrito := r.URL.Query().Get("distrito")
		regiao5 := r.URL.Query().Get("regiao5")
		nomeFeira := r.URL.Query().Get("nome_feira")
		bairro := r.URL.Query().Get("bairro")

		switch {
		case nomeFeira != "" || distrito != "" || regiao5 != "" || bairro != "":
			data, err = service.SearchStreetMarkets(distrito, regiao5, nomeFeira, bairro)
		default:
			data, err = service.ListStreetMarkets()
		}

		w.Header().Set("Content-Type", "application/json")

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.StreetMarket
		for _, d := range data {
			toJ = append(toJ, &presenter.StreetMarket{
				ID:         d.ID,
				Long:       d.Long,
				Lat:        d.Lat,
				Setcens:    d.Setcens,
				Areap:      d.Areap,
				Coddist:    d.Coddist,
				Distrito:   d.Distrito,
				Codsubpref: d.Codsubpref,
				Subprefe:   d.Subprefe,
				Regiao5:    d.Regiao5,
				Regiao8:    d.Regiao8,
				NomeFeira:  d.NomeFeira,
				Registro:   d.Registro,
				Logradouro: d.Logradouro,
				Numero:     d.Numero,
				Bairro:     d.Bairro,
				Referencia: d.Referencia,
				CreatedAt:  d.CreatedAt,
				UpdatedAt:  d.UpdatedAt,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func getStreetMarket(service streetmarket.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading streetmarket"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetStreetMarket(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.StreetMarket{
			ID:         data.ID,
			Long:       data.Long,
			Lat:        data.Lat,
			Setcens:    data.Setcens,
			Areap:      data.Areap,
			Coddist:    data.Coddist,
			Distrito:   data.Distrito,
			Codsubpref: data.Codsubpref,
			Subprefe:   data.Subprefe,
			Regiao5:    data.Regiao5,
			Regiao8:    data.Regiao8,
			NomeFeira:  data.NomeFeira,
			Registro:   data.Registro,
			Logradouro: data.Logradouro,
			Numero:     data.Numero,
			Bairro:     data.Bairro,
			Referencia: data.Referencia,
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createStreetMarket(service streetmarket.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding street market"
		var input struct {
			NomeFeira  string  `json:"nome_feira"`
			Logradouro string  `json:"logradouro"`
			Numero     string  `json:"numero"`
			Bairro     string  `json:"bairro"`
			Coddist    int     `json:"coddist"`
			Distrito   string  `json:"distrito"`
			Codsubpref int     `json:"codsubpref"`
			Subprefe   string  `json:"subpref"`
			Regiao5    string  `json:"regiao5"`
			Regiao8    string  `json:"regiao8"`
			Registro   string  `json:"registro"`
			Long       float64 `json:"long"`
			Lat        float64 `json:"lat"`
			Setcens    int     `json:"setcens"`
			Areap      int     `json:"areap"`
			Referencia string  `json:"referencia"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		// id, err := service.CreateBook(input.Title, input.Author, input.Pages, input.Quantity)
		id, err := service.CreateStreetMarket(input.Long, input.Lat, input.Setcens, input.Areap,
			input.Coddist, input.Distrito, input.Codsubpref,
			input.Subprefe, input.Regiao5, input.Regiao8, input.NomeFeira,
			input.Registro, input.Logradouro, input.Numero, input.Bairro, input.Referencia)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.StreetMarket{
			ID:         id,
			Long:       input.Long,
			Lat:        input.Lat,
			Setcens:    input.Setcens,
			Areap:      input.Areap,
			Coddist:    input.Coddist,
			Distrito:   input.Distrito,
			Codsubpref: input.Codsubpref,
			Subprefe:   input.Subprefe,
			Regiao5:    input.Regiao5,
			Regiao8:    input.Regiao8,
			NomeFeira:  input.NomeFeira,
			Registro:   input.Registro,
			Logradouro: input.Logradouro,
			Numero:     input.Numero,
			Bairro:     input.Bairro,
			Referencia: input.Referencia,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func deleteStreetMarket(service streetmarket.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing street market"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteStreetMarket(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeStreetMarketHandlers make url handlers
func MakeStreetMarketHandlers(r *mux.Router, n negroni.Negroni, service streetmarket.UseCase) {
	r.Handle("/v1/streetmarket", n.With(
		negroni.Wrap(listStreetMarkets(service)),
	)).Methods("GET", "OPTIONS").Name("listStreetMarkets")

	r.Handle("/v1/streetmarket", n.With(
		negroni.Wrap(createStreetMarket(service)),
	)).Methods("POST", "OPTIONS").Name("createStreetMarket")

	r.Handle("/v1/streetmarket/{id}", n.With(
		negroni.Wrap(getStreetMarket(service)),
	)).Methods("GET", "OPTIONS").Name("getStreetMarket")

	r.Handle("/v1/streetmarket/{id}", n.With(
		negroni.Wrap(deleteStreetMarket(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteStreetMarket")
}
