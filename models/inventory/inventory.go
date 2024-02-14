package inventory

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/models/item"
)

const (
	MAX_BACKPACK = 6
)

type InventoryView struct {
	Head item.Item
	Torso item.Item
	LeftHand item.Item
	RightHand item.Item
	Armor int
	Backpack []item.Item
	BackpackUsed int
	Ground []item.Item
}

type InventoryEditView struct {
	InventoryView
	Name string
	Surname string
	ItemTypeList []item.ItemType
}

func (v InventoryView) IntoEditView(name, surname string, itemTypeList []item.ItemType) InventoryEditView {
	return InventoryEditView{ v, name, surname, itemTypeList }
}

type Inventory struct {
	Head item.Item
	Torso item.Item
	LeftHand item.Item
	RightHand item.Item
	Backpack []item.Item
	ExtraSpace []item.Item
	Ground []item.Item
}

func New() Inventory {
	backpack := make([]item.Item, 0)
	extra := make([]item.Item, 0)
	ground := make([]item.Item, 0)

	return Inventory{Head: item.Empty(), Torso: item.Empty(), LeftHand: item.Empty(), RightHand: item.Empty(), Backpack: backpack, ExtraSpace: extra, Ground: ground}
}

// Unpack returns an Inventory struct by unpacking a single slice of items in the specific order of
// Head
// Torso
// LeftHand
// RightHand
// Backpack (padded with empty items to get to exactly 6 item size)
// ExtraSlots (padded with empty items to get to exactly inv.ExtraSpace() item size)
// Ground
func Unpack(flat []item.Item) (Inventory, error) {
	inv := New()
	inv.Head = flat[0]
	inv.Torso = flat[1]
	inv.LeftHand = flat[2]
	inv.RightHand = flat[3]
	flat = flat[3:]

	added := 0
	addedSize := 0
	for _, it := range flat {
		if addedSize + it.Size >= 6 {
			break
		}
		addedSize += it.Size
		added += 1
		err := inv.AddToBackpack(it)
		if err != nil {
			return inv, err
		}
	}

	flat = flat[added:]
	extra := inv.ExtraSpaceAvailable()
	added = 0
	addedSize = 0
	for _, it := range flat {
		if addedSize + it.Size >= extra {
			break
		}
		addedSize += it.Size
		added += 1
		err := inv.AddToExtraSpace(it)
		if err != nil {
			return inv, err
		}
	}

	flat = flat[added:]
	for _, it := range flat {
		inv.AddToGround(it)
	}

	return inv, nil
}

// Pack returns a single slice of items in the specific order of
// Head
// Torso
// LeftHand
// RightHand
// Backpack (padded with empty items to get to exactly 6 item size)
// ExtraSlots (padded with empty items to get to exactly inv.ExtraSpace() item size)
// Ground
func (inv Inventory) Pack() []item.Item {
	flat := make([]item.Item, 0)
	flat = append(flat, inv.Head)
	flat = append(flat, inv.Torso)
	flat = append(flat, inv.LeftHand)
	flat = append(flat, inv.RightHand)

	backUsed := inv.BackpackUsed()
	emptyNeeded := 6 - backUsed

	for _, it := range inv.Backpack {
		flat = append(flat, it)
	}

	for _ = range emptyNeeded {
		flat = append(flat, item.Empty())
	}

	extraUsed := inv.ExtraSpaceUsed()
	emptyNeeded = inv.ExtraSpaceAvailable() - extraUsed
	for _, it := range inv.ExtraSpace {
		flat = append(flat, it)
	}

	for _ = range emptyNeeded {
		flat = append(flat, item.Empty())
	}

	for _, it := range inv.Ground {
		flat = append(flat, it)
	}
	return flat
}

func (inv Inventory) IntoView() InventoryView {
	head := inv.Head
	torso := inv.Torso
	left := inv.LeftHand
	right := inv.RightHand
	armor := inv.ArmorEquipped()
	back := make([]item.Item, 0)
	slotsUsed := inv.BackpackUsed()

	for _, i := range inv.Backpack {
		back = append(back, i)
	}

	for i := 0; i <= (6 - slotsUsed); i++ {
		back = append(back, item.Empty())
	}

	return InventoryView{ Head: head, Torso: torso, LeftHand: left, RightHand: right, Armor: armor, Backpack: back, BackpackUsed: slotsUsed, Ground: inv.Ground }
}

