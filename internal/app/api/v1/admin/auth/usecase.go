package auth

import (
	"context"

	"github.com/anonychun/ecorp/internal/consts"
	"github.com/anonychun/ecorp/internal/current"
	"github.com/anonychun/ecorp/internal/entity"
)

func (u *Usecase) SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error) {
	admin, err := u.repository.Admin.FindByEmailAddress(ctx, req.EmailAddress)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	err = admin.ComparePassword(req.Password)
	if err != nil {
		return nil, consts.ErrInvalidCredentials
	}

	adminSession := &entity.AdminSession{
		AdminId:   admin.Id,
		IpAddress: req.IpAddress,
		UserAgent: req.UserAgent,
	}
	adminSession.GenerateToken()

	err = u.repository.AdminSession.Create(ctx, adminSession)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{Token: adminSession.Token}, nil
}

func (u *Usecase) SignOut(ctx context.Context, req SignOutRequest) error {
	adminSession, err := u.repository.AdminSession.FindByToken(ctx, req.Token)
	if err == consts.ErrRecordNotFound {
		return consts.ErrUnauthorized
	} else if err != nil {
		return err
	}

	err = u.repository.AdminSession.DeleteById(ctx, adminSession.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) Me(ctx context.Context) (*MeResponse, error) {
	admin := current.Admin(ctx)
	if admin == nil {
		return nil, consts.ErrUnauthorized
	}

	res := &MeResponse{}
	res.Admin.Id = admin.Id.String()
	res.Admin.EmailAddress = admin.EmailAddress

	return res, nil
}
