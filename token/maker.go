package token

import "time"

// Maker is an interface for managing Tokens
type Maker interface{

	// CreateToken creates new token for given username and duration
	CreateToken(username string,duration time.Duration) (string,error)

	// VerifyToken checks if the input token is valid or not
	VerifyToken(token string)(*Payload,error)
}