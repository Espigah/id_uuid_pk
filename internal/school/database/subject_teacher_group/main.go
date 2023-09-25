package subjectTeacherGroup

import (
	"fmt"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreateInput struct {
	SubjectID   int64
	SubjectUUID uuid.UUID

	TeacherID   int64
	TeacherUUID uuid.UUID

	GroupID   int64
	GroupUUID uuid.UUID
}

type SubjectTeacherGroup interface {
	Create(CreateInput) (*Model, error)
}

type subjectTeacherGroup struct {
	db        *pgx.Conn
	ID        int64 `json:"subject_teacher_group_id"`
	SubjectID int64 `json:"subject_id"`
	TeacherID int64 `json:"teacher_id"`
	GroupID   int64 `json:"group_id"`
}

type Input struct {
	Mode mode.Mode
	DB   *pgx.Conn
}

func NewFactory(input Input) (func() SubjectTeacherGroup, error) {

	m := mode.New(input.Mode)

	s := subjectTeacherGroup{
		db: input.DB,
	}

	switch {
	case m.IsSerial():
		return func() SubjectTeacherGroup {
			return &subjectTeacherGroupSerial{
				subjectTeacherGroup: s,
			}
		}, nil

	case m.IsUUID():
		return func() SubjectTeacherGroup {
			return &subjectTeacherGroupUUID{
				subjectTeacherGroup: s,
			}
		}, nil

	case m.IsPublicKey():
		return func() SubjectTeacherGroup {
			return &subjectTeacherGroupPublicKey{
				subjectTeacherGroup: s,
			}
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}

}