func (inv *Inventory) BackpackUsed() int {
	count := 0
	for _, i := range inv.Backpack {
		if i.Type != item.Bulk && i.Type != item.EmptySlot {
			count += i.Size
		}
	}
	return count
}

func (inv *Inventory) BackpackFull() bool {
	if len(inv.Backpack) > 5 {
		return true
	}

	filled := inv.BackpackUsed()
	if filled < 6 {
		return false
	}

	// How did we get here? Dunno, but if someone is trying to game the system it should work against them
	return true
}

func (inv Inventory) ExtraSpaceAvailable() int {
	// TODO: avoid adding to extra if the item type is EmptySlot or Bulk
	extra := 0
	if inv.Head.ActiveSlot == item.HeadSlot {
		extra += inv.Head.Storage
	}

	if inv.Head.ActiveSlot == item.HeadSlot {
		extra += inv.Head.Storage
	}

	if inv.Head.ActiveSlot == item.HeadSlot {
		extra += inv.Head.Storage
	}

	if inv.Head.ActiveSlot == item.HeadSlot {
		extra += inv.Head.Storage
	}

	for _, it := range inv.Backpack {
		if it.ActiveSlot == item.BackpackSlot {
			extra += it.Storage
		}
	}

	return extra
}

func (inv *Inventory) ExtraSpaceUsed() int {
	count := 0
	for _, i := range inv.ExtraSpace {
		if i.Type != item.Bulk && i.Type != item.EmptySlot {
			count += i.Size
		}
	}
	return count
}

func (inv Inventory) ArmorEquipped() int {
	armor := 0
	if inv.Head.ActiveSlot == item.HeadSlot {
		armor += inv.Head.Armor
	}

	if inv.Torso.ActiveSlot == item.TorsoSlot {
		armor += inv.Torso.Armor
	}

	if inv.LeftHand.ActiveSlot == item.HandSlot {
		armor += inv.LeftHand.Armor
	}

	if inv.RightHand.ActiveSlot == item.HandSlot {
		armor += inv.RightHand.Armor
	}

	for _, it := range inv.Backpack {
		if it.ActiveSlot == item.BackpackSlot {
			armor += it.Armor
		}
	}

	return armor
}

func (inv *Inventory) AddActive(i item.Item) error {
	toPlace := []item.Item{}
	if i.ActiveSlot == item.HandSlot {
		if i.Size == 1 {
			if inv.RightHand.Type == item.EmptySlot {
				toPlace = append(toPlace, inv.SetRightHand(i))
			} else {
				toPlace = append(toPlace, inv.SetLeftHand(i))
			}
		} else if i.Size == 2 {
			tmp := inv.SetHands(i, item.MakeBulk(i))
			toPlace = append(toPlace, tmp[0])
			toPlace = append(toPlace, tmp[1])
		}
	} else if i.ActiveSlot == item.HeadSlot {
		if i.Size == 1 {
			toPlace = append(toPlace, inv.SetHead(i))
		} else if i.Size == 2 {
			tmp := inv.SetBody(i, item.MakeBulk(i))
			toPlace = append(toPlace, tmp[0])
			toPlace = append(toPlace, tmp[1])
		}
	} else if i.ActiveSlot == item.TorsoSlot {
		if i.Size == 1 {
			toPlace = append(toPlace, inv.SetTorso(i))
		} else if i.Size == 2 {
			tmp := inv.SetBody(item.MakeBulk(i), i)
			toPlace = append(toPlace, tmp[0])
			toPlace = append(toPlace, tmp[1])
		}
	} else {
		// Backpack or Any slots
		toPlace = append(toPlace, i)
	}

	for _, thing := range toPlace {
		err := inv.AddToBackpack(thing)
		if err != nil {
			return err
		}
	}

	return  nil
}

func (inv *Inventory) SetLeftHand(i item.Item) item.Item {
	old := inv.LeftHand
	inv.LeftHand = i
	return old
}

func (inv *Inventory) SetRightHand(i item.Item) item.Item {
	old := inv.RightHand
	inv.RightHand = i
	return old
}

func (inv *Inventory) SetHands(left, right item.Item) [2]item.Item {
	var old [2]item.Item
	old[0] = inv.LeftHand
	old[1] = inv.RightHand

	inv.LeftHand = left
	inv.RightHand = right

	return old
}

func (inv *Inventory) SetHead(i item.Item) item.Item {
	old := inv.Head
	inv.Head = i
	return old
}

