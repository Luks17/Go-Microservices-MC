package handlers_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/devutils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAccountAPI(t *testing.T) {
	store := devutils.InitMockStore(t)

	account := devutils.RandomNewAccount()

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server, recorder := devutils.NewMockServer(store)

	url := fmt.Sprintf("/accounts/%d", account.ID)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
}
