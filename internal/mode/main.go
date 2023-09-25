package mode

import (
	"fmt"
	"os"
)

type Mode int

const (
	Serial Mode = iota
	UUID
	PublicKey
)

type mode struct {
	value Mode
}

func (m *mode) String() string {
	switch m.value {
	case Serial:
		return "serial"
	case UUID:
		return "uuid"
	case PublicKey:
		return "public_key"
	default:
		return ""
	}
}

func (m *mode) IsSerial() bool {
	return m.value == Serial
}

func (m *mode) IsUUID() bool {
	return m.value == UUID
}

func (m *mode) IsPublicKey() bool {
	return m.value == PublicKey
}

func New(m Mode) *mode {
	return &mode{
		value: m,
	}
}

func GetMode() Mode {
	m := os.Getenv("MODE")
	fmt.Printf("mode: %v\n", m)
	switch {
	case m == "uuid":
		return UUID
	case m == "public_key":
		return PublicKey
	case m == "serial":
		return Serial
	default:
		panic("Wrong mode passed")
	}

}
