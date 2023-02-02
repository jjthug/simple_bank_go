package token

import "time"

type Maker interface {

	//creates token for a user for a specific duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Checks if a token is valid
	VerifyToken(token string) (*Payload, error)
}
