package data

import (
	//"fmt"
	"context"
	"github.com/brnsampson/Hissenburg/models/character"
	//"github.com/brnsampson/Hissenburg/models/item"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
	"github.com/charmbracelet/log"
)

//func ItemDBToModel(ctx context.Context, q *sqlc.Queries, it sqlc.Item) (item.Item, error) {
//	tmp := item.Item {
//		Kind:        it.Kind,
//		Name:        it.Name,
//		Description: it.Description,
//		Value:       it.Value,
//		DiceCount:   0,
//		DiceSides:   0,
//		Armor:       it.Armor,
//		Storage:     it.Storage,
//		Size:        it.Size,
//		ActiveSize:  it.ActiveSize,
//		Slot:        it.Slot,
//		Stackable:   it.Stackable,
//		Icon:        it.Icon,
//	}
//
//	if it.DiceCount < 0 {
//		return tmp, fmt.Errorf("Invalid Item: DiceCount must be >= 0, was %d", it.DiceCount)
//	}
//	dc := uint64(it.DiceCount)
//	tmp.DiceCount = dc
//
//	if it.DiceSides < 0 {
//		return tmp, fmt.Errorf("Invalid Item: DiceSides must be >= 0, was %d", it.DiceSides)
//	}
//	ds := uint64(it.DiceSides)
//	tmp.DiceSides = ds
//
//	return tmp, nil
//}

//func GetItemRowToDB(it sqlc.GetItemRow) sqlc.Item {
//	return sqlc.Item{
//		ItemID:     it.ItemID,
//		Name:        it.Name,
//		Kind:        it.Kind,
//		Description: it.Description,
//		Value:       it.Value,
//		DiceCount:   it.DiceCount,
//		DiceSides:   it.DiceSides,
//		Armor:       it.Armor,
//		Storage:     it.Storage,
//		Size:        it.Size,
//		ActiveSize:  it.ActiveSize,
//		Slot:        it.Slot,
//		Stackable:   it.Stackable,
//		Icon:        it.Icon,
//	}
//}

//func GetItemFromKindAndNameRowToDB(it sqlc.GetItemFromKindAndNameRow) sqlc.Item {
//	return sqlc.Item{
//		ItemID:     it.ItemID,
//		Name:        it.Name,
//		Kind:        it.Kind,
//		Description: it.Description,
//		Value:       it.Value,
//		DiceCount:   it.DiceCount,
//		DiceSides:   it.DiceSides,
//		Armor:       it.Armor,
//		Storage:     it.Storage,
//		Size:        it.Size,
//		ActiveSize:  it.ActiveSize,
//		Slot:        it.Slot,
//		Stackable:   it.Stackable,
//		Icon:        it.Icon,
//	}
//}

//func GetItemRowToModel(ctx context.Context, q *sqlc.Queries, it sqlc.GetItemRow) (item.Item, error) {
//	tmp := GetItemRowToDB(it)
//	return ItemDBToModel(ctx, q, tmp)
//}
//
//func GetItemFromNameRowToModel(ctx context.Context, q *sqlc.Queries, it sqlc.GetItemFromNameRow) (item.Item, error) {
//	tmp := GetItemFromNameRowToDB(it)
//	return ItemDBToModel(ctx, q, tmp)
//}

