package controller

import (
	"github.com/davidgaspardev/gosrv/helpers"
)

type Controller = func(request *helpers.Request, response *helpers.Response)
