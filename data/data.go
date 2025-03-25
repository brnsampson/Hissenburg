package data

import (
	"github.com/brnsampson/Hissenburg/models/character"
	"context"
	"database/sql"
	"github.com/charmbracelet/log"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
)

type MetaRepo interface {
	GetUser(ctx context.Context, id int64) (sqlc.User, error)
	GetUserFromName(ctx context.Context, name string) (sqlc.User, error)
	CreateUser(ctx context.Context, name string) (sqlc.User, error)
	ListParties(ctx context.Context) ([]sqlc.Party, error)
	GetParty(ctx context.Context, id int64) (sqlc.Party, error)
	GetPartyFromName(ctx context.Context, name string) (sqlc.Party, error)
	CreateParty(ctx context.Context, party sqlc.Party) (sqlc.Party, error)
	ListVillages(ctx context.Context) ([]sqlc.Village, error)
	GetVillage(ctx context.Context, id int64) (sqlc.Village, error)
	GetVillageFromName(ctx context.Context, name string) (sqlc.Village, error)
	CreateVillage(ctx context.Context, name string) (sqlc.Village, error)
}

type CharRepo interface {
	MetaRepo
	TraitRepo
	IdentityRepo
	InventoryRepo
	ListCharacters(ctx context.Context) ([]sqlc.Character, error)
	ListCharacterViews(ctx context.Context) ([]character.Character, error)
	GetCharacter(ctx context.Context, id int64) (sqlc.Character, error)
	GetCharacterView(ctx context.Context, id int64) (character.Character, error)
	GetCharacterFromName(ctx context.Context, name, surname string) (sqlc.Character, error)
	GetCharacterViewFromName(ctx context.Context, name, surname string) (character.Character, error)
	CreateCharacter(ctx context.Context, c character.Character) (character.Character, error)
	DeleteCharacter(ctx context.Context, id int64) error
	GetAssociations(ctx context.Context, id int64) (character.Associations, error)
	GetAssociationsFromName(ctx context.Context, name, surname string) (character.Associations, error)
	ListAssociations(ctx context.Context) ([]character.Associations, error)
	UpdateAssociations(ctx context.Context, assoc character.Associations) error
	GetStatus(ctx context.Context, id int64) (character.Status, error)
	GetStatusFromName(ctx context.Context, name, surname string) (character.Status, error)
	UpdateStatus(ctx context.Context, status character.Status) error
	UpdateMaxStatus(ctx context.Context, status character.Status) error
	GetTraits(ctx context.Context, id int64) (character.Traits, error)
	GetTraitsFromName(ctx context.Context, name, surname string) (character.Traits, error)
	UpdateTraits(ctx context.Context, traits character.Traits) error
	GetIdentity(ctx context.Context, id int64) (character.Identity, error)
	GetIdentityFromName(ctx context.Context, name, surname string) (character.Identity, error)
	UpdateIdentity(ctx context.Context, ident character.Identity) error
}

type IdentityRepo interface {
	GetRandomGender(ctx context.Context) (sqlc.Gender, error)
	ListGenders(ctx context.Context) ([]sqlc.Gender, error)
	GetBackgroundFromTitle(ctx context.Context, title string) (sqlc.Background, error)
	GetRandomBackground(ctx context.Context) (sqlc.Background, error)
	ListBackgrounds(ctx context.Context) ([]sqlc.Background, error)
	GetRandomName(ctx context.Context) (sqlc.Name, error)
	ListNames(ctx context.Context) ([]sqlc.Name, error)
	GetRandomMasculineName(ctx context.Context) (sqlc.Name, error)
	ListMasculineNames(ctx context.Context) ([]string, error)
	GetRandomFeminineName(ctx context.Context) (sqlc.Name, error)
	ListFeminineNames(ctx context.Context) ([]string, error)
	GetRandomSurname(ctx context.Context) (sqlc.Surname, error)
	ListSurnames(ctx context.Context) ([]sqlc.Surname, error)
}

