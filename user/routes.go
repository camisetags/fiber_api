package user

import (
	"fiber_api/user/repositories"
	"fiber_api/user/services"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userRegisterParams struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func getUserRegisterParams(ctx *fiber.Ctx) (*userRegisterParams, error) {
	params := new(userRegisterParams)
	if err := ctx.BodyParser(params); err != nil {
		return nil, err
	}

	return params, nil
}

// Routes routes to handle user actions
func Routes(router fiber.Router, db *gorm.DB) {
	userRepo := repositories.UserRepoitory{}.
		SetConnection(db)

	router.Post("/", func(ctx *fiber.Ctx) error {
		params, paramsError := getUserRegisterParams(ctx)

		if paramsError != nil {
			return ctx.Status(400).
				JSON(fiber.Map{
					"error":   "INVALID_PARAMS",
					"message": paramsError.Error(),
				})
		}

		service := services.RegisterUserService{Repo: userRepo}
		newUser, creationError := service.Execute(
			services.UserFields{
				Name:     params.Name,
				Email:    params.Email,
				Password: params.Password,
			},
			params.PasswordConfirmation,
		)

		if creationError != nil {
			return ctx.Status(400).
				JSON(fiber.Map{
					"error":   "USER_CREATION",
					"message": creationError.Error(),
				})
		}

		return ctx.Status(200).
			JSON(fiber.Map{
				"name":  newUser.Name,
				"email": newUser.Email,
			})
	})
}
