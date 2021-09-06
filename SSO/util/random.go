package util

import (
	"crypto/rand"
	"fmt"

	"github.com/labstack/gommon/random"
)

func randToken() []byte {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return b
}

func NewTGT() string {
	return fmt.Sprintf("TGT-%s", random.String(16))
}

func NewTicket() string {
	ticket := fmt.Sprintf("ST-%s", random.String(16))
	return ticket
}

func NewSMSCode() string {
	return random.String(6, random.Numeric)
}
