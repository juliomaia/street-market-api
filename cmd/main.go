package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juliomaia/street-market-api/config"
	"github.com/juliomaia/street-market-api/infrastructure/repository"
	"github.com/juliomaia/street-market-api/pkg/metric"
	"github.com/juliomaia/street-market-api/usecase/streetmarket"
)

func handleParams() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Invalid query")
	}
	return os.Args[1], nil
}

func main() {
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	appMetric := metric.NewCLI("search")
	appMetric.Started()
	query, err := handleParams()
	if err != nil {
		log.Fatal(err.Error())
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	repo := repository.NewStreetMarketMySQL(db)
	service := streetmarket.NewService(repo)
	all, err := service.SearchStreetMarkets("", "", query, "")
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range all {
		fmt.Printf("%s %s \n", j.NomeFeira, j.Registro)
	}
	appMetric.Finished()
	err = metricService.SaveCLI(appMetric)
	if err != nil {
		log.Fatal(err)
	}
}
