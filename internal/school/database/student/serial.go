package student

import (
	"context"
	"errors"
)

type studentSerial struct {
	student
	ID int64 `json:"student_id"`
}

func (s *studentSerial) Create(input CreateInput) (*Model, error) {
	var model Model
	err := s.db.QueryRow(context.Background(), "INSERT INTO student(name, group_id) VALUES($1, $2) RETURNING (student_id)", input.Name, input.GroupID).Scan(&model.ID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating studentSerial, student:"+input.Name))
		return nil, err
	}
	model.Name = input.Name
	model.GroupID = input.GroupID
	return &model, nil

}
