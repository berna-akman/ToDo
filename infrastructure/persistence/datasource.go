package persistence

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"to-do-api/infrastructure/config"
)

type DataSource struct {
	*gorm.DB
}

func Connect(cfg config.Postgres) (*DataSource, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
	pg := postgres.Open(psqlInfo)
	db, err := gorm.Open(pg)

	return &DataSource{db}, err
}
