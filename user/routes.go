package user

import (
	// "fiber_api/user/models"
	"fiber_api/user/repositories"
	// "fiber_api/user/services"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)


type userRegisterParams struct {
	Name string	`json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func getUserRegisterParams(ctx *fiber.Ctx) (*userRegisterParams, error) {
	var params userRegisterParams
	if err := ctx.BodyParser(params); err != nil {
		return nil, err
	}

	return &params, nil
} 

// Routes routes to handle user actions
func Routes(router fiber.Router, db *gorm.DB) {
	transactionRepo := repositories.UserRepoitory{}.
		SetConnection(db)

	router.Post("/", func(ctx *fiber.Ctx) {
		params, paramsError := getUserRegisterParams(ctx)

		if paramsError != nil {
			ctx.Status(400).
				JSON(fiber.Map{
					"error": "INVALID_PARAMS",
					"message": paramsError.Error(),
				})
		}

		
	})
}
