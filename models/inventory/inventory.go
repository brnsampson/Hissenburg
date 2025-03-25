package inventory

import (
	"slices"
	"cmp"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
)

const (
	MAX_BACKPACK = 6
)

type InventoryError interface {
	error
	IsOutOfSpace() bool
}

type OutOfSpace struct {
	message string
}

func newOOS(message string) *OutOfSpace {
	return &OutOfSpace{ message: message }
}

func (oos OutOfSpace) Error() string {
	return oos.message
}

func (oos OutOfSpace) IsOutOfSpace() bool {
	return true
}

type Equipment struct {
	LeftHand *sqlc.ItemView
	RightHand *sqlc.ItemView
	Head *sqlc.ItemView
	Torso *sqlc.ItemView
}

type ItemUpdate struct {
	Left sqlc.ItemView
	Right sqlc.ItemView
}

type EquipmentUpdate struct {
	LeftHand *ItemUpdate
	RightHand *ItemUpdate
	Head *ItemUpdate
	Torso *ItemUpdate
}

func DiffEquipment(left, right Equipment) (onlyLeft, onlyRight Equipment, updates EquipmentUpdate) {
	if left.LeftHand != nil {
		if right.LeftHand != nil {
			if left.LeftHand.ID != right.LeftHand.ID {
				updates.LeftHand = &ItemUpdate{ Left: *left.LeftHand, Right: *right.LeftHand }
			}
		} else {
			// left is something but right is nil
			onlyLeft.LeftHand = left.LeftHand
		}
	} else {
		// right is something but left is nil
		onlyRight.LeftHand = right.LeftHand
	}

	if left.RightHand != nil {
		if right.RightHand != nil {
			if left.RightHand.ID != right.RightHand.ID {
				updates.RightHand = &ItemUpdate{ Left: *left.RightHand, Right: *right.RightHand }
			}
		} else {
			// left is something but right is nil
			onlyLeft.RightHand = left.RightHand
		}
	} else {
		// right is something but left is nil
		onlyRight.RightHand = right.RightHand
	}

	if left.Head != nil {
		if right.Head != nil {
			if left.Head.ID != right.Head.ID {
				updates.Head = &ItemUpdate{ Left: *left.Head, Right: *right.Head }
			}
		} else {
			// left is something but right is nil
			onlyLeft.Head = left.Head
		}
	} else {
		// right is something but left is nil
		onlyRight.Head = right.Head
	}

	if left.Torso != nil {
		if right.Torso != nil {
			if left.Torso.ID != right.Torso.ID {
				updates.Torso = &ItemUpdate{ Left: *left.Torso, Right: *right.Torso }
			}
		} else {
			// left is something but right is nil
			onlyLeft.Torso = left.Torso
		}
	} else {
		// right is something but left is nil
		onlyRight.Torso = right.Torso
	}
	return
}

type ContainerItem struct {
	Index int
	Item sqlc.ItemView
}

type ContainerItemUpdate struct {
	Index int
	Left ContainerItem
	Right ContainerItem
}

type Inventory struct {
	ID int64
	Armor int
	BackpackUsed int
	BonusSpaceUsed int
	BonusSpaceAvailable int
	Equipment Equipment
	Backpack []ContainerItem
	BonusSpace []ContainerItem
	Ground []ContainerItem
}

type InventoryUpdate struct {
	Equipment EquipmentUpdate
	Backpack []ContainerItemUpdate
	BonusSpace []ContainerItemUpdate
	Ground []ContainerItemUpdate
}

func DiffContainer(left, right []ContainerItem) (onlyLeft, onlyRight []ContainerItem, updates []ContainerItemUpdate) {
	lEnd := left[len(left)-1].Index
	rEnd := right[len(right)-1].Index
	longest := max(lEnd, rEnd)

	leftFull :=  make([]*ContainerItem, longest, longest)
	rightFull :=  make([]*ContainerItem, longest, longest)

	for _, ci := range left {
		leftFull[ci.Index] = &ci
	}
	for _, ci := range right {
		rightFull[ci.Index] = &ci
	}

	for i := range longest {
		if leftFull[i] != nil {
			if rightFull[i] != nil {
				if leftFull[i].Item.ID != rightFull[i].Item.ID {
					updates = append(updates, ContainerItemUpdate{ Index: i, Left: *leftFull[i], Right: *rightFull[i] })
				}
			} else {
				// left is something but right is nil
				onlyLeft = append(onlyLeft, *leftFull[i])
			}
		} else {
			// right is something but left is nil
			onlyRight = append(onlyRight, *rightFull[i])
		}
	}
	return
}

