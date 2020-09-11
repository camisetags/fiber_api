package transaction

import (
	"fiber_api/transaction/models"
	"fiber_api/transaction/repositories"
	"fiber_api/transaction/services"

	"github.com/gofiber/fiber"
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
func Routes(router fiber.Router, db *gorm.DB) {
	transactionRepo := repositories.TransactionRepository{}.
		SetConnection(db)

	router.Get("/", func(ctx *fiber.Ctx) {
		var transactions []fiber.Map

		for _, transaction := range transactionRepo.All() {
			transactions = append(transactions, fiber.Map{
				"id": transaction.ID,
				"name": transaction.Name,
				"type": transaction.Type,
				"value": transaction.Value,
			})
		}

		ctx.Status(200).JSON(fiber.Map{
			"transactions": transactions,
			"balance": transactionRepo.GetBalance(),
		})
	})

	router.Post("/", func(ctx *fiber.Ctx) {
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
		
		service := services.CreateTransactionService{
			Repo: transactionRepo,
		}
		newTransaction := models.Transaction{
			Name:  params.Title,
			Value: params.Value,
			Type:  params.Type,
		}
		createdTransaction, creationErr := service.Execute(newTransaction)

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
			JSON(fiber.Map{
				"id": createdTransaction.ID,
				"name": createdTransaction.Name,
				"type": createdTransaction.Type,
				"value": createdTransaction.Value,
			})
	})
}
