package types

type ErrorType int
const (
	Internal = iota
	User
)

type APIError interface {
	Error() string
	Type() ErrorType
}