//type Inventory struct {
//	ID int64
//	Armor int
//	HandsUsed int
//	BodyUsed int
//	BackpackUsed int
//	ExtraUsed int
//	ExtraSlots int
//	LeftHand item.Item
//	RightHand item.Item
//	Head item.Item
//	Torso item.Item
//	Backpack0 item.Item
//	Backpack1 item.Item
//	Backpack2 item.Item
//	Backpack3 item.Item
//	Backpack4 item.Item
//	Backpack5 item.Item
//	ExtraSpace0 item.Item
//	ExtraSpace1 item.Item
//	ExtraSpace2 item.Item
//	ExtraSpace3 item.Item
//	ExtraSpace4 item.Item
//	ExtraSpace5 item.Item
//	Ground0 item.Item
//	Ground1 item.Item
//	Ground2 item.Item
//	Ground3 item.Item
//	Ground4 item.Item
//	Ground5 item.Item
//}

func New() Inventory {
	equipment := Equipment {
		LeftHand: nil,
		RightHand: nil,
		Head: nil,
		Torso: nil,
	}

	return Inventory {
		ID: -1,
		Armor: 0,
		BackpackUsed: 0,
		BonusSpaceUsed: 0,
		BonusSpaceAvailable: 0,
		Equipment: equipment,
		Backpack: make([]ContainerItem, 0),
		BonusSpace: make([]ContainerItem, 0),
		Ground: make([]ContainerItem, 0),
	}
}

func DiffInventory(left, right Inventory) (onlyLeft, onlyRight Inventory, updates InventoryUpdate) {
	lEquip, rEquip, equipUp := DiffEquipment(left.Equipment, right.Equipment)
	lBackpack, rBackpack, backpackUp := DiffContainer(left.Backpack, right.Backpack)
	lBonus, rBonus, bonusUp := DiffContainer(left.BonusSpace, right.BonusSpace)
	lGround, rGround, groundUp := DiffContainer(left.Ground, right.Ground)

	onlyLeft.ID = -1
	onlyLeft.Equipment = lEquip
	onlyLeft.Backpack = lBackpack
	onlyLeft.BonusSpace = lBonus
	onlyLeft.Ground = lGround

	onlyRight.ID = -1
	onlyRight.Equipment = rEquip
	onlyRight.Backpack = rBackpack
	onlyRight.BonusSpace = rBonus
	onlyRight.Ground = rGround

	updates.Equipment = equipUp
	updates.Backpack = backpackUp
	updates.BonusSpace = bonusUp
	updates.Ground = groundUp

	return
}

//func NewFilled(it sqlc.Item) Inventory {
//	equipment := Equipment {
//		LeftHand: it,
//		RightHand: it,
//		Head: it,
//		Torso: it,
//	}
//
//	backpack := make([]ContainerItem, 0, 10)
//	for i := range 10 {
//		backpack = append(backpack, ContainerItem{ Index: i, Item: it })
//	}
//
//	inv := Inventory {
//		ID: -1,
//		Armor: 0,
//		BackpackUsed: 0,
//		BonusSpaceUsed: 0,
//		BonusSpaceAvailable: 0,
//		Equipment: equipment,
//		Backpack: backpack,
//		BonusSpace: make([]ContainerItem, 0),
//		Ground: make([]ContainerItem, 0),
//	}
//
//	inv.SetCalculatedValues()
//	return inv
//}

func (inv *Inventory) SetCalculatedValues() {
	// 3 armor is the max for cairn
	armor := countArmor(inv, 3)
	backpack := countBackpackUsed(inv)
	bonusUsed := countBonusSpaceUsed(inv)
	bonusSlots := countBonusSpace(inv)
	inv.Armor = armor
	inv.BonusSpaceAvailable = bonusSlots
	inv.BackpackUsed = backpack
	inv.BonusSpaceUsed = bonusUsed
}

func removeAndSort(arr []ContainerItem, idx int) *ContainerItem {
	var to_remove *ContainerItem

	if len(arr) == 0 {
		return nil
	}

	isSorted := slices.IsSortedFunc(arr, func(i, j ContainerItem) int {
		return cmp.Compare(i.Index, j.Index)
	})

	if !isSorted {
		slices.SortFunc(arr, func(i, j ContainerItem) int {
			return cmp.Compare(i.Index, j.Index)
		})
	}

	i, found := slices.BinarySearchFunc(arr, ContainerItem{Index: idx}, func(i, j ContainerItem) int {
		return cmp.Compare(i.Index, j.Index)
	})
	if found == true {
		to_remove = &arr[i]
		arr = append(arr[:i], arr[i+1:]...)
	}

	return to_remove
}

