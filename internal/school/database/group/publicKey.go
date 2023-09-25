package group

import (
	"context"
	"errors"
)

type groupPublicKey struct {
	group
}

func (t *groupPublicKey) Create(input CreateInput) (*Model, error) {
	var model Model

	err := t.db.QueryRow(context.Background(), "INSERT INTO \"group\"(name) VALUES($1) RETURNING group_id, public_key", input.Name).Scan(&model.ID, &model.PublicKey)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating groupSerial, title:"+input.Name))
		return nil, err
	}

	model.Name = input.Name
	return &model, nil

}
