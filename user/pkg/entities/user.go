package entities

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid()"` // db func
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"column:last_name; not null"`
	// FullName     string `gorm:"->;type:GENERATED ALWAYS AS (concat(first_name,' ',last_name));"`
	EmailAddress string `gorm:"unique"`
	Password     string
	Phone        string `gorm:"unique"`
	RoleId       uint
	Role         UserRole `gorm:"foreignKey:role_id"`
	DOB          time.Time
	CreatedAt    time.Time `gorm:"index"`
	LastLogin    time.Time `gorm:"index"`
	UpdatedAt    time.Time
	IsDeleted    bool      `gorm:"default:false"`
	DeletedAt    time.Time `gorm:"index"`
	// gorm.Model
}

type UserRole struct {
	ID       uint `gorm:"primaryKey,autoIncrement"`
	RoleName string
}

/*
ROLES
*/
func CreateRole(db *gorm.DB, role UserRole) (uint, error) {
	result := db.Create(&role)
	if result.Error != nil {
		return 0, result.Error
	}
	return role.ID, nil
}

func FetchAllRoles(db *gorm.DB) []UserRole {
	var roles []UserRole
	result := db.Find(&roles)
	if result.Error != nil {
		roles = []UserRole{} // empty slice
	}
	if result.RowsAffected == 0 {
		roles = []UserRole{} // empty slice
	}
	return roles
}

/*
USER
*/
func CreateUser(db *gorm.DB, user UserModel) (string, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.ID, nil
}

func FetchAllUsers(db *gorm.DB, limit int, offset int) []UserModel {
	var users []UserModel
	result := db.Preload("Role").Limit(limit).Offset((offset - 1) * limit).Find(&users)
	fmt.Println(len(users))
	if result.Error != nil {
		users = []UserModel{}
	}
	if result.RowsAffected == 0 {
		users = []UserModel{} // empty slice
	}
	return users
}

func CountUsers(db *gorm.DB) int64 {
	var total int64
	var users []UserModel
	result := db.Find(&users).Count(&total)
	fmt.Println(len(users))
	if result.Error != nil {
		total = 0
	}
	if result.RowsAffected == 0 {
		total = 0
	}
	return total
}

func FetchUserByColumn(db *gorm.DB, column string, value interface{}) []UserModel {
	var users []UserModel
	filter := fmt.Sprintf("%s = ?", column)
	condiction := fmt.Sprintf("%v", value)
	result := db.Preload("Role").Where(filter, condiction).Find(&users)
	if result.Error != nil {
		users = []UserModel{}
	}
	if result.RowsAffected == 0 {
		users = []UserModel{} // empty slice
	}
	return users
}

func UpdateUser(db *gorm.DB, id string, update map[string]interface{}) error {
	result := db.Model(UserModel{}).Where("id = ?", id).Updates(update)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func DeleteUser(db *gorm.DB, id string) error {
	result := db.Model(UserModel{}).Where("id = ?", id).Updates(map[string]interface{}{"IsDeleted": true, "DeletedAt": time.Now()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
