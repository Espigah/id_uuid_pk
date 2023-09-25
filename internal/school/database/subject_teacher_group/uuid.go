package subjectTeacherGroup

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type subjectTeacherGroupUUID struct {
	subjectTeacherGroup
	ID uuid.UUID `json:"id"`
}

func (t *subjectTeacherGroupUUID) Create(input CreateInput) (*Model, error) {
	var model Model
	err := t.db.QueryRow(
		context.Background(),
		"INSERT INTO subject_teacher_group(subject_id,teacher_id,group_id) VALUES($1,$2,$3) RETURNING (subject_teacher_group_id)", input.SubjectUUID, input.TeacherUUID, input.GroupUUID).
		Scan(&model.UUID)

	if err != nil {
		err = errors.Join(err, errors.New(fmt.Sprintf("msg:error creating subjectTeacherGroupSerial, subjectID:%s, subjectTeacherGroupID:%s, groupID:%s", input.SubjectID, input.TeacherID, input.GroupID)))
		return nil, err

	}

	model.SubjectUUID = input.SubjectUUID
	model.TeacherUUID = input.TeacherUUID
	model.GroupUUID = input.GroupUUID

	return &model, nil
}
