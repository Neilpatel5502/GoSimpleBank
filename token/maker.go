package token

import "time"

// Interface for managing tokens
type Maker interface {
	// creates a new token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// Checks if input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
