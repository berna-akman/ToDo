package persistence

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"to-do-api/pg/infrastructure/config"
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

func InitForTest() (*DataSource, sqlmock.Sqlmock, error) {
	mockConn, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockConn,
	}), &gorm.Config{NowFunc: func() time.Time {
		return time.Now().UTC()
	}})

	if err != nil {
		panic(err)
	}

	pg := &DataSource{DB: db}
	return pg, mock, nil
}