func (inv *Inventory) SetTorso(i item.Item) item.Item {
	old := inv.Torso
	inv.Torso = i
	return old
}

func (inv *Inventory) SetBody(left, right item.Item) [2]item.Item {
	var old [2]item.Item
	old[0] = inv.Head
	old[1] = inv.Torso

	inv.Head = left
	inv.Torso = right

	return old
}

func (inv *Inventory) SetBackpack(idx int, i item.Item) (old item.Item, err error) {
	// Sanity checks
	if idx < 0 || idx > 5 {
		return item.Empty(), fmt.Errorf("SetBackpack index out of bounds. Must be < 0 and > 5")
	}

	if i.Type == item.EmptySlot {
		// just deleting the existing item
		l := len(inv.Backpack) - 1
		if idx > l {
			return item.Empty(), fmt.Errorf("SetBackpack index out of bounds. Must be <= %d", l)
		}
		old = inv.Backpack[idx]
		inv.Backpack = append(inv.Backpack[:idx], inv.Backpack[idx + 1:]...)
		return old, nil
	}

	filled := inv.BackpackUsed()
	after := filled
	if i.Type != item.EmptySlot && i.Type != item.Bulk {
		after += i.Size
	}

	// If we explicitly want to add an item later in the backpack...
	l := len(inv.Backpack) - 1
	for _ = range (idx - l) {
		inv.Backpack = append(inv.Backpack, item.Empty())
	}

	old = inv.Backpack[idx]
	if i.Type != item.EmptySlot && i.Type != item.Bulk {
		after -= old.Size
	}

	if after >= 6 {
		return item.Empty(), fmt.Errorf("SetBackpack failed: Would result in too many backpack slots being filled")
	}

	inv.Backpack[idx] = i
	return old, nil
}

// AddToBackpack attemts to just add the given item to the backpack in a dumb way, and if it can't it will return an error
// Also, we just throw away any empty slots...
func (inv *Inventory) AddToBackpack(i item.Item) error {
	if i.Type == item.EmptySlot || i.Type == item.Bulk {
		return nil
	}

	filled := inv.BackpackUsed()
	if (filled + i.Size) > 6 {
		return fmt.Errorf("Item %s cannot fit in backpack", i.Name)
	}

	// pick the first empty slot if possible
	foundit := false
	idx := 0
	for j, other := range inv.Backpack {
		idx = j
		if other.Type == item.EmptySlot {
			foundit = true
			break
		}
	}

	if foundit {
		inv.Backpack[idx] = i
	} else {
		inv.Backpack = append(inv.Backpack, i)
	}

	return nil
}

// AddToExtraSlots attemts to just add the given item to the backpack in a dumb way, and if it can't it will return an error
// Also, we just throw away any empty slots...
func (inv *Inventory) AddToExtraSpace(i item.Item) error {
	if i.Type == item.EmptySlot || i.Type == item.Bulk {
		return nil
	}

	filled := inv.ExtraSpaceUsed()
	if (filled + i.Size) > 6 {
		return fmt.Errorf("Item %s cannot fit in extra storage space", i.Name)
	}

	// pick the first empty slot if possible
	foundit := false
	idx := 0
	for j, other := range inv.ExtraSpace {
		idx = j
		if other.Type == item.EmptySlot {
			foundit = true
			break
		}
	}

	if foundit {
		inv.ExtraSpace[idx] = i
	} else {
		inv.ExtraSpace = append(inv.ExtraSpace, i)
	}

	return nil
}

func (inv *Inventory) AddToGround(i item.Item) {
	if i.Type == item.EmptySlot || i.Type == item.Bulk {
		return
	}
	inv.Ground = append(inv.Ground, i)
	return
}

func (inv *Inventory) AddToStorage(i item.Item) {
	err := inv.AddToBackpack(i)
	if err != nil {
		err = inv.AddToExtraSpace(i)
		if err != nil {
			inv.AddToGround(i)
		}
	}
	return
}

func (inv *Inventory) DeleteGround(idx int) (old item.Item, err error) {
	l := len(inv.Ground)
	if idx < 0 || idx >= l {
		return item.Empty(), fmt.Errorf("Ground item %d does not exist", idx)
	}

	old = inv.Ground[idx]
	inv.Ground = append(inv.Ground[:idx], inv.Ground[idx+1:]...)
	return
}
