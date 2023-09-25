package group

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type groupUUID struct {
	group
	ID uuid.UUID `json:"id"`
}

func (t *groupUUID) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO \"group\"(name) VALUES($1) RETURNING (group_id)", input.Name).Scan(&model.UUID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating groupSerial, name:"+input.Name))
		return nil, err
	}

	model.Name = input.Name
	return &model, nil
}
