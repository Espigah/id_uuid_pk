package mark

import (
	"context"
	"errors"
	"fmt"
)

type markSerial struct {
	mark
	ID int64 `json:"mark_id"`
}

func (s *markSerial) Create(input CreateInput) (*Model, error) {
	var model Model
	err := s.db.QueryRow(context.Background(), "INSERT INTO mark(student_id, subject_id, mark) VALUES($1, $2, $3) RETURNING (mark_id)", input.StudentID, input.SubjectID, input.Mark).Scan(&model.ID)

	if err != nil {
		err = errors.Join(err, errors.New(fmt.Sprintf("msg:error creating markSerial, studentID:%s, subjectID:%s, mark:%s", input.StudentID, input.SubjectID, input.Mark)))
		return nil, err
	}
	model.StudentID = input.StudentID
	model.SubjectID = input.SubjectID
	model.Mark = input.Mark
	return &model, nil

}
