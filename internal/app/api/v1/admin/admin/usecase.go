package admin

import (
	"context"

	"github.com/anonychun/ecorp/internal/consts"
	"github.com/anonychun/ecorp/internal/entity"
	"github.com/samber/lo"
)

func (u *Usecase) FindAll(ctx context.Context) ([]*AdminDto, error) {
	admins, err := u.repository.Admin.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := lo.Map(admins, func(admin *entity.Admin, _ int) *AdminDto {
		return NewAdminDto(admin)
	})

	return res, nil
}

func (u *Usecase) FindById(ctx context.Context, req FindByIdRequest) (*AdminDto, error) {
	admin, err := u.repository.Admin.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrAdminNotFound
	} else if err != nil {
		return nil, err
	}

	return NewAdminDto(admin), nil
}

func (u *Usecase) Create(ctx context.Context, req CreateRequest) (*AdminDto, error) {
	admin := &entity.Admin{
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
	}

	err := admin.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if err := u.repository.Admin.Create(ctx, admin); err != nil {
		return nil, err
	}

	return NewAdminDto(admin), nil
}

func (u *Usecase) Update(ctx context.Context, req UpdateRequest) (*AdminDto, error) {
	admin, err := u.repository.Admin.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrAdminNotFound
	} else if err != nil {
		return nil, err
	}

	admin.Name = req.Name
	admin.EmailAddress = req.EmailAddress

	if err := u.repository.Admin.Update(ctx, admin); err != nil {
		return nil, err
	}

	return NewAdminDto(admin), nil
}

func (u *Usecase) Delete(ctx context.Context, req DeleteRequest) error {
	exists, err := u.repository.Admin.ExistsById(ctx, req.Id)
	if err != nil {
		return err
	}

	if !exists {
		return consts.ErrAdminNotFound
	}

	return u.repository.Transaction(ctx, func(ctx context.Context) error {
		err := u.repository.AdminSession.DeleteAllByAdminId(ctx, req.Id)
		if err != nil {
			return err
		}

		return u.repository.Admin.DeleteById(ctx, req.Id)
	})
}
