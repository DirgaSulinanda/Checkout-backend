package handler

import (
	"context"
	"reflect"
	"testing"

	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	productRepoMock "github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product/mocks"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	promoRepoMock "github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo/mocks"
	m "github.com/DirgaSulinanda/Checkout-Backend/internal/model/checkoutOutput"
	"github.com/golang/mock/gomock"
)

func Test_checkoutUC_DoCheckout(t *testing.T) {
	type fields struct {
		promoRepoMock   func(ctrl *gomock.Controller) promo.Repository
		productRepomock func(ctrl *gomock.Controller) product.Repository
	}
	type args struct {
		ctx   context.Context
		param m.CheckoutParam
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData m.CheckoutOutput
		wantErr  bool
	}{
		{
			name: "When scanned items are Macbook Pro, Raspberry Pi B, then get free Raspberry Pi B",
			fields: fields{
				promoRepoMock: func(ctrl *gomock.Controller) promo.Repository {
					mock := promoRepoMock.NewMockRepository(ctrl)
					mock.EXPECT().GetActivePromo(gomock.Any()).Return(getPromoMock(), nil)
					return mock
				},
				productRepomock: func(ctrl *gomock.Controller) product.Repository {
					mock := productRepoMock.NewMockRepository(ctrl)
					mock.EXPECT().GetProductDetails(gomock.Any(), []string{"43N23P", "234234"}).Return(
						[]product.Product{
							{SKU: "43N23P", Name: "MacBook Pro", Price: 5399.99, Quantity: 5},
							{SKU: "234234", Name: "Raspberry Pi B", Price: 30, Quantity: 2},
						},
						nil)
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				param: m.CheckoutParam{
					CashierName: "Dirga",
					Products: []m.ProductParam{
						{SKU: "43N23P", Quantity: 1},
						{SKU: "234234", Quantity: 1},
					},
				},
			},
			wantData: m.CheckoutOutput{
				Status:      m.StatusSuccess,
				CashierName: "Dirga",
				Products: []m.Product{
					{
						SKU:           "43N23P",
						Name:          "MacBook Pro",
						Quantity:      1,
						Price:         5399.99,
						TotalPrice:    5399.99,
						OriginalPrice: 5399.99, Promos: nil,
						IsOutOfStock: false,
					},
					{
						SKU:           "234234",
						Name:          "Raspberry Pi B",
						Quantity:      1,
						Price:         30,
						TotalPrice:    0,
						OriginalPrice: 30,
						Promos:        []string{"Free Rasp Pi"},
						IsOutOfStock:  false,
					},
				},
				SubTotal:      5399.99,
				OriginalPrice: 5429.99,
			},
		},
		{
			name: "When scanned items are 3 Google Home, then get free 1 Google Home",
			fields: fields{
				promoRepoMock: func(ctrl *gomock.Controller) promo.Repository {
					mock := promoRepoMock.NewMockRepository(ctrl)
					mock.EXPECT().GetActivePromo(gomock.Any()).Return(getPromoMock(), nil)
					return mock
				},
				productRepomock: func(ctrl *gomock.Controller) product.Repository {
					mock := productRepoMock.NewMockRepository(ctrl)
					mock.EXPECT().GetProductDetails(gomock.Any(), []string{"120P90"}).Return(
						[]product.Product{
							{SKU: "120P90", Name: "Google Home", Price: 49.99, Quantity: 10},
						},
						nil)
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				param: m.CheckoutParam{
					CashierName: "Dirga",
					Products: []m.ProductParam{
						{SKU: "120P90", Quantity: 3},
					},
				},
			},
			wantData: m.CheckoutOutput{
				Status:      m.StatusSuccess,
				CashierName: "Dirga",
				Products: []m.Product{
					{
						SKU:           "120P90",
						Name:          "Google Home",
						Quantity:      3,
						Price:         49.99,
						TotalPrice:    99.97999999999999,
						OriginalPrice: 149.97,
						Promos:        []string{"Buy 2 get 1 free Google Home"},
						IsOutOfStock:  false,
					},
				},
				SubTotal:      99.97999999999999,
				OriginalPrice: 149.97,
			},
		},
		{
			name: "When scanned items are 3 Alexa Speakers, then get discount 10% each",
			fields: fields{
				promoRepoMock: func(ctrl *gomock.Controller) promo.Repository {
					mock := promoRepoMock.NewMockRepository(ctrl)
					mock.EXPECT().GetActivePromo(gomock.Any()).Return(getPromoMock(), nil)
					return mock
				},
				productRepomock: func(ctrl *gomock.Controller) product.Repository {
					mock := productRepoMock.NewMockRepository(ctrl)
					mock.EXPECT().GetProductDetails(gomock.Any(), []string{"A304SD"}).Return(
						[]product.Product{
							{SKU: "A304SD", Name: "Alexa Speaker", Price: 109.5, Quantity: 10},
						},
						nil)
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				param: m.CheckoutParam{
					CashierName: "Dirga",
					Products: []m.ProductParam{
						{SKU: "A304SD", Quantity: 3},
					},
				},
			},
			wantData: m.CheckoutOutput{
				Status:      m.StatusSuccess,
				CashierName: "Dirga",
				Products: []m.Product{
					{
						SKU:           "A304SD",
						Name:          "Alexa Speaker",
						Quantity:      3,
						Price:         109.5,
						TotalPrice:    295.65,
						OriginalPrice: 328.5,
						Promos:        []string{"Alexa Spreaker 10% discount"},
						IsOutOfStock:  false,
					},
				},
				SubTotal:      295.65,
				OriginalPrice: 328.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &checkoutUC{
				promoRepo:   tt.fields.promoRepoMock(ctrl),
				productRepo: tt.fields.productRepomock(ctrl),
			}

			gotData, err := c.DoCheckout(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkoutUC.DoCheckout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("checkoutUC.DoCheckout() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
