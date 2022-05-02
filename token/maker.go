package token

import (
	"time"
)

type Maker interface {
	CreateToken(accountID int32, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}