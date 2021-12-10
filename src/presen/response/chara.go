package response

type CharacterListResponse struct {
	Characters []UserCharacter
}

type UserCharacter struct {
	UserCharacterID string
	CharacterID     string
	Name            string
}
