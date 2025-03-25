package character

import (
	"github.com/brnsampson/Hissenburg/gen/sqlc"
)

type CharacterListView struct {
	Characters []Character
}

type Associations struct {
	CharacterID int64
	User        sqlc.User
	Party       sqlc.Party
	Village     sqlc.Village
}

type AssociationsEdit struct {
	Name string
	Surname string
	Associations Associations
}

type Status struct {
	CharacterID int64
	Hp          int64
	MaxHp       int64
	Str         int64
	MaxStr      int64
	Dex         int64
	MaxDex      int64
	Will        int64
	MaxWill     int64
}

type StatusEdit struct {
	Name string
	Surname string
	Status Status
}

type Traits struct {
	CharacterID int64
	Physique    string
	Skin        string
	Hair        string
	Face        string
	Speech      string
	Clothing    string
	Virtue      string
	Vice        string
	Reputation  string
	Misfortune  string
}

type TraitsEdit struct {
	Name string
	Surname string
	Traits Traits
}

type Identity struct {
	CharacterID int64
	Gender      string
	Name        string
	Surname     string
	Age         int64
	Portrait    string
	Background  sqlc.Background
}

type IdentityEdit struct {
	Identity Identity
}

type Character struct {
	ID           int64
	Description  string
	Associations Associations
	Identity     Identity
	Traits       Traits
	Status       Status
	Inventory    int64
}

func New() Character {
	return Character{}
}
