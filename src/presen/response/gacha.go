package response

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

type GachaResult struct {
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
	IconURL     string `json:"icon_url"`
	Rarity      string `json:"rarity"`
}
