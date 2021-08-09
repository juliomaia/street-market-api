package entity

import (
	"math/big"
	"time"
)

//StreetMarket data
type StreetMarket struct {
	ID         ID
	Long       float32
	Lat        float32
	Setcens    big.Int
	Areap      big.Int
	Coddist    int
	Distrito   string
	Codsubpref int
	Subprefe   string
	Regiao5    string
	Regiao8    string
	NomeFeira  string
	Registro   string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//NewStreetMarket create a new StreetMarket
func NewStreetMarket(long float32, lat float32, setcens big.Int, areap big.Int, coddist int, distrito string, codsubpref int, subprefe string, regiao5 string, regiao8 string, nomeFeira string, registro string, logradouro string, numero string, bairro string, referencia string) (*StreetMarket, error) {
	sm := &StreetMarket{
		ID:         NewID(),
		Long:       long,
		Lat:        lat,
		Setcens:    setcens,
		Areap:      areap,
		Coddist:    coddist,
		Distrito:   distrito,
		Codsubpref: codsubpref,
		Subprefe:   subprefe,
		Regiao5:    regiao5,
		Regiao8:    regiao8,
		NomeFeira:  nomeFeira,
		Registro:   registro,
		Logradouro: logradouro,
		Numero:     numero,
		Bairro:     bairro,
		Referencia: referencia,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Time{},
	}

	err := sm.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return sm, nil
}

//Validate street-market
func (sm *StreetMarket) Validate() error {

	// TODO: add all fields validation
	if sm.NomeFeira == "" || sm.Logradouro == "" {
		return ErrInvalidEntity
	}

	return nil
}
