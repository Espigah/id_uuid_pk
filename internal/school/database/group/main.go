package group

import (
	"fmt"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/jackc/pgx/v5"
)

type CreateInput struct {
	Name      string
	PublicKey string
}

type Group interface {
	Create(CreateInput) (*Model, error)
}

type group struct {
	db   *pgx.Conn
	Name string `json:"name"`
}

type Input struct {
	Type mode.Mode
	DB   *pgx.Conn
}

func NewFactory(input Input) (func() Group, error) {

	m := mode.New(input.Type)

	g := group{
		db: input.DB,
	}

	switch {
	case m.IsSerial():
		return func() Group {
			return &groupSerial{
				group: g,
			}
		}, nil

	case m.IsUUID():
		return func() Group {
			return &groupUUID{
				group: g,
			}
		}, nil

	case m.IsPublicKey():
		return func() Group {
			return &groupPublicKey{
				group: g,
			}
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}

}
