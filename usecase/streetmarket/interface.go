package streetmarket

import (
	"github.com/juliomaia/street-market-api/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.StreetMarket, error)
	Search(distrito string, regiao5 string, nomeFeira string, bairro string) ([]*entity.StreetMarket, error)
	List() ([]*entity.StreetMarket, error)
}

//Writer streetMarket writer
type Writer interface {
	Create(e *entity.StreetMarket) (entity.ID, error)
	Update(e *entity.StreetMarket) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetStreetMarket(id entity.ID) (*entity.StreetMarket, error)
	SearchStreetMarkets(distrito string, regiao5 string, nomeFeira string, bairro string) ([]*entity.StreetMarket, error)
	ListStreetMarkets() ([]*entity.StreetMarket, error)
	CreateStreetMarket(long float64, lat float64, setcens int, areap int, coddist int, distrito string, codsubpref int, subprefe string, regiao5 string, regiao8 string, nomeFeira string, registro string, logradouro string, numero string, bairro string, referencia string) (entity.ID, error)
	UpdateStreetMarket(e *entity.StreetMarket) error
	DeleteStreetMarket(id entity.ID) error
}
