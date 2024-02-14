package status

//go:generate stringer -type=StatusType
type StatusType int

const (
	HP StatusType = iota
	MaxHP
	Str
	MaxStr
	Dex
	MaxDex
	Will
	MaxWill
)

type Status struct {
	HP uint8
	MaxHP uint8
	Str uint8
	MaxStr uint8
	Dex uint8
	MaxDex uint8
	Will uint8
	MaxWill uint8
}

func New() Status {
	return Status{HP: 0, MaxHP: 0, Str: 0, MaxStr: 0, Dex: 0, MaxDex: 0, Will: 0, MaxWill: 0}
}