type TraitRepo interface {
	GetRandomPhysique(ctx context.Context) (sqlc.Physique, error)
	ListPhysiques(ctx context.Context) ([]sqlc.Physique, error)
	GetRandomSkin(ctx context.Context) (sqlc.Skin, error)
	ListSkin(ctx context.Context) ([]sqlc.Skin, error)
	GetRandomHair(ctx context.Context) (sqlc.Hair, error)
	ListHair(ctx context.Context) ([]sqlc.Hair, error)
	GetRandomFace(ctx context.Context) (sqlc.Face, error)
	ListFaces(ctx context.Context) ([]sqlc.Face, error)
	GetRandomSpeech(ctx context.Context) (sqlc.Speech, error)
	ListSpeech(ctx context.Context) ([]sqlc.Speech, error)
	GetRandomClothing(ctx context.Context) (sqlc.Clothing, error)
	ListClothing(ctx context.Context) ([]sqlc.Clothing, error)
	GetRandomVirtue(ctx context.Context) (sqlc.Virtue, error)
	ListVirtues(ctx context.Context) ([]sqlc.Virtue, error)
	GetRandomVice(ctx context.Context) (sqlc.Vice, error)
	ListVices(ctx context.Context) ([]sqlc.Vice, error)
	GetRandomReputation(ctx context.Context) (sqlc.Reputation, error)
	ListReputations(ctx context.Context) ([]sqlc.Reputation, error)
	GetRandomMisfortune(ctx context.Context) (sqlc.Misfortune, error)
	ListMisfortunes(ctx context.Context) ([]sqlc.Misfortune, error)
}

