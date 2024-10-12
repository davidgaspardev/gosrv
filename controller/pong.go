package controller

import "github.com/davidgaspardev/gosrv/helpers"

func PongController() Controller {
	return func(req *helpers.Request, res *helpers.Response) {
		res.OkText("pong")
	}
}
