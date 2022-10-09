package repository

import (
	"fmt"

	"github.com/isaacRevan24/fitness-tracking/model"
)

func (r *Repo) AddGoalsRegister(request model.AddWeightGoalsReq) (model.AddWeightGoalsRes, error) {
	sqlStatement := `INSERT INTO goals(id, weight, steps) VALUES($1, $2, $3) RETURNING weight, steps`
	var weight float32 = 0.00
	var steps int32 = 0

	err := r.db.QueryRow(sqlStatement, request.ClientId, request.Weight, request.Steps).Scan(&weight, &steps)

	if err != nil {
		fmt.Println(err)
		return model.AddWeightGoalsRes{}, err
	}

	return model.AddWeightGoalsRes{Weight: weight, Steps: steps}, nil
}

func (r *Repo) GetGoalsRegister(clientId string) (model.GetGoalsRes, error) {
	sqlStatement := `SELECT weight, steps FROM goals WHERE id=$1`
	var goals model.GetGoalsRes
	err := r.db.QueryRow(sqlStatement, clientId).Scan(&goals.Weight, &goals.Steps)
	if err != nil {
		fmt.Println(err)
		return model.GetGoalsRes{}, err
	}
	return goals, nil
}

func (r *Repo) UpdateWeightGoal(request model.UpdateWeightGoalReq) (model.UpdateWeightGoalRes, error) {
	sqlStatement := `UPDATE goals SET weight=$2 WHERE id=$1 RETURNING weight`
	var goals model.UpdateWeightGoalRes
	err := r.db.QueryRow(sqlStatement, request.ClientId, request.Weight).Scan(&goals.Weight)
	if err != nil {
		fmt.Println(err)
		return model.UpdateWeightGoalRes{}, err
	}
	return goals, nil
}