func sortAndReplace(arr []ContainerItem, it ContainerItem) *ContainerItem {
	var to_replace *ContainerItem

	if len(arr) == 0 {
		arr = append(arr, it)
		return to_replace
	}

	isSorted := slices.IsSortedFunc(arr, func(i, j ContainerItem) int {
		return cmp.Compare(i.Index, j.Index)
	})

	if !isSorted {
		slices.SortFunc(arr, func(i, j ContainerItem) int {
			return cmp.Compare(i.Index, j.Index)
		})
	}

	idx, found := slices.BinarySearchFunc(arr, it, func(i, j ContainerItem) int {
		return cmp.Compare(i.Index, j.Index)
	})
	if found == true {
		to_replace = &arr[idx]
		arr[idx] = it
	} else {
		slices.Insert(arr, idx, it)
	}

	return to_replace
}

// AddToBackpack adds the given item to the backpack in sorted order. The Item will be added at inventory index x + 1,
// where x is the highest index currently in the inventory. If that index would exceed the maximum allowed, then an error
// is returned and the item is added to the ground instead.
func (inv *Inventory) AddToBackpack(it sqlc.ItemView, maxSize int) InventoryError {
	if len(inv.Backpack) > 0 {
		curMax := slices.MaxFunc(inv.Backpack, func(i, j ContainerItem) int {
			return cmp.Compare(i.Index, j.Index)
		})
		toAdd := ContainerItem{ Index: curMax.Index + 1, Item: it }
		if toAdd.Index >= maxSize {
			inv.AddToGround(toAdd.Item)
			return newOOS("Exceeded maximum capacity. Added to ground instead.")
		}

		tmp := sortAndReplace(inv.Backpack, toAdd)
		if tmp != nil {
			inv.AddToGround(tmp.Item)
			return newOOS("Somehow replaced an item while appending to backpack. Added to ground instead.")
		}
	} else {
		toAdd := ContainerItem{ Index: 1, Item: it }
		inv.Backpack = append(inv.Backpack, toAdd)
	}

	return nil
}

func (inv *Inventory) RemoveFromBackpack(idx int) *ContainerItem {
	return removeAndSort(inv.Backpack, idx)
}

// AddToBonusSpace adds the given item to the bonus inventory space in sorted order. The Item will be added at index x + 1,
// where x is the highest index currently in the bonus inventory space. If that index would exceed the maximum allowed, then an error
// is returned and the item is added to the ground instead.
func (inv *Inventory) AddToBonusSpace(it sqlc.ItemView) InventoryError {
	maxSize := countBonusSpace(inv)
	curMax := slices.MaxFunc(inv.BonusSpace, func(i, j ContainerItem) int {
		return cmp.Compare(i.Index, j.Index)
	})
	toAdd := ContainerItem{ Index: curMax.Index + 1, Item: it }
	if toAdd.Index >= maxSize {
		inv.AddToGround(toAdd.Item)
		return newOOS("Exceeded maximum capacity. Added to ground instead.")
	}

	tmp := sortAndReplace(inv.BonusSpace, toAdd)
	if tmp != nil {
		inv.AddToGround(tmp.Item)
		return newOOS("Somehow replaced an item while appending to bonus inventory space. Added to ground instead.")
	}

	return nil
}

func (inv *Inventory) RemoveFromBonusSpace(idx int) *ContainerItem {
	return removeAndSort(inv.BonusSpace, idx)
}

// AddToGround adds the given item to the backpack in sorted order. The Item will be added at ground index x + 1,
// where x is the highest index currently in the ground.
func (inv *Inventory) AddToGround(it sqlc.ItemView) *ContainerItem {
	curMax := slices.MaxFunc(inv.Ground, func(i, j ContainerItem) int {
		return cmp.Compare(i.Index, j.Index)
	})

	return sortAndReplace(inv.Ground, ContainerItem{ Index: curMax.Index + 1, Item: it})
}

func (inv *Inventory) RemoveFromGround(idx int) *ContainerItem {
	return removeAndSort(inv.Ground, idx)
}

// InsertIntoBackpack adds the given item to the backpack in sorted order. If an item is added with the same index as another
// existing item in the backpack, then the old item will be replaced and a pointer to it returned.
func (inv *Inventory) InsertIntoBackpack(it ContainerItem) *ContainerItem {
	return sortAndReplace(inv.Backpack, it)
}

