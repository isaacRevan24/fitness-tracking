package repository

import (
	"fmt"

	"github.com/isaacRevan24/fitness-tracking/model"
)

func (r *Repo) AddGoalsRegisters(request model.AddWeightGoalsReq) (model.AddWeightGoalsRes, error) {
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
