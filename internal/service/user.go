package service

import (
	"errors"
	"golang-ecommerce/entity"
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/model"
	"golang-ecommerce/pkg/bcrypt"
	"golang-ecommerce/pkg/jwt"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(param *model.UserRegister) error
	Login(param model.UserLogin) (model.LoginResponse, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserServcice struct {
	UserRepository repository.IUserRepository
	CartRepository repository.ICartRepository
	Bcrypt         bcrypt.Interface
	JwtAuth        jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, cartRepository repository.ICartRepository) IUserService {
	return &UserServcice{
		UserRepository: userRepository,
		Bcrypt:         bcrypt,
		JwtAuth:        jwtAuth,
		CartRepository: cartRepository,
	}
}

func (us *UserServcice) Register(param *model.UserRegister) error {
	hash, err := us.Bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	user := &entity.User{
		UserID:   id,
		Name:     param.Name,
		Email:    param.Email,
		Password: hash,
		RoleID:   2,
	}

	_, err = us.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	cart := entity.Cart{
		UserID: id,
	}

	err = us.CartRepository.CreateCart(&cart)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserServcice) Login(param model.UserLogin) (model.LoginResponse, error) {
	var result model.LoginResponse

	user, err := us.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})

	if err != nil {
		return result, err
	}

	err = us.Bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := us.JwtAuth.CreateJWTToken(user.UserID)
	if err != nil {
		return result, errors.New("failed to create token")
	}

	result.UserID = user.UserID
	result.Token = token
	result.RoleID = user.RoleID

	return result, nil
}

func (us *UserServcice) GetUser(param model.UserParam) (*entity.User, error) {
	return us.UserRepository.GetUser(param)
}
