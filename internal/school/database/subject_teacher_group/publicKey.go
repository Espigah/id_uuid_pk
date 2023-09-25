package subjectTeacherGroup

import (
	"context"
	"errors"
	"fmt"
)

type subjectTeacherGroupPublicKey struct {
	subjectTeacherGroup
}

func (t *subjectTeacherGroupPublicKey) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(
		context.Background(),
		"INSERT INTO subject_teacher_group(subject_id,teacher_id,group_id) VALUES($1,$2,$3) RETURNING (subject_teacher_group_id)", input.SubjectID, input.TeacherID, input.GroupID).
		Scan(&model.ID)

	if err != nil {
		err = errors.Join(err, errors.New(fmt.Sprintf("msg:error creating subjectTeacherGroupTeacherGroupSerial, subjectTeacherGroupID:%s, subjectTeacherGroupTeacherGroupID:%s, groupID:%s", input.SubjectID, input.TeacherID, input.GroupID)))
		return nil, err
	}

	model.SubjectID = input.SubjectID
	model.TeacherID = input.TeacherID
	model.GroupID = input.GroupID

	return &model, nil
}
