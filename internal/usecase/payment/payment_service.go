package payment

import (
	"context"
	"fmt"
	"sync"
	"time"

	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/payment"
	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/payment"
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
	totalPaymenToProcess := 2

	paymentProcess := make([]entity.PaymentProcess, totalPaymenToProcess)

	var wg sync.WaitGroup

	for x := 0; x < totalPaymenToProcess; x++ {
		paymentProcess[x] = entity.PaymentProcess{
			OrderId: int64(x),
			OutstandingAmount: float64(outstandingAmount),
			PaidAmount: float64(x * 1000),
		}
		
		wg.Add(2)
		go i.updateBalance(&paymentProcess[x], &wg)
		go i.updateBalance(&paymentProcess[x], &wg)
	}

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
	time.Sleep(3 * time.Second)
}

