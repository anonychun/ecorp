package auth

import (
	"context"

	"github.com/anonychun/benih/internal/consts"
	"github.com/anonychun/benih/internal/current"
	"github.com/anonychun/benih/internal/entity"
)

func (u *Usecase) SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error) {
	validationErr := u.validator.Struct(&req)
	if validationErr.IsFail() {
		return nil, validationErr
	}

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
	err := u.repository.AdminSession.DeleteByToken(ctx, req.Token)
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
	res.Admin.Name = admin.Name

	return res, nil
}
