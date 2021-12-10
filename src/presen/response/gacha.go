package response

type GachaDrawResponse struct {
	Results []GachaResult
}

type GachaResult struct {
	CharacterID string
	Name        string
}
