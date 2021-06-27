package postgres

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	"github.com/jmoiron/sqlx"
)

func Test_promoRepository_GetActivePromo(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []promo.Promo
		wantErr  bool
	}{
		{
			name: "When query is success, then return promo and no error",
			fields: fields{
				db: func() *sqlx.DB {
					mockDB, mock, _ := sqlmock.New()
					db := sqlx.NewDb(mockDB, "postgres")

					rows := mock.NewRows([]string{"id", "name", "description", "formula", "enabled"}).
						AddRow(1, "Free Rasp Pi", "Each sale of a MacBook Pro comes with a free Raspberry Pi B", "1*{43N23P}=1*{234234}", true).
						AddRow(2, "Buy 2 get 1 free Google Home", "Buy 3 Google Homes for the price of 2", "3*{120P90}=1*{120P90}", true).
						AddRow(3, "Alexa Spreaker 10% discount", "Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa Speakers", "3*{A304SD}=0.1n*{A304SD}", true)

					mock.ExpectQuery(regexp.QuoteMeta(getQuery())).WillReturnRows(rows)
					return db
				}(),
			},
			args: args{ctx: context.Background()},
			wantData: []promo.Promo{
				{ID: 1, Name: "Free Rasp Pi", Description: "Each sale of a MacBook Pro comes with a free Raspberry Pi B", Formula: "1*{43N23P}=1*{234234}", Enabled: true},
				{ID: 2, Name: "Buy 2 get 1 free Google Home", Description: "Buy 3 Google Homes for the price of 2", Formula: "3*{120P90}=1*{120P90}", Enabled: true},
				{ID: 3, Name: "Alexa Spreaker 10% discount", Description: "Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa Speakers", Formula: "3*{A304SD}=0.1n*{A304SD}", Enabled: true},
			},
			wantErr: false,
		},
		{
			name: "When db is nil, then return error",
			fields: fields{db: func() *sqlx.DB {
				return nil
			}()},
			args:     args{ctx: context.Background()},
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
			},
			wantData: nil,
			wantErr:  true,
		},
		{
			name: "When query error, then return error",
			fields: fields{
				db: func() *sqlx.DB {
					mockDB, mock, _ := sqlmock.New()
					db := sqlx.NewDb(mockDB, "postgres")
					mock.ExpectQuery(getQuery()).WillReturnError(errors.New("mock error"))
					return db
				}(),
			},
			args:     args{ctx: context.Background()},
			wantData: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &promoRepository{
				db: tt.fields.db,
			}
			gotData, err := p.GetActivePromo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("promoRepository.GetActivePromo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("promoRepository.GetActivePromo() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
