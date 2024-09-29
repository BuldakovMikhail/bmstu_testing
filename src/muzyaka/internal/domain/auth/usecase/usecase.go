package usecase

import (
	"github.com/pkg/errors"
	"src/internal/domain/user/repository"
	"src/internal/lib/jwt"
	"src/internal/lib/validation"
	"src/internal/models"
)

type AuthUseCase interface {
	SignUp(user *models.User) (*models.AuthToken, error)
	SignIn(email string, password string) (*models.AuthToken, error)
	BasicAuthorization(token *models.AuthToken) (uint64, string, error)
}

const (
	UserRole = "user"
)

type usecase struct {
	userRep       repository.UserRepository
	tokenProvider jwt.TokenProvider
	encryptor     Encryptor
}

func NewAuthUseCase(tokenProvider jwt.TokenProvider,
	userRep repository.UserRepository,
	enc Encryptor) AuthUseCase {
	return &usecase{
		tokenProvider: tokenProvider,
		userRep:       userRep,
		encryptor:     enc,
	}
}

func (u *usecase) SignUp(user *models.User) (*models.AuthToken, error) {
	if user.Password == "" {
		return nil, models.ErrInvalidPassword
	}

	if !validation.ValidateWithoutSpace(user.Password) {
		return nil, models.ErrInvalidPassword
	}

	if !validation.ValidateWithoutSpace(user.Email) {
		return nil, models.ErrInvalidLogin
	}

	encPassword, err := u.encryptor.EncodePassword([]byte(user.Password))
	if err != nil {
		return nil, errors.Wrap(err, "auth.usecase.SignUp encode error")
	}

	temp := user
	temp.Password = string(encPassword)
	id, err := u.userRep.AddUser(temp)
	temp.Id = id

	if err != nil {
		return nil, errors.Wrap(err, "auth.usecase.SignUp AddUser error")
	}

	jwtToken, err := u.tokenProvider.GenerateToken(temp)

	if err != nil {
		return nil, errors.Wrap(err, "auth.usecase.SignUp token generation error")
	}

	user.Password = ""
	return jwtToken, nil
}

func (u *usecase) SignIn(login string, password string) (*models.AuthToken, error) {
	repUser, err := u.userRep.GetUserByEmail(login)

	if err != nil {
		return nil, errors.Wrap(err, "auth.usecase.SignIn user get error")
	}

	err = u.encryptor.CompareHashAndPassword([]byte(repUser.Password), []byte(password))
	if err != nil {
		return nil, errors.Wrap(err, "auth.usecase.SignIn compare error")
	}

	jwtToken, err := u.tokenProvider.GenerateToken(repUser)
	if err != nil {
		return nil, errors.Wrap(err, "auth.usecase.SignIn token generation error")
	}

	return jwtToken, nil
}

func (u *usecase) BasicAuthorization(token *models.AuthToken) (uint64, string, error) {
	isValid, err := u.tokenProvider.IsTokenValid(token)
	if err != nil {
		return 0, "", errors.Wrap(err, "auth.usecase.Authorization token parse error")
	}

	if !isValid {
		return 0, "", models.ErrInvalidToken
	}

	tokenRole, err := u.tokenProvider.GetRole(token)
	if err != nil {
		return 0, "", errors.Wrap(err, "auth.usecase.Authorization token parse error")
	}

	tokenId, err := u.tokenProvider.GetId(token)
	if err != nil {
		return 0, "", errors.Wrap(err, "auth.usecase.Authorization token parse error")
	}

	user, err := u.userRep.GetUser(tokenId)
	if err != nil || user == nil {
		return 0, "", models.ErrNotFound
	}

	return tokenId, tokenRole, nil
}
