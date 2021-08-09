package presenter

import (
	"math/big"
	"time"

	"github.com/juliomaia/street-market-api/entity"
)

//StreetMarket data
type StreetMarket struct {
	ID         entity.ID `json:"id"`
	NomeFeira  string    `json:"nome_feira"`
	Logradouro string    `json:"logradouro"`
	Numero     string    `json:"numero"`
	Bairro     string    `json:"bairro"`
	Coddist    int       `json:"coddist"`
	Distrito   string    `json:"distrito"`
	Codsubpref int       `json:"codsubpref"`
	Subprefe   string    `json:"subpref"`
	Regiao5    string    `json:"regiao5"`
	Regiao8    string    `json:"regiao8"`
	Registro   string    `json:"registro"`
	Long       float32   `json:"long"`
	Lat        float32   `json:"lat"`
	Setcens    big.Int   `json:"setcens"`
	Areap      big.Int   `json:"areap"`
	Referencia string    `json:"referencia"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
