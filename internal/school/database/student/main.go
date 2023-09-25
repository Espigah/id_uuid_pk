package student

import (
	"fmt"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreateInput struct {
	Name      string
	GroupID   int64
	GroupUUID uuid.UUID
	PublicKey string
}

type Student interface {
	Create(CreateInput) (*Model, error)
}

type student struct {
	db    *pgx.Conn
	Title string `json:"title"`
}

type Input struct {
	Mode mode.Mode
	DB   *pgx.Conn
}

func NewFactory(input Input) (func() Student, error) {

	m := mode.New(input.Mode)

	s := student{
		db: input.DB,
	}

	switch {
	case m.IsSerial():
		return func() Student {
			return &studentSerial{
				student: s,
			}
		}, nil

	case m.IsUUID():
		return func() Student {
			return &studentUUID{
				student: s,
			}
		}, nil

	case m.IsPublicKey():
		return func() Student {
			return &studentPublicKey{
				student: s,
			}
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}

}