// InsertIntoBonusSpace adds the given item to the bonus space in sorted order. If an item is added with the same index as another
// existing item in the bonusspace, then the old item will be replaced and a pointer to it returned.
func (inv *Inventory) InsertIntoBonusSpace(it ContainerItem) *ContainerItem {
	return sortAndReplace(inv.BonusSpace, it)
}

// InsertIntoGround adds the given item to the bonus space in sorted order. If an item is added with the same index as another
// existing item in the ground, then the old item will be replaced and a pointer to it returned.
func (inv *Inventory) InsertIntoGround(it ContainerItem) *ContainerItem {
	return sortAndReplace(inv.Ground, it)
}

func countArmor(inv *Inventory, maxArmor int) int {
	armor := int64(0)
	if inv.Equipment.LeftHand.Slot == "hands" {
		armor += inv.Equipment.LeftHand.Armor
	}

	if inv.Equipment.RightHand.Slot == "hands" {
		armor += inv.Equipment.RightHand.Armor
	}

	if inv.Equipment.Head.Slot == "head" || inv.Equipment.Head.Slot == "body" {
		armor += inv.Equipment.Head.Armor
	}

	if inv.Equipment.Torso.Slot == "torso" || inv.Equipment.Head.Slot == "body" {
		armor += inv.Equipment.Torso.Armor
	}

	for _, it := range inv.Backpack {
		if it.Item.Slot == "backpack" {
			armor += it.Item.Armor
		}
	}

	for _, it := range inv.BonusSpace {
		if it.Item.Slot == "bonus" {
			armor += it.Item.Armor
		}
	}

	for _, it := range inv.Ground {
		if it.Item.Slot == "ground" {
			armor += it.Item.Armor
		}
	}

	return min(int(armor), maxArmor)
}

func leftHandSize(inv *Inventory) int {
	if inv.Equipment.LeftHand.Slot == "hand" {
		return int(inv.Equipment.LeftHand.ActiveSize)
	} else {
		return int(inv.Equipment.LeftHand.Size)
	}
}

func rightHandSize(inv *Inventory) int {
	if inv.Equipment.RightHand.Slot == "hand" {
		return int(inv.Equipment.RightHand.ActiveSize)
	} else {
		return int(inv.Equipment.RightHand.Size)
	}
}

func countHandSize(inv *Inventory) int {
	return leftHandSize(inv) + rightHandSize(inv)
}

func countHeadSize(inv *Inventory) int {
	if inv.Equipment.Head.Slot == "head" || inv.Equipment.Head.Slot == "body" {
		return int(inv.Equipment.Head.ActiveSize)
	} else {
		return int(inv.Equipment.Head.Size)
	}
}

func countTorsoSize(inv *Inventory) int {
	if inv.Equipment.Torso.Slot == "torso" || inv.Equipment.Torso.Slot == "body" {
		return int(inv.Equipment.Torso.ActiveSize)
	} else {
		return int(inv.Equipment.Torso.Size)
	}
}

func countBodySize(inv *Inventory) int {
	return countHeadSize(inv) + countTorsoSize(inv)
}

func countEquipmentSize(inv *Inventory) int {
	return leftHandSize(inv) + rightHandSize(inv) + countHeadSize(inv) + countTorsoSize(inv)
}

func countBackpackUsed(inv *Inventory) int {
	used := int64(0)

	for _, it := range inv.Backpack {
		if it.Item.Slot == "backpack" {
			used += it.Item.ActiveSize
		} else {
			used += it.Item.Size
		}
	}

	return int(used)
}

func countBonusSpaceUsed(inv *Inventory) int {
	used := int64(0)

	for _, it := range inv.BonusSpace {
		if it.Item.Slot == "bonus" || it.Item.Slot == "backpack" {
			used += it.Item.ActiveSize
		} else {
			used += it.Item.Size
		}
	}

	return int(used)
}

func countBonusSpace(inv *Inventory) int {
	bonus := int64(0)
	if inv.Equipment.LeftHand.Slot == "hand" {
		bonus += inv.Equipment.LeftHand.Storage
	}

	if inv.Equipment.RightHand.Slot == "hand" {
		bonus += inv.Equipment.RightHand.Storage
	}

	if inv.Equipment.Head.Slot == "head" || inv.Equipment.Head.Slot == "body" {
		bonus += inv.Equipment.Head.Storage
	}

	if inv.Equipment.Torso.Slot == "torso" || inv.Equipment.Head.Slot == "body" {
		bonus += inv.Equipment.Torso.Storage
	}

	for _, it := range inv.Backpack {
		if it.Item.Slot == "backpack" {
			bonus += it.Item.Storage
		}
	}

	return int(bonus)
}

