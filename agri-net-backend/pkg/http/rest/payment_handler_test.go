package rest

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
)

func TestPaymentHandler_Payin(t *testing.T) {
	type fields struct {
		Service payment.IPaymentService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phan := &PaymentHandler{
				Service: tt.fields.Service,
			}
			phan.Payin(tt.args.c)
		})
	}
}
