package domain

type Validator interface {
	Schema() ([]byte, error)
}
