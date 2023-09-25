package subject

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type subjectUUID struct {
	subject
	ID uuid.UUID `json:"id"`
}

func (t *subjectUUID) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO subject(title) VALUES($1) RETURNING (subject_id)", input.Title).Scan(&model.UUID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating subjectUUID, title:"+input.Title))
		return nil, err
	}
	model.Title = input.Title
	return &model, nil

}
