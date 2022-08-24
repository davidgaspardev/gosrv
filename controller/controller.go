package controller

import (
	"gosrv/helpers"
)

type Controller = func(request *helpers.Request, response *helpers.Response)
