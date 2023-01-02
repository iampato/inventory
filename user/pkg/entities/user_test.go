package entities

import (
	"reflect"
	"testing"

	config "github.com/iampato/inventory/user/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// connect db
	// connect to postgresql
	db = config.ConnectToDb()
	defer func() {
		// clean db
	}()
}

func TestCreateRole(t *testing.T) {
	type args struct {
		db   *gorm.DB
		role UserRole
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "Create roles",
			args: args{
				db: db,
				role: UserRole{
					RoleName: "ADMIN",
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateRole(tt.args.db, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchAllRoles(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want []UserRole
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchAllRoles(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchAllRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}
