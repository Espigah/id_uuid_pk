package student

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type studentUUID struct {
	student
	ID uuid.UUID `json:"id"`
}

func (s *studentUUID) Create(input CreateInput) (*Model, error) {
	var model Model
	err := s.db.QueryRow(context.Background(), "INSERT INTO student(name, group_id) VALUES($1, $2) RETURNING (student_id)", input.Name, input.GroupUUID).Scan(&model.UUID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating studentSerial, student:"+input.Name))
		return nil, err
	}
	model.Name = input.Name
	model.GroupUUID = input.GroupUUID
	return &model, nil

}
