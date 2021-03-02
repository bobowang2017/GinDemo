package dto

type UserDto struct {
	PageNum  int
	PageSize int
	Username string `binding:"required"` // string默认长度为255, 使用这种tag重设。
	Password string `binding:"required,min=6,max=12"`
	Name     string
	Sex      *int `binding:"required"`
	Birthday string
}
