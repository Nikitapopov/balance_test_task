package payment

import (
	"testing"
)

func TestStart(t *testing.T) {
	mockDbService := NewMockDbService()
	paymentService := NewPaymentsService(&mockDbService)
	err := paymentService.Start()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
