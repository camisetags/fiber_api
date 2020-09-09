package transaction

import (
	"fiber_api/transaction/models"
	"fiber_api/transaction/services"
	"fiber_api/transaction/repositories"
	
	"strconv"

	"github.com/gofiber/utils"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

// Routes routes to handle user actions
func Routes(app *fiber.App, db *gorm.DB) {
	group := app.Group("transactions")
	transactionRepo := repositories.TransactionRepository{}


	group.Get("/", func(ctx *fiber.Ctx) {
		transactions := transactionRepo.All()
		ctx.Status(200).JSON(transactions)
	})

	group.Post("/", func(ctx *fiber.Ctx) {
		title := utils.ImmutableString(ctx.Params("title"))
		paramType := utils.ImmutableString(ctx.Params("type"))
		value, convertErr := strconv.ParseUint(
			utils.ImmutableString(ctx.Params("value")), 
			10, 
			64,
		)

		if convertErr != nil {
			ctx.
				Status(400).
				JSON(fiber.Map{
					"error": "CONVERT_ERROR",
					"message": convertErr.Error(),
				})
			return
		}

		service := services.CreateTransactionService{
			Repo: transactionRepo,
		}
		createdTransaction, creationErr := service.Execute(services.CreateTransactionDTO{
			Transaction: models.Transaction{
				Name: title,
				Value: value,
				Type: paramType,
			},
		})

		if creationErr != nil {
			ctx.
				Status(400).
				JSON(fiber.Map{
					"error": "INVALID_TRANSACTION_TYPE",
					"message": creationErr.Error(),
				})
			return
		}

		ctx.
			Status(200).
			JSON(createdTransaction)
	})
}
