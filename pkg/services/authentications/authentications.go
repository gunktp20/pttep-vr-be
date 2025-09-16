package authentications

import (
	"context"
	"errors"
	"fmt"
	"pttep-vr-api/pkg/constant"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/jwt"
)

func (s *Service) GetTypes(ctx context.Context) ([]models.LoginType, error) {
	return s.repository.FindLoginTypes(ctx)
}

func (o *Service) Login(ctx context.Context, userLogin models.UserLogin) (models.User, string, error) {

	//? verify email
	var loginType models.LoginType
	var err error
	loginType.SetID(userLogin.LoginTypeID)
	loginType, err = o.repository.FindOneLoginTypes(ctx, loginType)
	if err != nil {
		return models.User{}, "", err
	}

	switch constant.TYPELoginType(loginType.Key) {
	//case constant.LoginType.Guest():
	//register new user for guest
	//case constant.LoginType.Telephone():
	//check duplicate telephone number, next to verify telephone
	case constant.LoginType.Email(), constant.LoginType.Contractor():
		// check duplicate email, next to verify email
		userLogin, err = o.repository.FindOneUserLoginsByUsernameAndPassword(ctx, userLogin)
		if err != nil {
			return models.User{}, "", err
		}
		if fmt.Sprintf("%v", userLogin.ID) == "0" {
			return models.User{}, "", errors.New("invalid username or password.")
		}
	default:
		return models.User{}, "", errors.New("invalid login type")
	}

	//find user information
	user, err := o.repository.FindOneUsers(userLogin.UserID)
	if err != nil {
		return models.User{}, "", err
	}
	//create token
	jwtObj := jwt.JWTAuthService()
	token := jwtObj.GenerateToken(user.Email, true)
	return user, token, nil
}
