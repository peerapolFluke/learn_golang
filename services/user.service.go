package services

import (
	"context"
	"fmt"
	"seamoor"
	"seamoor/cors/injection"
	"seamoor/cors/utils"
	"seamoor/ent"
	"seamoor/ent/user"
	"time"
)

func CreateUser(ctx context.Context, input *seamoor.CreateUserInput) (*ent.User, error) {
	//check email duplicate
	isAlreadyHaveUser := injection.DB.User.Query().
		Where(user.EmailEQ(input.Email)).
		ExistX(ctx)
	if isAlreadyHaveUser {
		return nil, utils.CustomGqlError(ctx, "This Email Already in use", 400)
	}

	//randomPassword := utils.RandomPassword(8)
	hashedPassword, errHash := utils.HashPassword("12345678")
	if errHash != nil {
		return nil, utils.CustomGqlError(ctx, errHash.Error(), 500)
	}
	fmt.Println("12345678", hashedPassword)
	user, err := injection.DB.User.Create().
		SetName(input.Name).
		SetSurname(input.Surname).
		SetPosition(input.Position).
		SetEmail(input.Email).
		SetRole(input.Role).
		SetPassword(hashedPassword).
		Save(ctx)
	if err != nil {
		return nil, utils.CustomGqlError(ctx, err.Error(), 500)
	}
	return user, nil
}

func Login(ctx context.Context, input *seamoor.LoginInput) (*seamoor.Credential, error) {
	incorrectError := utils.CustomGqlError(ctx, "Incorrect email or password", 400)

	u, err := injection.DB.User.Query().
		Where(user.EmailEQ(input.Email)).
		Only(ctx)
	if err != nil {
		return nil, incorrectError
	}
	checkPassword := utils.CheckPassword(input.Password, u.Password)
	if checkPassword != nil {
		return nil, incorrectError
	}
	injection.DB.User.Update().SetLastLoginAt(time.Now()).Where(user.IDEQ(u.ID)).Save(ctx)

	jwt, errJwt := utils.JwtGenerate(u.ID, u.Email)
	if errJwt != nil {
		return nil, errJwt
	}

	return &seamoor.Credential{
		AccessToken: jwt,
	}, nil
}

func GetUserById(ctx context.Context, id int) (*ent.User, error) {
	userData, err := injection.DB.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
