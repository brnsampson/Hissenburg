package data

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
)

type ItemRepo interface {
	ListItemKinds(ctx context.Context) ([]sqlc.ItemKind, error)
	ListItemSlots(ctx context.Context) ([]sqlc.ItemSlot, error)
	GetItemKind(ctx context.Context, id int64) (sqlc.ItemKind, error)
	GetItemSlot(ctx context.Context, id int64) (sqlc.ItemSlot, error)
	GetItemKindFromString(ctx context.Context, kind string) (sqlc.ItemKind, error)
	GetItemSlotFromString(ctx context.Context, slot string) (sqlc.ItemSlot, error)
	CreateItemKind(ctx context.Context, name string) (sqlc.ItemKind, error)
	CreateItemSlot(ctx context.Context, name string) (sqlc.ItemSlot, error)
	ListItems(ctx context.Context) ([]sqlc.ItemView, error)
	GetItem(ctx context.Context, id int64) (sqlc.Item, error)
	GetItemView(ctx context.Context, id int64) (sqlc.ItemView, error)
	GetItemFromKindAndName(ctx context.Context, kind, name string) (sqlc.ItemView, error)
	GetRandomItemFromKind(ctx context.Context, kind string) (sqlc.ItemView, error)
	CreateItem(ctx context.Context, it sqlc.ItemView) (sqlc.Item, error)
	UpdateItem(ctx context.Context, it sqlc.ItemView) error
	DeleteItem(ctx context.Context, id int64) error
	DeleteItemKind(ctx context.Context, id int64) error
	DeleteItemSlot(ctx context.Context, id int64) error
}

//func (r *Repository) ListItems(ctx context.Context) ([]item.Item, error) {
//	parsed := make([]item.Item, 0)
//	its, err := r.Queries.ListItems(ctx)
//	if err != nil {
//		return parsed, err
//	}
//
//	for _, it := range its {
//		kind, err := r.GetItemKind(ctx, it.Kind)
//		if err != nil {
//			return parsed, err
//		}
//
//		slot, err := r.GetItemSlot(ctx, it.Slot)
//		if err != nil {
//			return parsed, err
//		}
//
//		vm := item.Item {
//			Kind:        kind,
//			Name:        it.Name,
//			Description: it.Description,
//			Value:       it.Value,
//			DiceCount:   0,
//			DiceSides:   0,
//			Armor:       it.Armor,
//			Storage:     it.Storage,
//			Size:        it.Size,
//			ActiveSize:  it.ActiveSize,
//			Slot:        slot,
//			Stackable:   it.Stackable,
//			Icon:        it.Icon,
//		}
//
//		if it.DiceCount < 0 {
//			return parsed, fmt.Errorf("Invalid Item: DiceCount must be >= 0, was %d", it.DiceCount)
//		}
//		dc := uint64(it.DiceCount)
//		vm.DiceCount = dc
//
//		if it.DiceSides < 0 {
//			return parsed, fmt.Errorf("Invalid Item: DiceSides must be >= 0, was %d", it.DiceSides)
//		}
//		ds := uint64(it.DiceSides)
//		vm.DiceSides = ds
//
//		parsed = append(parsed, vm)
//	}
//
//	return parsed, nil
//}

func (r *Repository) GetItemFromKindAndName(ctx context.Context, kind, name string) (sqlc.ItemView, error) {
	queryParams := sqlc.GetItemFromKindAndNameParams { Kind: kind, Name: name }
	return r.Queries.GetItemFromKindAndName(ctx, queryParams)
}

//func (r *Repository) GetItem(ctx context.Context, id int64) (item.Item, error) {
//	it, err := r.Queries.GetItem(ctx, id)
//	if err != nil {
//		return item.Item{}, err
//	}
//
//	return ItemDBToModel(ctx, r.Queries, it)
//}

//func (r *Repository) GetRandomItemFromKind(ctx context.Context, kind string) (item.Item, error) {
//	it, err := r.Queries.GetRandomItemFromKind(ctx, kind)
//	if err != nil {
//		return item.Item{}, err
//	}
//
//	return ItemDBToModel(ctx, r.Queries, it)
//}

func (r *Repository) CreateItem(ctx context.Context, it sqlc.ItemView) (sqlc.Item, error) {
	kind, err := r.GetItemKindFromString(ctx, it.Kind)
	if err != nil {
		log.Error("Error looking up ItemKind", "it", it.Name, "kind", it.Kind)
		return sqlc.Item{}, err
	}
	slot, err := r.GetItemSlotFromString(ctx, it.Slot)
	if err != nil {
		log.Error("Error looking up ItemSlot", "it", it.Name, "slot", it.Slot)
		return sqlc.Item{}, err
	}

	itemParam := sqlc.CreateItemParams {
		Name: it.Name,
		Kind: kind.ID,
		Slot: slot.ID,
		Description: it.Description,
		Value: it.Value,
		DiceCount: it.DiceCount,
		DiceSides: it.DiceSides,
		Armor: it.Armor,
		Storage: it.Storage,
		Size: it.Size,
		ActiveSize: it.ActiveSize,
		Stackable: it.Stackable,
		Icon: it.Icon,
	}

	return r.Queries.CreateItem(ctx, itemParam)
}

func (r *Repository) UpdateItem(ctx context.Context, it sqlc.ItemView) error {
	oldItem, err := r.Queries.GetItem(ctx, it.ID)
	if err != nil {
		log.Error("Error looking up Item to update", "item", it.ID, "name", it.Name, "kind", it.Kind, "slot", it.Slot)
		return err
	}

	kind, err := r.GetItemKindFromString(ctx, it.Kind)
	if err != nil {
		log.Error("Error looking up ItemKind for Item update", "item", it.ID, "name", it.Name, "kind", it.Kind, "slot", it.Slot)
		return err
	}
	slot, err := r.GetItemSlotFromString(ctx, it.Slot)
	if err != nil {
		log.Error("Error looking up ItemSlot for Item update", "item", it.ID, "name", it.Name, "kind", it.Kind, "slot", it.Slot)
		return err
	}

	itemParam := sqlc.UpdateItemParams {
		Name: it.Name,
		Kind: kind.ID,
		Slot: slot.ID,
		Description: it.Description,
		Value: it.Value,
		DiceCount: int64(it.DiceCount),
		DiceSides: int64(it.DiceSides),
		Armor: it.Armor,
		Storage: it.Storage,
		Size: it.Size,
		ActiveSize: it.ActiveSize,
		Stackable: it.Stackable,
		Icon: it.Icon,
		ID: oldItem.ID,
	}

	return r.Queries.UpdateItem(ctx, itemParam)
}

