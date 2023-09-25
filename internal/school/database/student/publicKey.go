package student

import (
	"context"
	"errors"
)

type studentPublicKey struct {
	student
}

func (t *studentPublicKey) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO student(name, group_id) VALUES($1, $2) RETURNING student_id, public_key", input.Name, input.GroupID).Scan(&model.ID, &model.PublicKey)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating studentSerial, title:"+input.Name))
		return nil, err
	}
	model.Name = input.Name
	return &model, nil

}
