package auth

import (
	"context"

	"github.com/anonychun/ecorp/internal/consts"
	"github.com/anonychun/ecorp/internal/current"
	"github.com/anonychun/ecorp/internal/entity"
	"github.com/anonychun/ecorp/internal/repository"
)

func (u *Usecase) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	validationErr := u.validator.Struct(&req)
	isEmailAddressExists, err := u.repository.User.ExistsByEmailAddress(ctx, req.EmailAddress)
	if err != nil {
		return nil, err
	}

	if isEmailAddressExists {
		validationErr.AddError("emailAddress", consts.ErrEmailAddressAlreadyRegistered)
	}

	if validationErr.IsFail() {
		return nil, validationErr
	}

	user := &entity.User{
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
	}

	err = user.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	res := &SignUpResponse{}
	repository.Transaction(ctx, func(ctx context.Context) error {
		err = u.repository.User.Create(ctx, user)
		if err != nil {
			return err
		}

		userSession := &entity.UserSession{
			UserId:    user.Id,
			IpAddress: req.IpAddress,
			UserAgent: req.UserAgent,
		}
		userSession.GenerateToken()

		err = u.repository.UserSession.Create(ctx, userSession)
		if err != nil {
			return err
		}

		res.Token = userSession.Token
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Usecase) SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error) {
	validationErr := u.validator.Struct(&req)
	if validationErr.IsFail() {
		return nil, validationErr
	}

	user, err := u.repository.User.FindByEmailAddress(ctx, req.EmailAddress)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	err = user.ComparePassword(req.Password)
	if err != nil {
		return nil, consts.ErrInvalidCredentials
	}

	userSession := &entity.UserSession{
		UserId:    user.Id,
		IpAddress: req.IpAddress,
		UserAgent: req.UserAgent,
	}
	userSession.GenerateToken()

	err = u.repository.UserSession.Create(ctx, userSession)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{Token: userSession.Token}, nil
}

func (u *Usecase) SignOut(ctx context.Context, req SignOutRequest) error {
	err := u.repository.UserSession.DeleteByToken(ctx, req.Token)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) Me(ctx context.Context) (*MeResponse, error) {
	user := current.User(ctx)
	if user == nil {
		return nil, consts.ErrUnauthorized
	}

	res := &MeResponse{}
	res.User.Id = user.Id.String()
	res.User.Name = user.Name
	res.User.EmailAddress = user.EmailAddress

	return res, nil
}
