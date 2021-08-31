package tools

import (
	"lawSite/models"
)

func CreateToken(token string) *models.Token {
	return &models.Token{Data: &models.TokenData{Token: token}}
}
