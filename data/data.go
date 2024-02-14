package data

import (
	"github.com/brnsampson/Hissenburg/data/filesystem"
	"github.com/brnsampson/Hissenburg/data/static"
	"github.com/brnsampson/Hissenburg/models/character"
	"github.com/brnsampson/Hissenburg/models/inventory"
	"github.com/brnsampson/Hissenburg/models/item"
	"github.com/brnsampson/Hissenburg/models/status"
	"github.com/brnsampson/Hissenburg/models/trait"
)

type CharBackend interface {
	ListCharacters() ([]*character.Character, error)
	GetCharacter(name, surname string) (*character.Character, error)
	CreateCharacter(char *character.Character) (*character.Character, error)
	UpdateCharacter(name, surname string, char *character.Character) (*character.Character, error)
	DeleteCharacter(name, surname string) (*character.Character, error)
	UpdateInventory(name, surname string, inv inventory.Inventory) error
	UpdateStatus(name, surname string, status status.Status) error

	ListGender(filter string) []character.Gender
	PickGender() character.Gender
	ListName(gender character.Gender, filter string) []string
	PickName(character.Gender) string
	ListSurname(filter string) []string
	PickSurname() string
	ListBackground(filter string) []string
	PickBackground() string
	ListTrait(kind trait.TraitType, filter string) []string
	PickTrait(trait.TraitType) string
}

type ItemBackend interface {
	ListAllItems(filter string) []item.Item
	ListItem(kind item.ItemType, filter string) []item.Item
	GetItem(kind item.ItemType, name string) (item.Item, error)
	PickItem(kind item.ItemType) (item.Item, error)
}

type ItemData struct {
	static.StaticBackend
}

func NewItemData() (ItemData, error) {
	return ItemData{ static.New() }, nil
}

type CharData struct {
	static.StaticBackend
	fileBackend filesystem.FileBackend
}

func NewCharData(path string) (*CharData, error) {
	fb, err := filesystem.New(path)
	if err != nil {
		return nil, err
	}
	return &CharData{ static.New(), fb }, nil
}

func (b *CharData) ListCharacters() ([]*character.Character, error) {
	return b.fileBackend.ListCharacters()
}
func (b *CharData) GetCharacter(name, surname string) (*character.Character, error) {
	return b.fileBackend.GetCharacter(name, surname)
}
func (b *CharData) CreateCharacter(char *character.Character) (*character.Character, error) {
	err := b.fileBackend.UpdateCharacter(char)
	if err != nil {
		return nil, err
	}

	return b.fileBackend.GetCharacter(char.Name, char.Surname)
}

func (b *CharData) UpdateCharacter(name, surname string, char *character.Character) (*character.Character, error) {
	removeOld := false
	if name != char.Name || surname != char.Surname {
		removeOld = true
	}
	err := b.fileBackend.UpdateCharacter(char)
	if err != nil {
		return nil, err
	}

	if removeOld {
		err := b.fileBackend.DeleteCharacter(name, surname)
		if err != nil {
			return nil, err
		}
	}
	return b.fileBackend.GetCharacter(char.Name, char.Surname)
}
func (b *CharData) DeleteCharacter(name, surname string) (*character.Character, error) {
	char, err :=  b.fileBackend.GetCharacter(name, surname)
	if err != nil {
		return nil, err
	}

	err = b.fileBackend.DeleteCharacter(name, surname)
	if err != nil {
		return nil, err
	}
	return char, nil
}

func (b *CharData) UpdateInventory(name, surname string, inv inventory.Inventory) error {
	char, err := b.fileBackend.GetCharacter(name, surname)
	if err != nil {
		return err
	}

	char.Inventory = inv
	err = b.fileBackend.UpdateCharacter(char)
	if err != nil {
		return err
	}

	return nil
}

func (b *CharData) UpdateStatus(name, surname string, status status.Status) error {
	char, err := b.fileBackend.GetCharacter(name, surname)
	if err != nil {
		return err
	}

	char.Status = status
	err = b.fileBackend.UpdateCharacter(char)
	if err != nil {
		return err
	}

	return nil
}
