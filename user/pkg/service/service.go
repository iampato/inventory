package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iampato/inventory/user/pkg/entities"
	util "github.com/iampato/inventory/user/pkg/utils"
	"gorm.io/gorm"
)

// UserService describes the service.
type UserService interface {
	// Add your methods here
	CreateRole(ctx context.Context, roleName string) (err error)
	GetAllRoles(ctx context.Context) (roles []entities.UserRole)
	Create(ctx context.Context, firstName string, lastName string, emailAddress string, phone string, password string, roleId int, dob string) (user entities.UserModel, err error)
	GetAllUsers(ctx context.Context, limit int, offset int) (users []entities.UserModel, total int64, page int64, lastPage int64)
	// updateUser
	// deleteUser
	// login user

}

type basicUserService struct {
	db *gorm.DB
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService(db *gorm.DB) UserService {
	// Migrate the schema
	db.AutoMigrate(&entities.UserRole{}, &entities.UserModel{})

	// return
	return &basicUserService{
		db: db,
	}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware, db *gorm.DB) UserService {
	var svc UserService = NewBasicUserService(db)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserService) CreateRole(ctx context.Context, roleName string) error {
	ent := entities.UserRole{
		RoleName: roleName,
	}
	_, err := entities.CreateRole(b.db, ent)
	if err != nil {
		return err
	}
	return nil
}

// Create implements UserService
func (b *basicUserService) Create(ctx context.Context, firstName string, lastName string, emailAddress string, phone string, password string, roleId int, dob string) (entities.UserModel, error) {
	formattedDob, err := time.Parse("2006-01-02T15:04:05", dob)
	if err != nil {
		fmt.Printf("Error parsing time: %v", err)
		return entities.UserModel{}, err
	}
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return entities.UserModel{}, err
	}
	ent := entities.UserModel{
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Phone:        phone,
		RoleId:       uint(roleId),
		Password:     hashedPassword,
		DOB:          formattedDob,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	userId, err := entities.CreateUser(b.db, ent)
	if err != nil {
		return entities.UserModel{}, err
	}
	// send a notification of account creation

	users := entities.FetchUserByColumn(b.db, "id", userId)
	if len(users) > 0 {
		return users[0], nil
	}
	return entities.UserModel{}, errors.New("USER created by failed to fetch user")
}

// FetchAllRoles implements UserService
func (b *basicUserService) GetAllRoles(ctx context.Context) []entities.UserRole {
	roles := entities.FetchAllRoles(b.db)
	return roles
}

// FetchAllUsers implements UserService
func (b *basicUserService) GetAllUsers(ctx context.Context, limit int, offset int) ([]entities.UserModel, int64, int64, int64) {
	users := entities.FetchAllUsers(b.db, limit, offset)
	count := entities.CountUsers(b.db)

	return users, count, int64(offset), int64(float64(count / int64(limit)))
}
