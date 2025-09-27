package admin

import (
	"context"

	"github.com/anonychun/ecorp/internal/consts"
	"github.com/anonychun/ecorp/internal/entity"
	"github.com/samber/lo"
)

func (u *Usecase) FindAll(ctx context.Context) ([]*AdminBlueprint, error) {
	admins, err := u.repository.Admin.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := lo.Map(admins, func(admin *entity.Admin, _ int) *AdminBlueprint {
		return NewAdminBlueprint(admin)
	})

	return res, nil
}

func (u *Usecase) FindById(ctx context.Context, req FindByIdRequest) (*AdminBlueprint, error) {
	admin, err := u.repository.Admin.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrAdminNotFound
	} else if err != nil {
		return nil, err
	}

	return NewAdminBlueprint(admin), nil
}

func (u *Usecase) Create(ctx context.Context, req CreateRequest) (*AdminBlueprint, error) {
	admin := &entity.Admin{
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
	}

	err := admin.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = u.repository.Admin.Create(ctx, admin)
	if err != nil {
		return nil, err
	}

	return NewAdminBlueprint(admin), nil
}

func (u *Usecase) Update(ctx context.Context, req UpdateRequest) (*AdminBlueprint, error) {
	isEmailAddressExists, err := u.repository.Admin.ExistsByEmailAddressAndNotId(ctx, req.EmailAddress, req.Id)
	if err != nil {
		return nil, err
	}

	if isEmailAddressExists {
		return nil, consts.ErrEmailAddressAlreadyRegistered
	}

	admin, err := u.repository.Admin.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrAdminNotFound
	} else if err != nil {
		return nil, err
	}

	admin.Name = req.Name
	admin.EmailAddress = req.EmailAddress

	err = u.repository.Admin.Update(ctx, admin)
	if err != nil {
		return nil, err
	}

	return NewAdminBlueprint(admin), nil
}
