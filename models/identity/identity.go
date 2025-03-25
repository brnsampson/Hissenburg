package identity

import "github.com/brnsampson/Hissenburg/gen/sqlc"

type Identity struct {
	ID         int64
	Name       string
	Surname    string
	Age        int64
	Portrait   string
	Gender     sqlc.Gender
	Background sqlc.Background
}

func New() Identity {
	return Identity{ ID: -1, Gender: sqlc.Gender{}, Background: sqlc.Background{} }
}
