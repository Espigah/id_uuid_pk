package subject

import (
	"context"
	"errors"
)

type subjectPublicKey struct {
	subject
}

func (t *subjectPublicKey) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO subject(title) VALUES($1) RETURNING subject_id, public_key", input.Title).Scan(&model.ID, &model.PublicKey)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating subjectPublicKey, title:"+input.Title))
		return nil, err
	}
	model.Title = input.Title
	return &model, nil

}
