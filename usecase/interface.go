package streetmarket

import (
	"github.com/juliomaia/street-market-api/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.StreetMarket, error)
	Search(query string) ([]*entity.StreetMarket, error)
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
	SearchStreetMarkets(query string) ([]*entity.StreetMarket, error)
	ListStreetMarkets() ([]*entity.StreetMarket, error)
	// CreateStreetMarket(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateStreetMarket(e *entity.StreetMarket) error
	DeleteStreetMarket(id entity.ID) error
}
