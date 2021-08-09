package streetmarket

import (
	"strings"
	"time"

	"github.com/juliomaia/street-market-api/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//TODO: implement
//CreateStreetMarket create a book
// func (s *Service) CreateStreetMarket(title string, author string, pages int, quantity int) (entity.ID, error) {
// 	b, err := entity.NewStreetMarket(title, author, pages, quantity)
// 	if err != nil {
// 		return b.ID, err
// 	}
// 	return s.repo.Create(b)
// }

//GetStreetMarket get a streetMarket
func (s *Service) GetStreetMarket(id entity.ID) (*entity.StreetMarket, error) {
	sm, err := s.repo.Get(id)
	if sm == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return sm, nil
}

//SearchStreetMarkets search streetMarkets
func (s *Service) SearchStreetMarkets(query string) ([]*entity.StreetMarket, error) {
	streetMarkets, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(streetMarkets) == 0 {
		return nil, entity.ErrNotFound
	}
	return streetMarkets, nil
}

//ListStreetMarkets list streetMarkets
func (s *Service) ListStreetMarkets() ([]*entity.StreetMarket, error) {
	streetMarkets, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(streetMarkets) == 0 {
		return nil, entity.ErrNotFound
	}
	return streetMarkets, nil
}

//DeleteStreetMarket Delete a streetMarket
func (s *Service) DeleteStreetMarket(id entity.ID) error {
	_, err := s.GetStreetMarket(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateStreetMarket Update a streetMarket
func (s *Service) UpdateStreetMarket(e *entity.StreetMarket) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
