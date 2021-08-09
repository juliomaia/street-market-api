package repository

import (
	"database/sql"
	"time"

	"github.com/juliomaia/street-market-api/entity"
)

//StreetMarketMySQL mysql repo
type StreetMarketMySQL struct {
	db *sql.DB
}

//NewStreetMarketMySQL create new repository
func NewStreetMarketMySQL(db *sql.DB) *StreetMarketMySQL {
	return &StreetMarketMySQL{
		db: db,
	}
}

//TODO
//Create a street market
func (r *StreetMarketMySQL) Create(e *entity.StreetMarket) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into street_market ( created_at) 
		values(?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.NomeFeira,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

//Get a street market
func (r *StreetMarketMySQL) Get(id entity.ID) (*entity.StreetMarket, error) {

	stmt, err := r.db.Prepare(`select * from street_market where id = ?`)
	if err != nil {
		return nil, err
	}

	var sm entity.StreetMarket

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&sm.ID, &sm.Long, &sm.Lat, &sm.Setcens, &sm.Areap, sm.Coddist,
			&sm.Distrito, &sm.Codsubpref, &sm.Subprefe,
			&sm.Regiao5, &sm.Regiao8, &sm.NomeFeira, &sm.Registro,
			&sm.Logradouro, &sm.Numero, &sm.Bairro, &sm.Referencia,
			&sm.CreatedAt, &sm.UpdatedAt)

	}

	return &sm, nil
}

//Update a book
func (r *StreetMarketMySQL) Update(e *entity.StreetMarket) error {
	//TODO SQL
	// e.UpdatedAt = time.Now()
	// _, err := r.db.Exec("update book set title = ?, author = ?, pages = ?, quantity = ?, updated_at = ? where id = ?", e.Title, e.Author, e.Pages, e.Quantity, e.UpdatedAt.Format("2006-01-02"), e.ID)
	// if err != nil {
	// 	return err
	// }
	return nil
}

//Search street markets
func (r *StreetMarketMySQL) Search(query string) ([]*entity.StreetMarket, error) {
	stmt, err := r.db.Prepare(`select * from street_market where name_feira like ?`)
	if err != nil {
		return nil, err
	}

	var streetMarkets []*entity.StreetMarket

	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sm entity.StreetMarket

		err = rows.Scan(&sm.ID, &sm.Long, &sm.Lat, &sm.Setcens, &sm.Areap, sm.Coddist,
			&sm.Distrito, &sm.Codsubpref, &sm.Subprefe,
			&sm.Regiao5, &sm.Regiao8, &sm.NomeFeira, &sm.Registro,
			&sm.Logradouro, &sm.Numero, &sm.Bairro, &sm.Referencia,
			&sm.CreatedAt, &sm.UpdatedAt)

		if err != nil {
			return nil, err
		}
		streetMarkets = append(streetMarkets, &sm)
	}

	return streetMarkets, nil
}

//List street markets
func (r *StreetMarketMySQL) List() ([]*entity.StreetMarket, error) {

	stmt, err := r.db.Prepare(`select * from street_market order by 1 desc`)
	if err != nil {
		return nil, err
	}

	var streetMarkets []*entity.StreetMarket

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var sm entity.StreetMarket

		err = rows.Scan(&sm.ID, &sm.Long, &sm.Lat, &sm.Setcens, &sm.Areap, sm.Coddist,
			&sm.Distrito, &sm.Codsubpref, &sm.Subprefe,
			&sm.Regiao5, &sm.Regiao8, &sm.NomeFeira, &sm.Registro,
			&sm.Logradouro, &sm.Numero, &sm.Bairro, &sm.Referencia,
			&sm.CreatedAt, &sm.UpdatedAt)

		if err != nil {
			return nil, err
		}
		streetMarkets = append(streetMarkets, &sm)
	}

	return streetMarkets, nil
}

//Delete a street market
func (r *StreetMarketMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from street_market where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
