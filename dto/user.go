package dto

import "GinDemo/models"

type UserDto struct {
	models.User
	PageNum  int
	PageSize int
}
