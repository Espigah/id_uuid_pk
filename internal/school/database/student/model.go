package student

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

type Model struct {
	ID        int64     `json:"student_id"`
	UUID      uuid.UUID `json:"uuid"`
	PublicKey xid.ID    `json:"public_key"`
	Name      string    `json:"name"`
	GroupID   int64     `json:"group_id"`
	GroupUUID uuid.UUID
}
