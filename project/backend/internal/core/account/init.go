package account

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	account := api.V2.Group("/account", middlewares.Auth)
	{
		account.GET(
			"/",
			list,
		)
		account.POST(
			"/",
			create,
		)
		account.PATCH(
			"/",
			update,
		)
		account.DELETE(
			"/",
			delete,
		)
	}
}