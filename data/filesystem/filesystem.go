package filesystem

import (
	"os"
	"io/fs"
	"path/filepath"
	"encoding/json"
	"github.com/brnsampson/Hissenburg/models/character"
	"github.com/charmbracelet/log"
)

func readCharFile(path string) (*character.Character, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var char character.Character
	err = json.Unmarshal(data, &char)
	if err != nil {
		return nil, err
	}
	return &char, nil
}

func newCharWalkFunc() (*[]*character.Character, fs.WalkDirFunc, error) {
	c := make([]*character.Character, 0)
	f := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			char, err := readCharFile(path)
			if err != nil {
				return err
			}
			c = append(c, char)
		}
		return nil
	}
	return &c, f, nil
}


type FileBackend struct {
	BasePath string
}

func New(path string) (FileBackend, error) {
	log.Debug("Created new FileBackend", "path", path)
	p, err := filepath.Abs(path)
	fb := FileBackend{ BasePath: p }
	return fb, err
}

func (fb FileBackend) UpdateCharacter(char *character.Character) error {
	dirPath := filepath.Join(fb.BasePath, char.Surname)
	err := os.MkdirAll(dirPath, 0750)
	if err != nil {
		return err
	}

	path := filepath.Join(fb.BasePath, char.Surname, char.Name + ".json")
	log.Debug("Updating character file", "path", path)
	enc, err := json.Marshal(char)
	if err != nil {
		return err
	}
	os.WriteFile(path, enc, 0666)
	return nil
}

func (fb FileBackend) ListCharacters() ([]*character.Character, error) {
	chars, f, err := newCharWalkFunc()
	if err != nil {
		return *chars, err
	}
	err = filepath.WalkDir(fb.BasePath, f)
	if err != nil {
		return *chars, err
	}
	return *chars, err
}

func (fb FileBackend) GetCharacter(name, surname string) (*character.Character, error) {
	path := filepath.Join(fb.BasePath, surname, name + ".json")
	log.Debug("Reading character file", "path", path)
	return readCharFile(path)
}

func (fb FileBackend) DeleteCharacter(name, surname string) error {
	path := filepath.Join(fb.BasePath, surname, name + ".json")
	log.Debug("Deleting character file", "path", path)

	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
