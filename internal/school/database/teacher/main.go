package teacher

import (
	"fmt"

	"github.com/Espigah/uuid_vs_serial_vs_public_key/internal/mode"
	"github.com/jackc/pgx/v5"
	"github.com/rs/xid"
)

type CreateInput struct {
	Name      string
	PublicKey xid.ID
}

type Teacher interface {
	Create(CreateInput) (*Model, error)
	GetTeacherMark(*Model) (interface{}, error)
	GetLatest() (*Model, error)
}

type teacher struct {
	db *pgx.Conn
}

type Input struct {
	Mode mode.Mode
	DB   *pgx.Conn
}

func NewFactory(input Input) (func() Teacher, error) {

	m := mode.New(input.Mode)

	t := teacher{
		db: input.DB,
	}

	switch {
	case m.IsSerial():
		return func() Teacher {
			return &teacherSerial{
				teacher: t,
			}
		}, nil

	case m.IsUUID():
		return func() Teacher {
			return &teacherUUID{
				teacher: t,
			}
		}, nil

	case m.IsPublicKey():
		return func() Teacher {
			return &teacherPublicKey{
				teacher: t,
			}
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}

}
