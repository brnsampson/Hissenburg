package main

import (
	"os"
	"path/filepath"
	"encoding/json"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/charmbracelet/log"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/brnsampson/Hissenburg/models/item"
)

const STATIC_PATH = "./static"

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

func run() error {
	log.SetLevel(log.DebugLevel)
	log.Debug("started DB population script")
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3")
	if err != nil {
		return err
	}
	log.Debug("Opened DB connection")

	queries := sqlc.New(db)

	// Populate database
	itemList := make([]item.Item, 0)
	err = readJson("./items/core.json", &itemList)
	if err != nil {
		return err
	}

	for _, it := range itemList {
		log.Debug("Attempting to create item if it does not exist", "item", it)
		found, err := queries.GetItemFromName(ctx, it.Name)
		if err != nil {
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

			itemParam := sqlc.CreateItemParams {
				//Kind: kind.Name,
				Kind: it.Kind,
				Name: it.Name,
				Description: it.Description,
				Value: it.Value,
				DiceCount: int64(it.DiceCount),
				DiceSides: int64(it.DiceSides),
				Armor: it.Armor,
				Storage: it.Storage,
				Size: it.Size,
				ActiveSize: it.ActiveSize,
				//Slot: slot.ItemSlotID,
				Slot: it.Slot,
				Stackable: it.Stackable,
				Icon: it.Icon,
			}

			newItem, err := queries.CreateItem(ctx, itemParam)
			if err != nil {
				log.Error("Error Creating item", "item", it.Name, "kind", it.Kind)
				return err
			}
			log.Info("Created new item", "item", newItem)
		} else {
			log.Debug("Found item to already exist", "item", it.Name, "record", found)
		}
	}

	log.Debug("Attempting to fetch item 'Boondoggle'")
	boon, err := queries.GetItemFromName(ctx, "Boondoggle")
	if err != nil {
		log.Debug("Boondoggle did not exist and err WAS NOT nil! WORKING AS INTENDED!", "error", err)
	} else {
		log.Error("Boondoggle should not exist and err WAS nil! THAT IS A PROBLEM!", "boondoggle", boon)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

