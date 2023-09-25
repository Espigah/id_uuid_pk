package group

import (
	"context"
	"errors"
)

type groupSerial struct {
	group
	ID int64 `json:"group_id"`
}

func (t *groupSerial) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO \"group\"(name) VALUES($1) RETURNING (group_id)", input.Name).Scan(&model.ID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating groupSerial, name:"+input.Name))
		return nil, err
	}

	model.Name = input.Name
	return &model, nil
}