//func InventoryDBToModel(ctx context.Context, q *sqlc.Queries, inv sqlc.Inventory) (inventory.Inventory, error) {
//	tmp := inventory.New()
//	tmp.ID = inv.ID
//
//	invalidTmp, err := q.GetItemFromName(ctx, "Invalid")
//	if err != nil {
//		return tmp, err
//	}
//	invalid, err := ItemDBToModel(ctx, q, invalidTmp)
//
//	var left item.Item
//	leftTmp, err := q.GetItem(ctx, inv.LeftHand)
//	if err != nil {
//		left = invalid
//	} else {
//		left, err = ItemDBToModel(ctx, q, leftTmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.LeftHand = left
//
//	var right item.Item
//	rightTmp, err := q.GetItem(ctx, inv.RightHand)
//	if err != nil {
//		right = invalid
//	} else {
//		right, err = ItemDBToModel(ctx, q, rightTmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.RightHand = right
//
//	var head item.Item
//	headTmp, err := q.GetItem(ctx, inv.Head)
//	if err != nil {
//		head = invalid
//	} else {
//		head, err = ItemDBToModel(ctx, q, headTmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Head = head
//
//	var torso item.Item
//	torsoTmp, err := q.GetItem(ctx, inv.Torso)
//	if err != nil {
//		torso = invalid
//	} else {
//		torso, err = ItemDBToModel(ctx, q, torsoTmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Torso = torso
//
//	var backpack0 item.Item
//	backpack0Tmp, err := q.GetItem(ctx, inv.Backpack0)
//	if err != nil {
//		backpack0 = invalid
//	} else {
//		backpack0, err = ItemDBToModel(ctx, q, backpack0Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Backpack0 = backpack0
//
//	var backpack1 item.Item
//	backpack1Tmp, err := q.GetItem(ctx, inv.Backpack1)
//	if err != nil {
//		backpack1 = invalid
//	} else {
//		backpack1, err = ItemDBToModel(ctx, q, backpack1Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Backpack1 = backpack1
//
//	var backpack2 item.Item
//	backpack2Tmp, err := q.GetItem(ctx, inv.Backpack2)
//	if err != nil {
//		backpack2 = invalid
//	} else {
//		backpack2, err = ItemDBToModel(ctx, q, backpack2Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Backpack2 = backpack2
//
//	var backpack3 item.Item
//	backpack3Tmp, err := q.GetItem(ctx, inv.Backpack3)
//	if err != nil {
//		backpack3 = invalid
//	} else {
//		backpack3, err = ItemDBToModel(ctx, q, backpack3Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Backpack3 = backpack3
//
//	var backpack4 item.Item
//	backpack4Tmp, err := q.GetItem(ctx, inv.Backpack4)
//	if err != nil {
//		backpack4 = invalid
//	} else {
//		backpack4, err = ItemDBToModel(ctx, q, backpack4Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Backpack4 = backpack4
//
//	var backpack5 item.Item
//	backpack5Tmp, err := q.GetItem(ctx, inv.Backpack5)
//	if err != nil {
//		backpack5 = invalid
//	} else {
//		backpack5, err = ItemDBToModel(ctx, q, backpack5Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Backpack5 = backpack5
//
//	var extra0 item.Item
//	extra0Tmp, err := q.GetItem(ctx, inv.ExtraSpace0)
//	if err != nil {
//		extra0 = invalid
//	} else {
//		extra0, err = ItemDBToModel(ctx, q, extra0Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.ExtraSpace0 = extra0
//
//	var extra1 item.Item
//	extra1Tmp, err := q.GetItem(ctx, inv.ExtraSpace1)
//	if err != nil {
//		extra1 = invalid
//	} else {
//		extra1, err = ItemDBToModel(ctx, q, extra1Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.ExtraSpace1 = extra1
//
//	var extra2 item.Item
//	extra2Tmp, err := q.GetItem(ctx, inv.ExtraSpace2)
//	if err != nil {
//		extra2 = invalid
//	} else {
//		extra2, err = ItemDBToModel(ctx, q, extra2Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.ExtraSpace2 = extra2
//
//	var extra3 item.Item
//	extra3Tmp, err := q.GetItem(ctx, inv.ExtraSpace3)
//	if err != nil {
//		extra3 = invalid
//	} else {
//		extra3, err = ItemDBToModel(ctx, q, extra3Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.ExtraSpace3 = extra3
//
//	var extra4 item.Item
//	extra4Tmp, err := q.GetItem(ctx, inv.ExtraSpace4)
//	if err != nil {
//		extra4 = invalid
//	} else {
//		extra4, err = ItemDBToModel(ctx, q, extra4Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.ExtraSpace4 = extra4
//
//	var extra5 item.Item
//	extra5Tmp, err := q.GetItem(ctx, inv.ExtraSpace5)
//	if err != nil {
//		extra5 = invalid
//	} else {
//		extra5, err = ItemDBToModel(ctx, q, extra5Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.ExtraSpace5 = extra5
//
//	var ground0 item.Item
//	ground0Tmp, err := q.GetItem(ctx, inv.Ground0)
//	if err != nil {
//		ground0 = invalid
//	} else {
//		ground0, err = ItemDBToModel(ctx, q, ground0Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Ground0 = ground0
//
//	var ground1 item.Item
//	ground1Tmp, err := q.GetItem(ctx, inv.Ground1)
//	if err != nil {
//		ground1 = invalid
//	} else {
//		ground1, err = ItemDBToModel(ctx, q, ground1Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Ground1 = ground1
//
//	var ground2 item.Item
//	ground2Tmp, err := q.GetItem(ctx, inv.Ground2)
//	if err != nil {
//		ground2 = invalid
//	} else {
//		ground2, err = ItemDBToModel(ctx, q, ground2Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Ground2 = ground2
//
//	var ground3 item.Item
//	ground3Tmp, err := q.GetItem(ctx, inv.Ground3)
//	if err != nil {
//		ground3 = invalid
//	} else {
//		ground3, err = ItemDBToModel(ctx, q, ground3Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Ground3 = ground3
//
//	var ground4 item.Item
//	ground4Tmp, err := q.GetItem(ctx, inv.Ground4)
//	if err != nil {
//		ground4 = invalid
//	} else {
//		ground4, err = ItemDBToModel(ctx, q, ground4Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Ground4 = ground4
//
//	var ground5 item.Item
//	ground5Tmp, err := q.GetItem(ctx, inv.Ground5)
//	if err != nil {
//		ground5 = invalid
//	} else {
//		ground5, err = ItemDBToModel(ctx, q, ground5Tmp)
//		if err != nil {
//			return tmp, err
//		}
//	}
//	tmp.Ground5 = ground5
//	tmp.SetCalculatedValues()
//
//	return tmp, nil
//}