func countGround(inv *Inventory) int {
	return len(inv.Ground)
}

func (inv *Inventory) DropOverweight() {
	hands := countHandSize(inv)
	body := countBodySize(inv)
	backpack := countBackpackUsed(inv)
	bonusUsed := countBonusSpaceUsed(inv)
	bonusSpace:= countBonusSpace(inv)

	if hands > 2 {
		if leftHandSize(inv) > rightHandSize(inv) {
			inv.AddToGround(*inv.Equipment.RightHand)
			inv.Equipment.RightHand = nil
		} else {
			inv.AddToGround(*inv.Equipment.LeftHand)
			inv.Equipment.LeftHand = nil
		}
	}

	if body > 2 {
		if countHeadSize(inv) > countTorsoSize(inv) {
			inv.AddToGround(*inv.Equipment.Torso)
			inv.Equipment.Torso = nil
		} else {
			inv.AddToGround(*inv.Equipment.Head)
			inv.Equipment.Head = nil
		}
	}

	if backpack > 6 {
		to_remove := make([]int, 0)
		seen := 0
		for _, it := range inv.Backpack {
			if seen > 6 {
				inv.AddToGround(it.Item)
				to_remove = append(to_remove, it.Index)
			} else {
				if it.Item.Slot == "backpack" {
					seen += int(it.Item.ActiveSize)
				} else {
					seen += int(it.Item.Size)
				}
			}
		}
		for _, idx := range to_remove {
			inv.RemoveFromBackpack(idx)
		}
	}

	bonusSpace = countBonusSpace(inv)
	if bonusUsed > bonusSpace {
		to_remove := make([]int, 0)
		seen := 0
		for _, it := range inv.BonusSpace {
			if seen > 6 {
				inv.AddToGround(it.Item)
				to_remove = append(to_remove, it.Index)
			} else {
				if it.Item.Slot == "bonus" || it.Item.Slot == "backpack" {
					seen += int(it.Item.ActiveSize)
				} else {
					seen += int(it.Item.Size)
				}
			}
		}
		for _, idx := range to_remove {
			inv.RemoveFromBonusSpace(idx)
		}
	}

	return
}
//
//// TruncateExtra will remove items from the extra list and drop them on the ground until the number of slots used by
//// the extra list can fit into the available extra slot capacity. Returns true if changes have been made and false if
//// no changes were made. This is the case regardless of whether an error is returned.
//func (inv *Inventory) TruncateExtra(empty item.Item) (changed bool, err error) {
//	changed = false
//	extra := countExtraUsed(inv)
//	extraSlots := countExtraSlots(inv)
//	if extra >= extraSlots {
//		return
//	}
//
//	currCount := 0
//
//	extras := []*item.Item{ &inv.ExtraSpace0, &inv.ExtraSpace1, &inv.ExtraSpace2, &inv.ExtraSpace3, &inv.ExtraSpace4, &inv.ExtraSpace5 }
//	for _, e := range extras {
//		nextSize := e.CurrentSize("Backpack")
//		if extraSlots < currCount + nextSize {
//			// Would overencumber the inventory
//			err = inv.AddToGround(*e)
//			*e = empty
//			changed = true
//		} else {
//			currCount += nextSize
//		}
//	}
//
//	return
//}

//func FromModel(ctx context.Context, repo *data.Repo, inv sqlc.Inventory) (Inventory, error) {
//	empty, err := repo.GetItemFromName(ctx, "Empty")
//	if err != nil {
//		return Inventory{}, fmt.Errorf("Could not find Empty item")
//	}
//
//	left := empty
//	right := empty
//	head := empty
//	torso := empty
//	b1 := empty
//	e1 := empty
//	g1 := empty
//	b2 := empty
//	e2 := empty
//	g2 := empty
//	b3 := empty
//	e3 := empty
//	g3 := empty
//	b4 := empty
//	e4 := empty
//	g4 := empty
//	b5 := empty
//	e5 := empty
//	g5 := empty
//	b6 := empty
//	e6 := empty
//	g6 := empty
//
//	if inv.LeftHand.Valid {
//		left = inv.LeftHand.Int64
//	}
//
//	mv := Inventory {
//		ID: inv.InventoryID,
//
//	}
//}

//type InventoryEditView struct {
//	Inventory
//	Name string
//	Surname string
//	ItemTypeList []item.ItemType
//}
//
//func (v InventoryView) IntoEditView(name, surname string, itemTypeList []item.ItemType) InventoryEditView {
//	return InventoryEditView{ v, name, surname, itemTypeList }
//}

