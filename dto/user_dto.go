package dto

type UserDto struct {
	Nik        string `json:"nik" form:"nik" validate:"required"`
	Name       string `json:"name" form:"name" validate:"required"`
	Email      string `json:"email" form:"email" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
	IsActive   int    `json:"is_active" form:"is_active" validate:"required"`
	TypeUser   int    `json:"type_user" form:"type_user"`
	UserRoleID string `json:"user_role_id" form:"user_role_id" validate:"required"`
}

type UserUpdateDto struct {
	Name     string `json:"name" form:"name" query:"name" validate:"required"`
	Email    string `json:"email" form:"email" query:"email" validate:"required"`
	Password string `json:"password" form:"password" query:"password"`
	Image    string `json:"image" form:"image" query:"image"`
}
