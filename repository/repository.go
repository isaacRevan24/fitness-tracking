//go:generate go run github.com/golang/mock/mockgen -source repository.go -destination mock/repository_mock.go -package mock
package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/isaacRevan24/fitness-tracking/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "fitness"
)

type Repo struct {
	db *sql.DB
}

type TrackingRepository interface {
	AddWeightRegister(request model.AddWeightRegisterReq) (string, error)
	GetWeightRegister(clientId string, weightId string) (model.GetWeightRegisterRes, error)
	UpdateWeightRegister(request model.UpdateWeightRegisterReq) error
	DeleteWeightRegister(request model.DeleteWeightRegisterReq) error
}

type GoalsRepository interface {
	AddGoalsRegister(request model.AddWeightGoalsReq) (model.AddWeightGoalsRes, error)
}

func NewTrackingRepository() TrackingRepository {
	repo, _ := getConnection()
	return repo
}

func NewGoalsRepository() GoalsRepository {
	repo, _ := getConnection()
	return repo
}

func getConnection() (*Repo, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return nil, err
	}
	return &Repo{db: db}, nil
}
