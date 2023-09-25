package subject

import (
	"context"
	"errors"
)

type subjectSerial struct {
	subject
	ID int64 `json:"subject_id"`
}

func (t *subjectSerial) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(context.Background(), "INSERT INTO subject(title) VALUES($1) RETURNING (subject_id)", input.Title).Scan(&model.ID)

	if err != nil {
		err = errors.Join(err, errors.New("msg:error creating subjectSerial, title:"+input.Title))
		return nil, err
	}
	model.Title = input.Title
	return &model, nil

}
