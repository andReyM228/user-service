package car_trading

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"testing"
	"user_service/internal/domain"
	"user_service/internal/repository/balances"
	"user_service/internal/repository/transactions"
)

func initDatabase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "tx_service")

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func TestService_SendCoins(t *testing.T) {
	db := initDatabase()

	type fields struct {
		balanceRepo     balances.Repository
		transactionRepo transactions.Repository
	}
	type args struct {
		tx domain.Transactions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful transfer",
			fields: fields{
				balanceRepo:     balances.NewRepository(db, logrus.New()),
				transactionRepo: transactions.NewRepository(db, logrus.New()),
			},
			args: args{domain.Transactions{
				UserIDFrom: 1,
				UserIDTo:   2,
				Amount:     10,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				balanceRepo:     tt.fields.balanceRepo,
				transactionRepo: tt.fields.transactionRepo,
			}
			if err := s.SendCoins(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("SendCoins() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
