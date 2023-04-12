package client

type CharacterClient interface {
	GetAllCharacters() ([]*CharacterClientResponse, error)
}

type CharacterClientResponse struct {
	Name        string
	Description string
	ClientID    string
	Copyright   string
	ImageURL    string
}

func NewCharacterClientResponse(name, description, clientID, copyright, imageURL string) *CharacterClientResponse {
	return &CharacterClientResponse{
		Name:        name,
		Description: description,
		ClientID:    clientID,
		Copyright:   copyright,
		ImageURL:    imageURL,
	}
}
