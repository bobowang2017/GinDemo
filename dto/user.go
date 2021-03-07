package dto

type UserDto struct {
	PageNum  int
	PageSize int
	Username string `binding:"required"`
	Password string `binding:"required,min=6,max=12"`
	Name     string
	Sex      *int `binding:"required"`
	Birthday string
}

type UserAddrDto struct {
	UserId    int
	Consignee string `binding:"required"`
	Province  int    `binding:"required,gte=1,lte=40"`
	City      int    `binding:"required,gte=1,lte=40"`
	District  int    `binding:"required,gte=1,lte=40"`
	Street    string
	Zipcode   string
	Mobile    string `binding:"required,min=11,max=11"`
	Default   bool
}