//type Inventory struct {
//	Head item.Item
//	Torso item.Item
//	LeftHand item.Item
//	RightHand item.Item
//	Backpack []item.Item
//	ExtraSpace []item.Item
//	Ground []item.Item
//}
//
//func New() Inventory {
//	backpack := make([]item.Item, 0)
//	extra := make([]item.Item, 0)
//	ground := make([]item.Item, 0)
//
//	return Inventory{Head: item.Empty(), Torso: item.Empty(), LeftHand: item.Empty(), RightHand: item.Empty(), Backpack: backpack, ExtraSpace: extra, Ground: ground}
//}
//
//// Unpack returns an Inventory struct by unpacking a single slice of items in the specific order of
//// Head
//// Torso
//// LeftHand
//// RightHand
//// Backpack (padded with empty items to get to exactly 6 item size)
//// ExtraSlots (padded with empty items to get to exactly inv.ExtraSpace() item size)
//// Ground
//func Unpack(flat []item.Item) (Inventory, error) {
//	inv := New()
//	inv.Head = flat[0]
//	inv.Torso = flat[1]
//	inv.LeftHand = flat[2]
//	inv.RightHand = flat[3]
//	flat = flat[3:]
//
//	added := 0
//	addedSize := 0
//	for _, it := range flat {
//		if addedSize + it.Size >= 6 {
//			break
//		}
//		addedSize += it.Size
//		added += 1
//		err := inv.AddToBackpack(it)
//		if err != nil {
//			return inv, err
//		}
//	}
//
//	flat = flat[added:]
//	extra := inv.ExtraSpaceAvailable()
//	added = 0
//	addedSize = 0
//	for _, it := range flat {
//		if addedSize + it.Size >= extra {
//			break
//		}
//		addedSize += it.Size
//		added += 1
//		err := inv.AddToExtraSpace(it)
//		if err != nil {
//			return inv, err
//		}
//	}
//
//	flat = flat[added:]
//	for _, it := range flat {
//		inv.AddToGround(it)
//	}
//
//	return inv, nil
//}
//
//// Pack returns a single slice of items in the specific order of
//// Head
//// Torso
//// LeftHand
//// RightHand
//// Backpack (padded with empty items to get to exactly 6 item size)
//// ExtraSlots (padded with empty items to get to exactly inv.ExtraSpace() item size)
//// Ground
//func (inv Inventory) Pack() []item.Item {
//	flat := make([]item.Item, 0)
//	flat = append(flat, inv.Head)
//	flat = append(flat, inv.Torso)
//	flat = append(flat, inv.LeftHand)
//	flat = append(flat, inv.RightHand)
//
//	backUsed := inv.BackpackUsed()
//	emptyNeeded := 6 - backUsed
//
//	for _, it := range inv.Backpack {
//		flat = append(flat, it)
//	}
//
//	for _ = range emptyNeeded {
//		flat = append(flat, item.Empty())
//	}
//
//	extraUsed := inv.ExtraSpaceUsed()
//	emptyNeeded = inv.ExtraSpaceAvailable() - extraUsed
//	for _, it := range inv.ExtraSpace {
//		flat = append(flat, it)
//	}
//
//	for _ = range emptyNeeded {
//		flat = append(flat, item.Empty())
//	}
//
//	for _, it := range inv.Ground {
//		flat = append(flat, it)
//	}
//	return flat
//}
//
//func (inv Inventory) IntoView() InventoryView {
//	head := inv.Head
//	torso := inv.Torso
//	left := inv.LeftHand
//	right := inv.RightHand
//	armor := inv.ArmorEquipped()
//	back := make([]item.Item, 0)
//	slotsUsed := inv.BackpackUsed()
//
//	for _, i := range inv.Backpack {
//		back = append(back, i)
//	}
//
//	for i := 0; i <= (6 - slotsUsed); i++ {
//		back = append(back, item.Empty())
//	}
//
//	return InventoryView{ Head: head, Torso: torso, LeftHand: left, RightHand: right, Armor: armor, Backpack: back, BackpackUsed: slotsUsed, Ground: inv.Ground }
//}
//
//func (inv *Inventory) BackpackUsed() int {
//	count := 0
//	for _, i := range inv.Backpack {
//		if i.Type != item.Bulk && i.Type != item.EmptySlot {
//			count += i.Size
//		}
//	}
//	return count
//}
//
//func (inv *Inventory) BackpackFull() bool {
//	if len(inv.Backpack) > 5 {
//		return true
//	}
//
//	filled := inv.BackpackUsed()
//	if filled < 6 {
//		return false
//	}
//
//	// How did we get here? Dunno, but if someone is trying to game the system it should work against them
//	return true
//}
//
//func (inv Inventory) ExtraSpaceAvailable() int {
//	// TODO: avoid adding to extra if the item type is EmptySlot or Bulk
//	extra := 0
//	if inv.Head.ActiveSlot == item.Head {
//		extra += inv.Head.Storage
//	}
//
//	if inv.Head.ActiveSlot == item.Head {
//		extra += inv.Head.Storage
//	}
//
//	if inv.Head.ActiveSlot == item.Head {
//		extra += inv.Head.Storage
//	}
//
//	if inv.Head.ActiveSlot == item.Head {
//		extra += inv.Head.Storage
//	}
//
//	for _, it := range inv.Backpack {
//		if it.ActiveSlot == item.Backpack {
//			extra += it.Storage
//		}
//	}
//
//	return extra
//}
//
//func (inv *Inventory) ExtraSpaceUsed() int {
//	count := 0
//	for _, i := range inv.ExtraSpace {
//		if i.Type != item.Bulk && i.Type != item.EmptySlot {
//			count += i.Size
//		}
//	}
//	return count
//}
//
//func (inv Inventory) ArmorEquipped() int {
//	armor := 0
//	if inv.Head.ActiveSlot == item.Head {
//		armor += inv.Head.Armor
//	}
//
//	if inv.Torso.ActiveSlot == item.Torso {
//		armor += inv.Torso.Armor
//	}
//
//	if inv.LeftHand.ActiveSlot == item.Hand {
//		armor += inv.LeftHand.Armor
//	}
//
//	if inv.RightHand.ActiveSlot == item.Hand {
//		armor += inv.RightHand.Armor
//	}
//
//	for _, it := range inv.Backpack {
//		if it.ActiveSlot == item.Backpack {
//			armor += it.Armor
//		}
//	}
//
//	return armor
//}
//
//func (inv *Inventory) AddActive(i item.Item) error {
//	toPlace := []item.Item{}
//	if i.ActiveSlot == item.Hand {
//		if i.Size == 1 {
//			if inv.RightHand.Type == item.EmptySlot {
//				toPlace = append(toPlace, inv.SetRightHand(i))
//			} else {
//				toPlace = append(toPlace, inv.SetLeftHand(i))
//			}
//		} else if i.Size == 2 {
//			tmp := inv.SetHands(i, item.MakeBulk(i))
//			toPlace = append(toPlace, tmp[0])
//			toPlace = append(toPlace, tmp[1])
//		}
//	} else if i.ActiveSlot == item.Head {
//		if i.Size == 1 {
//			toPlace = append(toPlace, inv.SetHead(i))
//		} else if i.Size == 2 {
//			tmp := inv.SetBody(i, item.MakeBulk(i))
//			toPlace = append(toPlace, tmp[0])
//			toPlace = append(toPlace, tmp[1])
//		}
//	} else if i.ActiveSlot == item.Torso {
//		if i.Size == 1 {
//			toPlace = append(toPlace, inv.SetTorso(i))
//		} else if i.Size == 2 {
//			tmp := inv.SetBody(item.MakeBulk(i), i)
//			toPlace = append(toPlace, tmp[0])
//			toPlace = append(toPlace, tmp[1])
//		}
//	} else {
//		// Backpack or Any slots
//		toPlace = append(toPlace, i)
//	}
//
//	for _, thing := range toPlace {
//		err := inv.AddToBackpack(thing)
//		if err != nil {
//			return err
//		}
//	}
//
//	return  nil
//}
//
//func (inv *Inventory) SetLeftHand(i item.Item) item.Item {
//	old := inv.LeftHand
//	inv.LeftHand = i
//	return old
//}
//
//func (inv *Inventory) SetRightHand(i item.Item) item.Item {
//	old := inv.RightHand
//	inv.RightHand = i
//	return old
//}
//
//func (inv *Inventory) SetHands(left, right item.Item) [2]item.Item {
//	var old [2]item.Item
//	old[0] = inv.LeftHand
//	old[1] = inv.RightHand
//
//	inv.LeftHand = left
//	inv.RightHand = right
//
//	return old
//}
//
//func (inv *Inventory) SetHead(i item.Item) item.Item {
//	old := inv.Head
//	inv.Head = i
//	return old
//}
//
//func (inv *Inventory) SetTorso(i item.Item) item.Item {
//	old := inv.Torso
//	inv.Torso = i
//	return old
//}
//
//func (inv *Inventory) SetBody(left, right item.Item) [2]item.Item {
//	var old [2]item.Item
//	old[0] = inv.Head
//	old[1] = inv.Torso
//
//	inv.Head = left
//	inv.Torso = right
//
//	return old
//}
//
//func (inv *Inventory) SetBackpack(idx int, i item.Item) (old item.Item, err error) {
//	// Sanity checks
//	if idx < 0 || idx > 5 {
//		return item.Empty(), fmt.Errorf("SetBackpack index out of bounds. Must be < 0 and > 5")
//	}
//
//	if i.Type == item.EmptySlot {
//		// just deleting the existing item
//		l := len(inv.Backpack) - 1
//		if idx > l {
//			return item.Empty(), fmt.Errorf("SetBackpack index out of bounds. Must be <= %d", l)
//		}
//		old = inv.Backpack[idx]
//		inv.Backpack = append(inv.Backpack[:idx], inv.Backpack[idx + 1:]...)
//		return old, nil
//	}
//
//	filled := inv.BackpackUsed()
//	after := filled
//	if i.Type != item.EmptySlot && i.Type != item.Bulk {
//		after += i.Size
//	}
//
//	// If we explicitly want to add an item later in the backpack...
//	l := len(inv.Backpack) - 1
//	for _ = range (idx - l) {
//		inv.Backpack = append(inv.Backpack, item.Empty())
//	}
//
//	old = inv.Backpack[idx]
//	if i.Type != item.EmptySlot && i.Type != item.Bulk {
//		after -= old.Size
//	}
//
//	if after >= 6 {
//		return item.Empty(), fmt.Errorf("SetBackpack failed: Would result in too many backpack slots being filled")
//	}
//
//	inv.Backpack[idx] = i
//	return old, nil
//}
//
//// AddToBackpack attemts to just add the given item to the backpack in a dumb way, and if it can't it will return an error
//// Also, we just throw away any empty slots...
//func (inv *Inventory) AddToBackpack(i item.Item) error {
//	if i.Type == item.EmptySlot || i.Type == item.Bulk {
//		return nil
//	}
//
//	filled := inv.BackpackUsed()
//	if (filled + i.Size) > 6 {
//		return fmt.Errorf("Item %s cannot fit in backpack", i.Name)
//	}
//
//	// pick the first empty slot if possible
//	foundit := false
//	idx := 0
//	for j, other := range inv.Backpack {
//		idx = j
//		if other.Type == item.EmptySlot {
//			foundit = true
//			break
//		}
//	}
//
//	if foundit {
//		inv.Backpack[idx] = i
//	} else {
//		inv.Backpack = append(inv.Backpack, i)
//	}
//
//	return nil
//}
//
//// AddToExtraSlots attemts to just add the given item to the backpack in a dumb way, and if it can't it will return an error
//// Also, we just throw away any empty slots...
//func (inv *Inventory) AddToExtraSpace(i item.Item) error {
//	if i.Type == item.EmptySlot || i.Type == item.Bulk {
//		return nil
//	}
//
//	filled := inv.ExtraSpaceUsed()
//	if (filled + i.Size) > 6 {
//		return fmt.Errorf("Item %s cannot fit in extra storage space", i.Name)
//	}
//
//	// pick the first empty slot if possible
//	foundit := false
//	idx := 0
//	for j, other := range inv.ExtraSpace {
//		idx = j
//		if other.Type == item.EmptySlot {
//			foundit = true
//			break
//		}
//	}
//
//	if foundit {
//		inv.ExtraSpace[idx] = i
//	} else {
//		inv.ExtraSpace = append(inv.ExtraSpace, i)
//	}
//
//	return nil
//}
//
//func (inv *Inventory) AddToGround(i item.Item) {
//	if i.Type == item.EmptySlot || i.Type == item.Bulk {
//		return
//	}
//	inv.Ground = append(inv.Ground, i)
//	return
//}
//
//func (inv *Inventory) AddToStorage(i item.Item) {
//	err := inv.AddToBackpack(i)
//	if err != nil {
//		err = inv.AddToExtraSpace(i)
//		if err != nil {
//			inv.AddToGround(i)
//		}
//	}
//	return
//}
//
//func (inv *Inventory) DeleteGround(idx int) (old item.Item, err error) {
//	l := len(inv.Ground)
//	if idx < 0 || idx >= l {
//		return item.Empty(), fmt.Errorf("Ground item %d does not exist", idx)
//	}
//
//	old = inv.Ground[idx]
//	inv.Ground = append(inv.Ground[:idx], inv.Ground[idx+1:]...)
//	return
//}
