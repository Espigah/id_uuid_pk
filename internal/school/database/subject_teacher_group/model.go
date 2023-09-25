package subjectTeacherGroup

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

type Model struct {
	ID          int64     `json:"subject_teacher_group_id"`
	UUID        uuid.UUID `json:"uuid"`
	PublicKey   xid.ID    `json:"public_key"`
	SubjectID   int64     `json:"subject_id"`
	SubjectUUID uuid.UUID
	TeacherID   int64 `json:"teacher_id"`
	TeacherUUID uuid.UUID
	GroupID     int64 `json:"group_id"`
	GroupUUID   uuid.UUID
}
