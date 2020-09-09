package transaction

import (
	"fiber_api/transaction/models"
	"fiber_api/transaction/repositories"
	"fiber_api/transaction/services"

	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
	"gorm.io/gorm"
)

type creationParams struct {
	Type  string
	Title string
	Value uint64
}

func parseCreationParams(ctx *fiber.Ctx) (*creationParams, error) {
	title := utils.ImmutableString(ctx.Params("title"))
	paramType := utils.ImmutableString(ctx.Params("type"))
	value, convertErr := strconv.ParseUint(
		utils.ImmutableString(ctx.Params("value")),
		10,
		64,
	)

	if convertErr != nil {
		return nil, convertErr
	}

	return &creationParams{
		Type:  paramType,
		Value: value,
		Title: title,
	}, nil
}

// Routes routes to handle user actions
func Routes(app *fiber.App, db *gorm.DB) {
	group := app.Group("transactions")
	transactionRepo := repositories.TransactionRepository{}

	group.Get("/", func(ctx *fiber.Ctx) {
		transactions := transactionRepo.All()
		ctx.Status(200).JSON(transactions)
	})

	group.Post("/", func(ctx *fiber.Ctx) {
		params, convertErr := parseCreationParams(ctx)

		if convertErr != nil {
			ctx.
				Status(400).
				JSON(fiber.Map{
					"error":   "CONVERT_ERROR",
					"message": convertErr.Error(),
				})
			return
		}

		service := services.CreateTransactionService{Repo: transactionRepo}
		createdTransaction, creationErr := service.Execute(services.CreateTransactionDTO{
			Transaction: models.Transaction{
				Name:  params.Title,
				Value: params.Value,
				Type:  params.Type,
			},
		})

		if creationErr != nil {
			ctx.
				Status(400).
				JSON(fiber.Map{
					"error":   "CREATION_ERROR",
					"message": creationErr.Error(),
				})
			return
		}

		ctx.
			Status(200).
			JSON(createdTransaction)
	})
}
