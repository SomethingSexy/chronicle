package game

type GameCharacterRequest struct {
	ID      string `jsonapi:"primary,game-characters"`
	GameId  string `jsonapi:"attr,gameId"`
	WorldId string `jsonapi:"attr,worldId"`
	Name    string `jsonapi:"attr,name"`
	Type    string `jsonapi:"attr,type"`
}
