package startup

import (
	"os"
	"path/filepath"
	"encoding/json"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/charmbracelet/log"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/brnsampson/Hissenburg/data"
)

const STATIC_PATH = "./static"

var DEFAULT_ITEM_LISTS = []string{"core"}

func readJson[T any](path string, s *T) error {
	p := filepath.Join(STATIC_PATH, path)

	log.Debug("Decoding file", "path", p)

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}

func populateItemList(ctx context.Context, repo data.ItemRepo, listPath string) error {
	itemList := make([]sqlc.ItemView, 0)
	err := readJson(listPath, &itemList)
	if err != nil {
		return err
	}

	for _, it := range itemList {
		log.Debug("Attempting to create item if it does not exist", "item", it)
		found, err := repo.GetItemFromKindAndName(ctx, it.Kind, it.Name)
		if err != nil {
			log.Debug("Did not find item already created", "error", err)
			//kind, err := queries.GetItemKindFromString(ctx, tmp.Kind)
			//if err != nil {
			//	log.Error("Error looking up ItemKind", "item", tmp.Name, "kind", tmp.Kind)
			//	return err
			//}
			//slot, err := queries.GetItemSlotFromString(ctx, tmp.Slot)
			//if err != nil {
			//	log.Error("Error looking up ItemSlot", "item", tmp.Name, "slot", tmp.Slot)
			//	return err
			//}

			//itemParam := sqlc.ItemView {
			//	Name: it.Name,
			//	Kind: it.Kind,
			//	Slot: it.Slot,
			//	Description: it.Description,
			//	Value: it.Value,
			//	DiceCount: int64(it.DiceCount),
			//	DiceSides: int64(it.DiceSides),
			//	Armor: it.Armor,
			//	Storage: it.Storage,
			//	Size: it.Size,
			//	ActiveSize: it.ActiveSize,
			//	Stackable: it.Stackable,
			//	Icon: it.Icon,
			//}

			newItem, err := repo.CreateItem(ctx, it)
			if err != nil {
				log.Error("Error Creating item", "item", it.Name, "kind", it.Kind)
				return err
			}
			log.Info("Created new item", "item", newItem)
		} else {
			log.Debug("Found item to already exist", "item", it.Name, "record", found)
		}
	}
	return nil
}

func PopulateMissingItems(itemLists []string) error {
	ctx := context.Background()
	itemLists = append(itemLists, DEFAULT_ITEM_LISTS...)
	log.Debug("Starting populating items from static lists", "item_lists", itemLists)

	db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3")
	if err != nil {
		return err
	}
	repo := data.NewRepo(db)

	log.Debug("Opened DB connection")

	// Populate database
	for _, l := range itemLists {
		path := filepath.Join("./", "items", l + ".json")
		err := populateItemList(ctx, repo, path)
		if err != nil {
			return err
		}
	}

	return nil
}
