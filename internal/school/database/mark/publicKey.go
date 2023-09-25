package mark

import (
	"context"
	"errors"
	"fmt"
)

type markPublicKey struct {
	mark
}

func (s *markPublicKey) Create(input CreateInput) (*Model, error) {
	var model Model
	err := s.db.QueryRow(context.Background(),
		"INSERT INTO mark(student_id, subject_id, mark) VALUES($1, $2, $3) RETURNING mark_id, public_key", input.StudentID, input.SubjectID, input.Mark).
		Scan(&model.ID, &model.PublicKey)

	if err != nil {
		err = errors.Join(err, errors.New(fmt.Sprintf("msg:error creating markSerial, studentID:%s, subjectID:%s, mark:%s", input.StudentID, input.SubjectID, input.Mark)))
		return nil, err
	}
	model.StudentID = input.StudentID
	model.SubjectID = input.SubjectID
	model.Mark = input.Mark
	return &model, nil

}
