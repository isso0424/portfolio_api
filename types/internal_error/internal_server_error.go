package internal_error

import (
	"log"

	"github.com/isso0424/portfolio_api/types"
)

type InternalError struct {
	message string
}

func New(message string) InternalError {
	return InternalError{ message }
}

func(err InternalError) Type() types.ErrorType {
	return types.Internal
}

func(err InternalError) Error() string {
	log.Printf("Error: %s", err.message)

	return "internal server error"
}
