package payment

import (
	"context"
	"fmt"
	"sync"
	"time"

	entity "lunba-e-commerce/internal/domain/entity/payment"
	repository "lunba-e-commerce/internal/domain/repository/payment"
)

type PaymentService interface {
	Process(ctx context.Context, orderId int64) (*entity.Payment, error)
}

type paymentServiceImpl struct {
	repo repository.PaymentRepository
	mu sync.Mutex
}

func New(r repository.PaymentRepository) PaymentService {
	if r == nil {
		panic("payment repository must be implemented.")
	}
	return &paymentServiceImpl{
		repo: r,
	}
}

func (i *paymentServiceImpl) Process(ctx context.Context, orderId int64) (*entity.Payment, error) {
	outstandingAmount := 6000

	var wg sync.WaitGroup

	paymentProcess := entity.PaymentProcess{
		OrderId: orderId,
		OutstandingAmount: float64(outstandingAmount),
		PaidAmount: float64(orderId * 1000),
	}

	wg.Add(2)
	go i.updateBalance(&paymentProcess, &wg)
	go i.updateBalance(&paymentProcess, &wg)

	wg.Wait()

	return nil, nil
}

func (i *paymentServiceImpl) updateBalance(e *entity.PaymentProcess, wg *sync.WaitGroup) {
	defer wg.Done()
	
	i.mu.Lock()
	defer i.mu.Unlock()

	// fmt.Printf("Processing payment with orderId: %d", e.OrderId)

	e.Balance = e.OutstandingAmount - e.PaidAmount
	e.RandomAmount += e.Balance

	fmt.Printf("%v\n", e)
	time.Sleep(1 * time.Second)
}

