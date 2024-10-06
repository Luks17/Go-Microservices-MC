package devutils

import (
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db"
	mockdb "github.com/Luks17/Go-Microservices-MC/db/mock"
	"go.uber.org/mock/gomock"
)

func InitMockStore(t *testing.T) *mockdb.MockStore {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)
	db.DBStore = store

	return store
}
