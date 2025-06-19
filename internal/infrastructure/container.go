package infrastructure

import (
	"simple-withdraw-api/internal/balance"
	"simple-withdraw-api/internal/config"
	"simple-withdraw-api/internal/domain"
	"simple-withdraw-api/internal/user"
	"simple-withdraw-api/internal/withdrawal"
	"simple-withdraw-api/pkg/xlogger"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var (
	cfg config.Config

	UserRepository domain.UserRepository
	BalanceRepository domain.BalanceRepository
	WithdrawalRepository domain.WithdrawalRepository

	UserSvc domain.UserService
	BalanceSvc domain.BalanceService
	WithdrawalSvc domain.WithdrawalService
)

func init (){
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err = env.Parse(&cfg); err != nil {
		panic(err)
	}
	xlogger.Setup(cfg)
	xlogger.Logger.Info().Msgf("Config: %+v", cfg)
	dbSetup()

	UserRepository = user.NewMysqlRepository(db)
	BalanceRepository = balance.NewBalanceRepository(db)
	WithdrawalRepository = withdrawal.NewWithdrawalRepository(db)

	BalanceSvc = balance.NewBalanceService(BalanceRepository)
	UserSvc = user.NewUserService(UserRepository,BalanceSvc)
	WithdrawalSvc = withdrawal.NewWithdrawalService(UserSvc,BalanceSvc,WithdrawalRepository)
}