package mark

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

type Model struct {
	ID          int64     `json:"mark_id"`
	UUID        uuid.UUID `json:"uuid"`
	PublicKey   xid.ID    `json:"public_key"`
	StudentID   int64
	SubjectID   int64
	StudentUUID uuid.UUID
	SubjectUUID uuid.UUID
	Mark        int64 `json:"mark"`
}
