package payment

import "BalanceRange/internal/db"

type MockDbService struct {
}

func NewMockDbService() db.Storage {
	return &MockDbService{}
}

func (s *MockDbService) GetBalance(clientId int) (res int, err error) {
	return 1, nil
}

func (d *MockDbService) CreateInvoice(amount int, clientID int) {
}

func (d *MockDbService) CreateWithdraw(amount int, clientID int) {
}

func (d *MockDbService) Clear() {
}
