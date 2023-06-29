package payment

import (
	"BalanceRange/internal/db"
	"fmt"
	"math/rand"
	"sync"
)

type PaymentService struct {
	dbService db.Storage
}

type PaymentServiceInterface interface {
	Start() error
}

func NewPaymentsService(dbService *db.Storage) PaymentServiceInterface {
	return &PaymentService{
		dbService: *dbService,
	}
}

func (s *PaymentService) Start() error {
	for i := 0; i < 30; i++ {
		go s.dbService.CreateInvoice(randInRange(10, 50), randInRange(1, 3))
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.dbService.CreateWithdraw(randInRange(40, 100), randInRange(1, 3))
		}()
	}
	wg.Wait()

	balance1, err := s.dbService.GetBalance(1)
	if err != nil {
		return err
	}
	fmt.Println(balance1)

	balance2, err := s.dbService.GetBalance(2)
	if err != nil {
		return err
	}
	fmt.Println(balance2)

	balance3, err := s.dbService.GetBalance(3)
	if err != nil {
		return err
	}
	fmt.Println(balance3)

	return nil
}

func randInRange(min, max int) int {
	return int(float64(rand.Intn(max-min+1) + min))
}
