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
	Name         string `json:"name" validate:"required" field:"name" label:"Name"`
	EmailAddress string `json:"emailAddress" validate:"required|email" field:"emailAddress" label:"Email address"`
	Password     string `json:"password" validate:"required|minLen:8" field:"password" label:"Password"`
}

type UpdateRequest struct {
	Id           string `param:"id"`
	Name         string `json:"name" validate:"required" field:"name" label:"Name"`
	EmailAddress string `json:"emailAddress" validate:"required|email" field:"emailAddress" label:"Email address"`
}
