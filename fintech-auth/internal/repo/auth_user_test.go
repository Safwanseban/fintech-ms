package repo

import (
	"errors"
	"fintechGo/internal/types"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {

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
	// for _, tt := range tests {
	// 	t.Run(tt.Name, func(t *testing.T) {
	// 		mockDB, sqlMock, err := sqlmock.New()
	// 		require.NoError(t, err)

	// 		db, err := gorm.Open( gorm.Config{
	// 			ConnPool: mockDB,
	// 		}, mockDB)

	// 		u := &Db{
	// 			db: db,
	// 		}
	// 		if tt.MockDb != nil {
	// 			tt.MockDb(sqlMock)
	// 		}
	// 		if tt.WantCall {
	// 			err := u.CreateUser(tt.input)
	// 			if err != nil {
	// 				require.EqualError(t, err, tt.expectedError.Error())
	// 			}

	// 		}
	// 	})
	// }

}
