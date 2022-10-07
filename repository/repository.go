//go:generate go run github.com/golang/mock/mockgen -source repository.go -destination mock/repository_mock.go -package mock
package repository

import (
	"database/sql"
	"errors"
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
	GetWeightRegister(clientId string, weightId string) (model.GetWeightRegisterRes, error)
	UpdateWeightRegister(request model.UpdateWeightRegisterReq) error
	DeleteWeightRegister(request model.DeleteWeightRegisterReq) error
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
		fmt.Println(insertError)
		return "", insertError
	}
	return id, nil
}

func (r *Repo) GetWeightRegister(clientId string, weightId string) (model.GetWeightRegisterRes, error) {
	sqlStatement := `SELECT created_at, weight FROM weight_track WHERE id=$1 AND weight_id=$2`
	var created_at time.Time
	var weight float32
	error := r.db.QueryRow(sqlStatement, clientId, weightId).Scan(&created_at, &weight)
	if error != nil {
		fmt.Println(error)
		return model.GetWeightRegisterRes{}, error
	}
	return model.GetWeightRegisterRes{Weight: weight, CreatedAt: created_at.String()}, nil
}

func (r *Repo) UpdateWeightRegister(request model.UpdateWeightRegisterReq) error {
	sqlStatement := `UPDATE weight_track SET weight=$1 WHERE weight_id=$2 AND id=$3`
	res, error := r.db.Exec(sqlStatement, request.Weight, request.WeightTrackId, request.ClientId)
	if error != nil {
		fmt.Println(error)
		return error
	}
	count, rowError := res.RowsAffected()
	if rowError != nil {
		fmt.Println(rowError)
		return rowError
	}

	if count < 1 {
		fmt.Println("Registros no afectado")
		return errors.New("row no afectado")
	}
	fmt.Println(count)
	return nil
}

func (r *Repo) DeleteWeightRegister(request model.DeleteWeightRegisterReq) error {
	sqlStatement := `DELETE FROM weight_track WHERE weight_id=$1 AND id=$2`
	res, error := r.db.Exec(sqlStatement, request.WeightTrackId, request.ClientId)
	if error != nil {
		fmt.Println(error)
		return error
	}
	count, rowError := res.RowsAffected()
	if rowError != nil {
		fmt.Println(rowError)
		return rowError
	}

	if count < 1 {
		fmt.Println("Registros no afectado")
		return errors.New("row no afectado")
	}
	return nil
}
