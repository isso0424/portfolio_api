package user_error

import (
	"log"

	"github.com/isso0424/portfolio_api/types"
)

type UserError struct {
	message string
}

func New(message string) UserError {
	return UserError{ message }
}

func(err UserError) Type() types.ErrorType {
	return types.User
}

func(err UserError) Error() string {
	log.Printf("Error: %s", err.message)

	return err.message
}
