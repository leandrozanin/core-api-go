package server

import (
	"context"
	"fmt"

	"github.com/leandrozanin/core-api-go/middleware/header"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

type AdapterName string

const (
	ApiFiber  AdapterName = "fiber"
	ApiLambda AdapterName = "lambda"
)

var fiberLambda *fiberadapter.FiberLambda

type Api struct {
	Name   string
	Domain string
	Adp    AdapterName
	App    *fiber.App
}

func (Server *Api) InitApi() {
	Server.App = fiber.New()
	Server.App.Use(header.New(header.Config{
		Name:   &Server.Name,
		Domain: &Server.Domain,
	}))

	Server.App.Hooks().OnRoute(func(r fiber.Route) error {
		fmt.Print("Method: " + r.Method + " " + r.Path + "\n")

		return nil
	})

	Server.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func (Server *Api) Start(port string) {

	if Server.Adp == ApiLambda {

		fiberLambda = fiberadapter.New(Server.App)
		lambda.Start(HandlerLambda)
		fmt.Println("Iniciando Lambda")

	} else {

		err := Server.App.Listen(port)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fiber porta: ", port)
		}
	}

}

func HandlerLambda(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, request)
}
