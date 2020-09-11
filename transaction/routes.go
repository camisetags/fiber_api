package transaction

import (
	// "fmt"
	
	"fiber_api/transaction/models"
	"fiber_api/transaction/repositories"
	"fiber_api/transaction/services"

	// "strconv"

	"github.com/gofiber/fiber"
	// "github.com/gofiber/utils"
	"gorm.io/gorm"
)

type creationParams struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Value uint64 `json:"value"`
}

func parseCreationParams(ctx *fiber.Ctx) (*creationParams, error) {
	params := new(creationParams)
	if err := ctx.BodyParser(params); err != nil {
		return nil, err
	}

	return params, nil
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
			JSON(*createdTransaction)
	})
}
