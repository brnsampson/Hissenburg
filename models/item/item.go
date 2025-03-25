package item

import "github.com/brnsampson/Hissenburg/gen/sqlc"

type ItemOverviewView struct {
	Kinds []sqlc.ItemKind
	Slots []sqlc.ItemSlot
	Items []sqlc.ItemView
}

type ItemEditView struct {
	Kinds []sqlc.ItemKind
	Slots []sqlc.ItemSlot
	Item sqlc.ItemView
}

//type Item struct {
//	Kind        string
//	Name        string
//	Description string
//	Value       int64
//	DiceCount   uint64
//	DiceSides   uint64
//	Armor       int64
//	Storage     int64
//	Size        int64
//	ActiveSize  int64
//	Slot        string
//	Stackable   bool
//	Icon        string
//}
//
//func New() Item {
//	return Item {}
//}
//
//func (i Item) CurrentSize(slot string) int {
//	if i.Slot == slot {
//		return int(i.ActiveSize)
//	} else {
//		return int(i.Size)
//	}
//}
