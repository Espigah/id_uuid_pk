package teacher

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

type Model struct {
	ID        int64     `json:"teacher_id"`
	UUID      uuid.UUID `json:"uuid"`
	PublicKey xid.ID    `json:"public_key"`
	Name      string    `json:"name"`
}

type TeacherMark struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Mark    int64  `json:"mark"`
}
