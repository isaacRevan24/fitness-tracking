package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/isaacRevan24/fitness-tracking/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "postgres"
	dbname   = "fitness"
)

type repo struct {
	db *sql.DB
}

type TrackingRepository interface {
	AddWeightRegister(request *model.AddWeightRegisterReq) (string, error)
}

func NewTrackingRepository() TrackingRepository {
	repo, _ := getConnection()
	return repo
}

func getConnection() (*repo, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return nil, err
	}

	return &repo{db: db}, nil
}

func (r *repo) AddWeightRegister(request *model.AddWeightRegisterReq) (string, error) {
	defer r.db.Close()
	sqlStatement := `INSERT INTO weight_track(id, weight, created_at) VALUES($1, $2, $3) RETURNING weight_track_id`
	id := uuid.UUID{}
	insertError := r.db.QueryRow(sqlStatement, &request.ClientId, &request.Weight, &request.CreatedAt).Scan(id)
	if insertError != nil {
		return "", insertError
	}
	return id.String(), nil
}
