package mark

import (
	"fmt"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreateInput struct {
	PublicKey   string
	StudentID   int64
	SubjectID   int64
	StudentUUID uuid.UUID
	SubjectUUID uuid.UUID
	Mark        int64
}

type Mark interface {
	Create(CreateInput) (*Model, error)
}

type mark struct {
	db    *pgx.Conn
	Title string `json:"title"`
}

type Input struct {
	Mode mode.Mode
	DB   *pgx.Conn
}

func NewFactory(input Input) (func() Mark, error) {

	m := mode.New(input.Mode)

	mk := mark{
		db: input.DB,
	}

	switch {
	case m.IsSerial():
		return func() Mark {
			return &markSerial{
				mark: mk,
			}
		}, nil

	case m.IsUUID():
		return func() Mark {
			return &markUUID{
				mark: mk,
			}
		}, nil

	case m.IsPublicKey():
		return func() Mark {
			return &markPublicKey{
				mark: mk,
			}
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}

}
