package entity_test

import (
	"testing"

	"github.com/juliomaia/street-market-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStreetMarket(t *testing.T) {
	longitude := -23.538862
	latitude := -46.6048588
	setcens := 66645649999
	areap := 99945649999
	coddist := 99
	distrito := "Brás"
	codsubpref := 101
	subprefe := "Brás"
	regiao5 := "Leste"
	regiao8 := "Leste 1"
	nomeFeira := "Feira do Rolo"
	registro := "9099-9"
	logradouro := "Avenida Celso Garcia"
	numero := "958"
	bairro := "Brás"
	referencia := "Praça do Rolo"

	sm, err := entity.NewStreetMarket(longitude, latitude, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nomeFeira, registro, logradouro, numero, bairro, referencia)
	assert.Nil(t, err)
	assert.Equal(t, sm.NomeFeira, "Feira do Rolo")
	assert.Equal(t, sm.Logradouro, "Avenida Celso Garcia")
	assert.Equal(t, sm.Coddist, 99)
	assert.Equal(t, sm.Regiao5, "Leste")
	assert.NotNil(t, sm.ID)
}
