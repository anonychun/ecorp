package admin

import "github.com/anonychun/ecorp/internal/entity"

type AdminBlueprint struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewAdminBlueprint(admin *entity.Admin) *AdminBlueprint {
	return &AdminBlueprint{
		Id:   admin.Id.String(),
		Name: admin.Name,
	}
}

type FindByIdRequest struct {
	Id string `param:"id"`
}

type CreateRequest struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type UpdateRequest struct {
	Id           string `param:"id"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
}
