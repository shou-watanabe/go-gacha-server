package response

type CharacterListResponse struct {
	Characters []UserCharacter `json:"user_character"`
}

type UserCharacter struct {
	UserCharacterID string `json:"user_character_id"`
	CharacterID     string `json:"character_id"`
	Name            string `json:"name"`
	IconURL         string `json:"icon_url"`
}
