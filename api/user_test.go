package api

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	mockdb "github.com/FilledEther20/Reg_Bank/db/mock"
// 	"github.com/FilledEther20/Reg_Bank/db/sqlc"
// 	"github.com/FilledEther20/Reg_Bank/util"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateUserAPI(t *testing.T) {
// 	user, password := randomUser(t)

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"username": user.Username,
// 				"password": password,
// 				"fullname": user.FullName,
// 				"email":    user.Email,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := sqlc.CreateUserParams{
// 					Username:          user.Username,
// 					HashedPassword:    gomock.Any().(string), // we don’t check actual hash here
// 					FullName:          user.FullName,
// 					Email:             user.Email,
// 					PasswordChangedAt: gomock.Any().(time.Time),
// 				}
// 				store.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusCreated, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "InvalidEmail",
// 			body: gin.H{
// 				"username": user.Username,
// 				"password": password,
// 				"fullname": user.FullName,
// 				"email":    "invalid-email",
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				// store should not be called
// 				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "DBError",
// 			body: gin.H{
// 				"username": user.Username,
// 				"password": password,
// 				"fullname": user.FullName,
// 				"email":    user.Email,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(sqlc.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(data))
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func randomUser(t *testing.T) (sqlc.User, string) {
// 	password := util.RandomString(6)
// 	hashedPassword, err := util.HashPassword(password)
// 	require.NoError(t, err)

// 	user := sqlc.User{
// 		Username:          util.RandomOwner(6),
// 		HashedPassword:    hashedPassword,
// 		FullName:          util.RandomOwner(8),
// 		Email:             util.RandomEmail(5),
// 		PasswordChangedAt: time.Now(),
// 	}
// 	return user, password
// }
