// package book

// import (
// 	"testing"
// 	"time"

// 	"github.com/juliomaia/street-market-api/entity"

// 	"github.com/stretchr/testify/assert"
// )

// func newFixtureStreetMarket() *entity.StreetMarket {
// 	return &entity.StreetMarket{
// 		Title:     "I Am Ozzy",
// 		Author:    "Ozzy Osbourne",
// 		Pages:     294,
// 		Quantity:  1,
// 		CreatedAt: time.Now(),
// 	}
// }

// func Test_Create(t *testing.T) {
// 	repo := newInmem()
// 	m := NewService(repo)
// 	u := newFixtureStreetMarket()
// 	_, err := m.CreateStreetMarket(u.Title, u.Author, u.Pages, u.Quantity)
// 	assert.Nil(t, err)
// 	assert.False(t, u.CreatedAt.IsZero())
// }

// func Test_SearchAndFind(t *testing.T) {
// 	repo := newInmem()
// 	m := NewService(repo)
// 	u1 := newFixtureStreetMarket()
// 	u2 := newFixtureStreetMarket()
// 	u2.Title = "Lemmy: Biography"

// 	uID, _ := m.CreateStreetMarket(u1.Title, u1.Author, u1.Pages, u1.Quantity)
// 	_, _ = m.CreateStreetMarket(u2.Title, u2.Author, u2.Pages, u2.Quantity)

// 	t.Run("search", func(t *testing.T) {
// 		c, err := m.SearchStreetMarkets("ozzy")
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, len(c))
// 		assert.Equal(t, "I Am Ozzy", c[0].Title)

// 		c, err = m.SearchStreetMarkets("dio")
// 		assert.Equal(t, entity.ErrNotFound, err)
// 		assert.Nil(t, c)
// 	})
// 	t.Run("list all", func(t *testing.T) {
// 		all, err := m.ListStreetMarkets()
// 		assert.Nil(t, err)
// 		assert.Equal(t, 2, len(all))
// 	})

// 	t.Run("get", func(t *testing.T) {
// 		saved, err := m.GetStreetMarket(uID)
// 		assert.Nil(t, err)
// 		assert.Equal(t, u1.Title, saved.Title)
// 	})
// }

// func Test_Update(t *testing.T) {
// 	repo := newInmem()
// 	m := NewService(repo)
// 	u := newFixtureStreetMarket()
// 	id, err := m.CreateStreetMarket(u.Title, u.Author, u.Pages, u.Quantity)
// 	assert.Nil(t, err)
// 	saved, _ := m.GetStreetMarket(id)
// 	saved.Title = "Lemmy: Biography"
// 	assert.Nil(t, m.UpdateStreetMarket(saved))
// 	updated, err := m.GetStreetMarket(id)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "Lemmy: Biography", updated.Title)
// }

// func TestDelete(t *testing.T) {
// 	repo := newInmem()
// 	m := NewService(repo)
// 	u1 := newFixtureStreetMarket()
// 	u2 := newFixtureStreetMarket()
// 	u2ID, _ := m.CreateStreetMarket(u2.Title, u2.Author, u2.Pages, u2.Quantity)

// 	err := m.DeleteStreetMarket(u1.ID)
// 	assert.Equal(t, entity.ErrNotFound, err)

// 	err = m.DeleteStreetMarket(u2ID)
// 	assert.Nil(t, err)
// 	_, err = m.GetStreetMarket(u2ID)
// 	assert.Equal(t, entity.ErrNotFound, err)
// }
