package repo

import (
	"database/sql"
	"errors"
	"fintechGo/configs"
	"fintechGo/internal/types"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/knadh/koanf"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	conf := configs.NewConfig()
	type args struct {
	}

	tests := []struct {
		Name          string
		input         *types.AuthUser
		args          args
		WantCall      bool
		MockDb        func(sqlmock.Sqlmock)
		expectedError error
	}{

		{
			Name: "error inserting",
			input: &types.AuthUser{
				Name:     "safwan",
				Email:    "safwan@gmail.com",
				Password: "1234",
			},
			args:     args{},
			WantCall: true,
			MockDb: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO auth_users(name,email,password)
				VALUES($1,$2,$3);`)).WithArgs("SAFWAN", "safwan@gmail.com", "123").WillReturnError(errors.New("no records"))
			},
			expectedError: errors.New("something wronng"),
		},
	}
	fmt.Println(tests)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockDB, sqlMock, err := sqlmock.New()
			require.NoError(t, err)
			db := newMockDB(mockDB, conf)
			require.NotNil(t, db)
			u := &Db{
				db: db,
			}
			if tt.MockDb != nil {
				tt.MockDb(sqlMock)
			}
			if tt.WantCall {
				err := u.CreateUser(tt.input)
				if err != nil {
					require.EqualError(t, err, tt.expectedError.Error())
				}

			}
		})
	}

}

func newMockDB(db *sql.DB, logger *koanf.Koanf) *gorm.DB {

	gormHandler, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	if err != nil {
		return nil
	}

	return gormHandler
}
