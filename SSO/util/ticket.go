package util

import (
	"fmt"
	"time"
	"crypto/rand"

)
func randToken() []byte {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return b
}
func NewTicket()string{
	ticket := fmt.Sprintf("ST-%v%v%v%x",time.Now().Year(),int(time.Now().Month()),time.Now().Day(),randToken())
	return ticket
}
