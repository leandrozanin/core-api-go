package server

import (
	"github.com/gofiber/fiber/v2"
)

func (Server *Api) Get(path string, handlers ...fiber.Handler) fiber.Router {
	return Server.App.Get(path, handlers...)
}

func (Server *Api) Head(path string, handlers ...fiber.Handler) fiber.Router {
	return Server.App.Head(path, handlers...)
}

func (Server *Api) Post(path string, handlers ...fiber.Handler) fiber.Router {
	return Server.App.Post(path, handlers...)
}

func (Server *Api) Put(path string, handlers ...fiber.Handler) fiber.Router {
	return Server.App.Put(path, handlers...)
}

func (Server *Api) Delete(path string, handlers ...fiber.Handler) fiber.Router {
	return Server.App.Delete(path, handlers...)
}
