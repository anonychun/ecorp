package admin

import "github.com/anonychun/ecorp/internal/entity"

type AdminDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewAdminDto(admin *entity.Admin) *AdminDto {
	return &AdminDto{
		Id:   admin.Id.String(),
		Name: admin.Name,
	}
}

type FindByIdRequest struct {
	Id string `param:"id"`
}

type CreateRequest struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type UpdateRequest struct {
	Id           string `param:"id"`
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

type DeleteRequest struct {
	Id string `param:"id"`
}
