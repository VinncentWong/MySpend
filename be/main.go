package main

import (
	"log"
	bcCache "module/app/budget/cache"
	bHandler "module/app/budget/handler"
	bRepo "module/app/budget/repository"
	bUsecase "module/app/budget/usecase"
	fHandler "module/app/financial_account/handler"
	fRepo "module/app/financial_account/repository"
	fUsecase "module/app/financial_account/usecase"
	pHandler "module/app/payment/handler"
	pRepo "module/app/payment/repository"
	pUsecase "module/app/payment/usecase"
	"module/app/user/handler"
	"module/app/user/repository"
	"module/app/user/usecase"
	"module/infrastructure"
	"module/rest"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil && os.Getenv("device") != "docker" {
		log.Fatalf("can't load env properties, %s", err.Error())
		return
	}

	err = infrastructure.InitDb()
	if err != nil {
		log.Fatal("exception on database initialization")
		return
	}

	err = infrastructure.Migrate()
	if err != nil {
		log.Fatal("exception when migrating")
		return
	}

	err = infrastructure.InitStorage()
	if err != nil {
		log.Fatal("exception when init storage")
		return
	}

	result := infrastructure.InitRedis()
	if result != nil {
		log.Fatal("exception when init redis")
		return
	}

	err = infrastructure.InitCache()
	if err != nil {
		log.Fatal("exception when init cache")
		return
	}
	// Routing Initialization
	r := gin.Default()

	// Use cors
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}))

	// Make bucket
	err = infrastructure.CreateBucket("paymentproof")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// User App
	userRepo := repository.NewUserRepository()
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Financial Account App
	financialAccountRepo := fRepo.NewFinancialAccountRepository()
	financialAccountService := fUsecase.NewFinancialAccountService(financialAccountRepo)
	financialAccountHandler := fHandler.NewFinancialAccountHandler(financialAccountService)

	// Payment History App
	paymentRepo := pRepo.NewPaymentRepository()
	paymentService := pUsecase.NewPaymentService(paymentRepo)
	paymentHandler := pHandler.NewPaymentHandler(paymentService)

	// Budget History App
	budgetRepo := bRepo.NewBudgetRepository()
	budgetCache := bcCache.NewBudgetCache()
	budgetService := bUsecase.NewBudgetService(budgetRepo, budgetCache)
	budgetHandler := bHandler.NewBudgetHandler(budgetService)

	rest.UserRouting(r, userHandler)
	rest.FinancialAccountRouting(r, financialAccountHandler)
	rest.CheckHealthRouting(r)
	rest.PaymentHistoryRouting(r, paymentHandler)
	rest.BudgetRouting(r, budgetHandler)

	r.Run(":8000")
}
