package subject

import (
	"fmt"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/jackc/pgx/v5"
)

type CreateInput struct {
	Title     string
	PublicKey string
}

type Subject interface {
	Create(CreateInput) (*Model, error)
}

type subject struct {
	db    *pgx.Conn
	Title string `json:"title"`
}

type Input struct {
	Mode mode.Mode
	DB   *pgx.Conn
}

func NewFactory(input Input) (func() Subject, error) {

	m := mode.New(input.Mode)

	s := subject{
		db: input.DB,
	}

	switch {
	case m.IsSerial():
		return func() Subject {
			return &subjectSerial{
				subject: s,
			}
		}, nil

	case m.IsUUID():
		return func() Subject {
			return &subjectUUID{
				subject: s,
			}
		}, nil

	case m.IsPublicKey():
		return func() Subject {
			return &subjectPublicKey{
				subject: s,
			}
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}

}
