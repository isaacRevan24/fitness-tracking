//go:generate go run github.com/golang/mock/mockgen -source repository.go -destination mock/repository_mock.go -package mock
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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
	GetWeightRegister(request model.GetWeightRegisterReq)
}

func NewTrackingRepository() TrackingRepository {
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

func (r *Repo) AddWeightRegister(request model.AddWeightRegisterReq) (string, error) {
	sqlStatement := `INSERT INTO weight_track(weight_id, id, weight, created_at) VALUES($1, $2, $3, $4)`
	id := uuid.New().String()
	_, insertError := r.db.Exec(sqlStatement, id, request.ClientId, request.Weight, request.CreatedAt)
	if insertError != nil {
		return "", insertError
	}
	return id, nil
}

func (r *Repo) GetWeightRegister(request model.GetWeightRegisterReq) {
	sqlStatement := `SELECT created_at, weight FROM weight_track WHERE id=$1 AND weight_track_id=$2`
	var created_at time.Time
	var weight float32
	r.db.QueryRow(sqlStatement, request.ClientId, request.WeightTrackId).Scan(&created_at, &weight)
}
