package response

import (
	"Template-Go/models"
)

type RegisterResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}