func CharDBToModel(ctx context.Context, q *sqlc.Queries, c sqlc.Character) (character.Character, error) {
	traits := character.Traits{
		Physique: c.Physique,
		Skin: c.Skin,
		Hair: c.Hair,
		Face: c.Face,
		Speech: c.Speech,
		Clothing: c.Clothing,
		Virtue: c.Virtue,
		Vice: c.Vice,
		Reputation: c.Reputation,
		Misfortune: c.Misfortune,
	}

	status := character.Status{
		Hp: c.Hp,
		MaxHp: c.MaxHp,
		Str: c.Str,
		MaxStr: c.MaxStr,
		Dex: c.Dex,
		MaxDex: c.MaxDex,
		Will: c.Will,
		MaxWill: c.MaxWill,
	}

	tmp := character.Character{
		ID: c.ID,
		Description: c.Description,
		//Associations: assoc,
		//Identity: ident,
		Traits: traits,
		Status: status,
		Inventory: c.Inventory,
	}

	background, err := q.GetBackground(ctx, c.Background)
	if err != nil {
		log.Error("Failed to lookup Background for Character", "cid", c.ID, "name", c.Name, "surname", c.Surname)
		return tmp, err
	}

	ident := character.Identity{
		Gender: c.Gender,
		Name: c.Name,
		Surname: c.Surname,
		Age: c.Age,
		Portrait: c.Portrait,
		Background: background,
	}
	tmp.Identity = ident

	user, err := q.GetUser(ctx, c.User)
	if err != nil {
		log.Error("Failed to lookup User for Character", "cid", c.ID, "user", c.User)
		return tmp, err
	}

	party, err := q.GetParty(ctx, c.Party)
	if err != nil {
		log.Error("Failed to lookup Party for Character", "cid", c.ID, "party", c.Party)
		return tmp, err
	}

	village, err := q.GetVillage(ctx, c.Village)
	if err != nil {
		log.Error("Failed to lookup Village for Character", "cid", c.ID, "village", c.Village)
		return tmp, err
	}

	assoc := character.Associations{
		User: user,
		Party: party,
		Village: village,
	}
	tmp.Associations = assoc

	return tmp, nil
}

func CharViewRowToModel(ctx context.Context, q *sqlc.Queries, c sqlc.GetCharacterViewRow) character.Character {
	assoc := character.Associations{
		User: c.User,
		Party: c.Party,
		Village: c.Village,
	}

	ident := character.Identity{
		Gender: c.Gender,
		Name: c.Name,
		Surname: c.Surname,
		Age: c.Age,
		Portrait: c.Portrait,
		Background: c.Background,
	}

	traits := character.Traits{
		Physique: c.Physique,
		Skin: c.Skin,
		Hair: c.Hair,
		Face: c.Face,
		Speech: c.Speech,
		Clothing: c.Clothing,
		Virtue: c.Virtue,
		Vice: c.Vice,
		Reputation: c.Reputation,
		Misfortune: c.Misfortune,
	}

	status := character.Status{
		Hp: c.Hp,
		MaxHp: c.MaxHp,
		Str: c.Str,
		MaxStr: c.MaxStr,
		Dex: c.Dex,
		MaxDex: c.MaxDex,
		Will: c.Will,
		MaxWill: c.MaxWill,
	}

	tmp := character.Character{
		ID: c.ID,
		Description: c.Description,
		Associations: assoc,
		Identity: ident,
		Traits: traits,
		Status: status,
		Inventory: c.Inventory.ID,
	}

	return tmp
}

func CharViewFromNameRowToModel(ctx context.Context, q *sqlc.Queries, c sqlc.GetCharacterViewFromNameRow) character.Character {
	assoc := character.Associations{
		User: c.User,
		Party: c.Party,
		Village: c.Village,
	}

	ident := character.Identity{
		Gender: c.Gender,
		Name: c.Name,
		Surname: c.Surname,
		Age: c.Age,
		Portrait: c.Portrait,
		Background: c.Background,
	}

	traits := character.Traits{
		Physique: c.Physique,
		Skin: c.Skin,
		Hair: c.Hair,
		Face: c.Face,
		Speech: c.Speech,
		Clothing: c.Clothing,
		Virtue: c.Virtue,
		Vice: c.Vice,
		Reputation: c.Reputation,
		Misfortune: c.Misfortune,
	}

	status := character.Status{
		Hp: c.Hp,
		MaxHp: c.MaxHp,
		Str: c.Str,
		MaxStr: c.MaxStr,
		Dex: c.Dex,
		MaxDex: c.MaxDex,
		Will: c.Will,
		MaxWill: c.MaxWill,
	}

	tmp := character.Character{
		ID: c.ID,
		Description: c.Description,
		Associations: assoc,
		Identity: ident,
		Traits: traits,
		Status: status,
		Inventory: c.Inventory.ID,
	}

	return tmp
}
