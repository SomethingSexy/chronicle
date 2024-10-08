package domain

func NewCharacterType(t string) CharacterType {
	switch t {
	case "NPC":
		return NPC
	case "PC":
		return PC
	}

	return NPC
}

type CharacterType int

const (
	NPC = iota
	PC
)

func (t CharacterType) String() string {
	return [...]string{"NPC", "PC"}[t-1]
}

// Represents a character that has been added to a world
type WorldCharacter struct {
	Type CharacterType
}