type Repository struct {
	*sqlc.Queries
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repository {
	queries := sqlc.New(db)
	return &Repository{ queries, db }
}

//func (r *Repository) GetItemKind(ctx context.Context, id int64) (sqlc.ItemKind, error) {
//	return r.queries.GetItemKind(ctx, id)
//}
//
//func (r *Repository) GetItemKindFromString(ctx context.Context, kind string) (sqlc.ItemKind, error) {
//	return r.queries.GetItemKindFromString(ctx, kind)
//}
//
//func (r *Repository) GetItemSlot(ctx context.Context, id int64) (sqlc.ItemSlot, error) {
//	return r.queries.GetItemSlot(ctx, id)
//}
//
//func (r *Repository) GetItemSlotFromString(ctx context.Context, slot string) (sqlc.ItemSlot, error) {
//	return r.queries.GetItemSlotFromString(ctx, slot)
//}
func (r *Repository) CreateParty(ctx context.Context, party sqlc.Party) (sqlc.Party, error){
	params := sqlc.CreatePartyParams{Name: party.Name, Description: party.Description}

	return r.Queries.CreateParty(ctx, params)
}

func (r *Repository) GetAssociations(ctx context.Context, id int64) (character.Associations, error){
	assoc, err := r.Queries.GetAssociations(ctx, id)
	if err != nil {
		return character.Associations{}, err
	}

	tmp := character.Associations{
		CharacterID: assoc.ID,
		User: assoc.User,
		Party: assoc.Party,
		Village: assoc.Village,
	}

	return tmp, nil
}

func (r *Repository) GetAssociationsFromName(ctx context.Context, name, surname string) (character.Associations, error){
	params := sqlc.GetAssociationsFromNameParams{Name: name, Surname: surname}
	assoc, err := r.Queries.GetAssociationsFromName(ctx, params)
	if err != nil {
		return character.Associations{}, err
	}

	tmp := character.Associations{
		CharacterID: assoc.ID,
		User: assoc.User,
		Party: assoc.Party,
		Village: assoc.Village,
	}

	return tmp, nil
}

func (r *Repository) ListAssociations(ctx context.Context) ([]character.Associations, error) {
	associations := make([]character.Associations, 0)
	rows, err := r.Queries.ListAssociations(ctx)
	if err != nil {
		return associations, err
	}

	for _, row := range rows {
		tmp := character.Associations{
			CharacterID: row.ID,
			User: row.User,
			Party: row.Party,
			Village: row.Village,
		}
		associations = append(associations, tmp)
	}

	return associations, nil
}

func (r *Repository) UpdateAssociations(ctx context.Context, assoc character.Associations) error{
	params := sqlc.UpdateAssociationsParams{
		User: assoc.User.ID,
		Party: assoc.Party.ID,
		Village: assoc.Village.ID,
		ID: assoc.CharacterID,
	}

	return r.Queries.UpdateAssociations(ctx, params)
}

func (r *Repository) GetIdentity(ctx context.Context, id int64) (character.Identity, error) {
	i, err := r.Queries.GetIdentity(ctx, id)
	if err != nil {
		return character.Identity{}, err
	}

	tmp := character.Identity{
		CharacterID: i.ID,
		Gender: i.Gender,
		Name: i.Name,
		Surname: i.Surname,
		Age: i.Age,
		Portrait: i.Portrait,
		Background: i.Background,
	}

	return tmp, nil
}

func (r *Repository) GetIdentityFromName(ctx context.Context, name, surname string) (character.Identity, error) {
	params := sqlc.GetIdentityFromNameParams{Name: name, Surname: surname}
	i, err := r.Queries.GetIdentityFromName(ctx, params)
	if err != nil {
		return character.Identity{}, err
	}

	tmp := character.Identity{
		CharacterID: i.ID,
		Gender: i.Gender,
		Name: i.Name,
		Surname: i.Surname,
		Age: i.Age,
		Portrait: i.Portrait,
		Background: i.Background,
	}

	return tmp, nil
}

func (r *Repository) UpdateIdentity(ctx context.Context, i character.Identity) error {
	params := sqlc.UpdateIdentityParams {
		ID: i.CharacterID,
		Name: i.Name,
		Surname: i.Surname,
		Age: i.Age,
		Portrait: i.Portrait,
		Gender: i.Gender,
		Background: i.Background.ID,
	}

	return r.Queries.UpdateIdentity(ctx, params)
}

func (r *Repository) GetTraits(ctx context.Context, id int64) (character.Traits, error) {
	t, err := r.Queries.GetTraits(ctx, id)
	if err != nil {
		return character.Traits{}, err
	}

	traits := character.Traits {
		CharacterID: t.ID,
		Physique: t.Physique,
		Skin: t.Skin,
		Hair: t.Hair,
		Face: t.Face,
		Speech: t.Speech,
		Clothing: t.Clothing,
		Virtue: t.Virtue,
		Vice: t.Vice,
		Reputation: t.Reputation,
		Misfortune: t.Misfortune,
	}

	return traits, nil
}

func (r *Repository) GetTraitsFromName(ctx context.Context, name, surname string) (character.Traits, error) {
	params := sqlc.GetTraitsFromNameParams{Name: name, Surname: surname}
	t, err := r.Queries.GetTraitsFromName(ctx, params)
	if err != nil {
		return character.Traits{}, err
	}

	traits := character.Traits {
		CharacterID: t.ID,
		Physique: t.Physique,
		Skin: t.Skin,
		Hair: t.Hair,
		Face: t.Face,
		Speech: t.Speech,
		Clothing: t.Clothing,
		Virtue: t.Virtue,
		Vice: t.Vice,
		Reputation: t.Reputation,
		Misfortune: t.Misfortune,
	}

	return traits, nil
}

func (r *Repository) UpdateTraits(ctx context.Context, traits character.Traits) error {
	params := sqlc.UpdateTraitsParams {
		ID: traits.CharacterID,
		Physique: traits.Physique,
		Skin: traits.Skin,
		Hair: traits.Hair,
		Face: traits.Face,
		Speech: traits.Speech,
		Clothing: traits.Clothing,
		Virtue: traits.Virtue,
		Vice: traits.Vice,
		Reputation: traits.Reputation,
		Misfortune: traits.Misfortune,
	}

	return r.Queries.UpdateTraits(ctx, params)
}

func (r *Repository) GetStatus(ctx context.Context, id int64) (character.Status, error) {
	s, err := r.Queries.GetStatus(ctx, id)
	if err != nil {
		return character.Status{}, err
	}

	status := character.Status {
		CharacterID: s.ID,
		Hp: s.Hp,
		MaxHp: s.MaxHp,
		Str: s.Str,
		MaxStr: s.MaxStr,
		Dex: s.Dex,
		MaxDex: s.MaxDex,
		Will: s.Will,
		MaxWill: s.MaxWill,
	}

	return status, nil
}

func (r *Repository) GetStatusFromName(ctx context.Context, name, surname string) (character.Status, error) {
	params := sqlc.GetStatusFromNameParams{Name: name, Surname: surname}
	s, err := r.Queries.GetStatusFromName(ctx, params)
	if err != nil {
		return character.Status{}, err
	}

	status := character.Status {
		CharacterID: s.ID,
		Hp: s.Hp,
		MaxHp: s.MaxHp,
		Str: s.Str,
		MaxStr: s.MaxStr,
		Dex: s.Dex,
		MaxDex: s.MaxDex,
		Will: s.Will,
		MaxWill: s.MaxWill,
	}

	return status, nil
}

func (r *Repository) UpdateStatus(ctx context.Context, status character.Status) error {
	params := sqlc.UpdateStatusParams {
		ID: status.CharacterID,
		Hp: status.Hp,
		Str: status.Str,
		Dex: status.Dex,
		Will: status.Will,
	}
	return r.Queries.UpdateStatus(ctx, params)
}

func (r *Repository) UpdateMaxStatus(ctx context.Context, status character.Status) error {
	params := sqlc.UpdateMaxStatusParams {
		ID: status.CharacterID,
		MaxHp: status.MaxHp,
		MaxStr: status.MaxStr,
		MaxDex: status.MaxDex,
		MaxWill: status.MaxWill,
	}
	return r.Queries.UpdateMaxStatus(ctx, params)
}

// Handled automatically by embedded r.Queries
//func (r *Repository) DeleteTraits(ctx context.Context, id int64) error {
//}

func (r *Repository) ListCharacterViews(ctx context.Context) ([]character.Character, error) {
	chars := make([]character.Character, 0)

	charList, err := r.Queries.ListCharacters(ctx)
	if err != nil {
		return chars, err
	}

	for _, c := range charList {
		tmp, err := CharDBToModel(ctx, r.Queries, c)
		if err != nil {
			return chars, err
		}
		chars = append(chars, tmp)
	}

	return chars, nil
}

//func (r *Repository) GetCharacter(ctx context.Context, id int64) (character.Character, error) {
//	char, err := r.Queries.GetCharacter(ctx, id)
//	if err != nil {
//		return character.Character{}, err
//	}
//
//	return CharDBToModel(ctx, r.Queries, char)
//}

//automatic
//func (r *Repository) GetCharacter(ctx context.Context, id int64) (sqlc.Character, error)

func (r *Repository) GetCharacterView(ctx context.Context, id int64) (character.Character, error) {
	char, err := r.Queries.GetCharacterView(ctx, id)
	if err != nil {
		return character.Character{}, err
	}

	return CharViewRowToModel(ctx, r.Queries, char), nil
}

func (r *Repository) GetCharacterFromName(ctx context.Context, name, surname string) (sqlc.Character, error) {
	params := sqlc.GetCharacterFromNameParams{ Name: name, Surname: surname }
	return r.Queries.GetCharacterFromName(ctx, params)
}

func (r *Repository) GetCharacterViewFromName(ctx context.Context, name, surname string) (character.Character, error) {
	params := sqlc.GetCharacterViewFromNameParams{ Name: name, Surname: surname }
	char, err := r.Queries.GetCharacterViewFromName(ctx, params)
	if err != nil {
		return character.Character{}, err
	}

	return CharViewFromNameRowToModel(ctx, r.Queries, char), nil
}

func (r *Repository) CreateCharacter(ctx context.Context, c character.Character) (character.Character, error) {
	user, err := r.GetUser(ctx, c.Associations.User.ID)
	if err != nil {
		log.Error("Failed to lookup user when creating new character", "user", c.Associations.User.Name, "uid", c.Associations.User.ID, "error", err)
		return character.Character{}, err
	}

	party, err := r.GetParty(ctx, c.Associations.Party.ID)
	if err != nil {
		log.Error("Failed to lookup party when creating new character", "party", c.Associations.Party.Name, "party_id", c.Associations.Party.ID, "error", err)
		return character.Character{}, err
	}

	village, err := r.GetVillage(ctx, c.Associations.Village.ID)
	if err != nil {
		log.Error("Failed to lookup village when creating new character", "village", c.Associations.Village.Name, "uid", c.Associations.Village.ID, "error", err)
		return character.Character{}, err
	}

	inv, err := r.CreateInventory(ctx)
	if err != nil {
		log.Error("Failed to create new inventory for character", "inventory", inv, "character", c, "error", err)
		return character.Character{}, err
	}

	charParams := sqlc.CreateCharacterParams {
		Description: c.Description,
		User: user.ID,
		Party: party.ID,
		Village: village.ID,
		Inventory: inv,
		Gender: c.Identity.Gender,
		Name: c.Identity.Name,
		Surname: c.Identity.Surname,
		Age: c.Identity.Age,
		Portrait: c.Identity.Portrait,
		Background: c.Identity.Background.ID,
		Physique: c.Traits.Physique,
		Skin: c.Traits.Skin,
		Hair: c.Traits.Hair,
		Face: c.Traits.Face,
		Speech: c.Traits.Speech,
		Clothing: c.Traits.Clothing,
		Virtue: c.Traits.Virtue,
		Vice: c.Traits.Vice,
		Reputation: c.Traits.Reputation,
		Misfortune: c.Traits.Misfortune,
		Hp: c.Status.Hp,
		MaxHp: c.Status.MaxHp,
		Str: c.Status.Str,
		MaxStr: c.Status.MaxStr,
		Dex: c.Status.Dex,
		MaxDex: c.Status.MaxDex,
		Will: c.Status.Will,
		MaxWill: c.Status.MaxWill,
	}
	newChar, err := r.Queries.CreateCharacter(ctx, charParams)
	if err != nil {
		log.Error("Failed to create new character", "character", charParams, "error", err)
		return character.Character{}, err
	}

	log.Debug("Created Character", "character", newChar)
	return CharDBToModel(ctx, r.Queries, newChar)
}
