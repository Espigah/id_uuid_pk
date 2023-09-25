package mark

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type markUUID struct {
	mark
	ID uuid.UUID `json:"id"`
}

func (s *markUUID) Create(input CreateInput) (*Model, error) {
	var model Model
	err := s.db.QueryRow(context.Background(), "INSERT INTO mark(student_id, subject_id, mark) VALUES($1, $2, $3) RETURNING (mark_id)", input.StudentUUID, input.SubjectUUID, input.Mark).Scan(&model.UUID)

	if err != nil {
		err = errors.Join(err, errors.New(fmt.Sprintf("msg:error creating markSerial, studentID:%s, subjectID:%s, mark:%s", input.StudentUUID, input.SubjectUUID, input.Mark)))
		return nil, err
	}
	model.StudentUUID = input.StudentUUID
	model.SubjectUUID = input.SubjectUUID
	model.Mark = input.Mark
	return &model, nil

}
