//go:generate go run github.com/golang/mock/mockgen -source repository.go -destination mock/repository_mock.go -package mock
package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/isaacRevan24/fitness-tracking/model"
	_ "github.com/lib/pq"
)

func init() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "fitness"
	connection = connectionInfo{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

var connection connectionInfo

type connectionInfo struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

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
	GetGoalsRegister(clientId string) (model.GetGoalsRes, error)
	UpdateWeightGoal(request model.UpdateWeightGoalReq) (model.UpdateWeightGoalRes, error)
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
	fmt.Println("connection getter")
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			connection.host, connection.port, connection.user, connection.password, connection.dbname))
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return nil, err
	}
	return &Repo{db: db}, nil
}
