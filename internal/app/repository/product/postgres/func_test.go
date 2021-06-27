package postgres

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	"github.com/jmoiron/sqlx"
)

func Test_productRepository_GetProductDetails(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx context.Context
		sku []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []product.Product
		wantErr  bool
	}{
		{
			name: "When query is success, then return products and no error",
			fields: fields{
				db: func() *sqlx.DB {
					mockDB, mock, _ := sqlmock.New()
					db := sqlx.NewDb(mockDB, "postgres")

					rows := mock.NewRows([]string{"sku", "name", "price", "quantity"}).
						AddRow("A304SD", "Alexa Speaker", 109.50, 10).
						AddRow("234234", "Raspberry Pi B", 30, 2)

					q, _, _ := getQuery([]string{"234234", "A304SD"})
					q = db.Rebind(q)
					mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("234234", "A304SD").WillReturnRows(rows)
					return db
				}(),
			},
			args: args{ctx: context.Background(), sku: []string{"234234", "A304SD"}},
			wantData: []product.Product{
				{SKU: "A304SD", Name: "Alexa Speaker", Price: 109.5, Quantity: 10},
				{SKU: "234234", Name: "Raspberry Pi B", Price: 30, Quantity: 2},
			},
			wantErr: false,
		},
		{
			name: "When db is nil, then return error",
			fields: fields{db: func() *sqlx.DB {
				return nil
			}()},
			args:     args{ctx: context.Background(), sku: []string{"234234", "A304SD"}},
			wantData: nil,
			wantErr:  true,
		},
		{
			name: "When ctx is timeout, then return error",
			fields: fields{db: func() *sqlx.DB {
				return nil
			}()},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel()
					return ctx
				}(),
				sku: []string{"234234", "A304SD"}},
			wantData: nil,
			wantErr:  true,
		},
		{
			name: "When query error, then return error",
			fields: fields{
				db: func() *sqlx.DB {
					mockDB, mock, _ := sqlmock.New()
					db := sqlx.NewDb(mockDB, "postgres")
					q, _, _ := getQuery([]string{"234234", "A304SD"})
					q = db.Rebind(q)
					mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("234234", "A304SD").WillReturnError(errors.New("mock error"))
					return db
				}(),
			},
			args:     args{ctx: context.Background(), sku: []string{"234234", "A304SD"}},
			wantData: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &productRepository{
				db: tt.fields.db,
			}
			gotData, err := p.GetProductDetails(tt.args.ctx, tt.args.sku)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.GetProductDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("productRepository.GetProductDetails() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
