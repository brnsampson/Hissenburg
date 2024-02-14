package background

//go:generate stringer -type=Background
type Background int
const (
	Alchemist Background = iota
	Blacksmith
	Butcher
	Burglar
	Carpenter
	Cleric
	Gambler
	Gravedigger
	Herbalist
	Hunter
	Magician
	Mercenary
	Merchant
	Miner
	Outlaw
	Performer
	Pickpocket
	Smuggler
	Servant
	Ranger
)

func (n Background) Count() int {
	return 20
}
