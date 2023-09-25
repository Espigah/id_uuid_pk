package subject

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

type Model struct {
	ID        int64     `json:"subject_id"`
	UUID      uuid.UUID `json:"uuid"`
	PublicKey xid.ID    `json:"public_key"`
	Title     string    `json:"title"`
}
