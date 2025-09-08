package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/FilledEther20/Reg_Bank/db/mock"
	"github.com/FilledEther20/Reg_Bank/db/sqlc"
	"github.com/FilledEther20/Reg_Bank/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAccountApi(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() sqlc.Account {
	return sqlc.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(5),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
}
