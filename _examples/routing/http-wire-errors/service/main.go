package main

import (
	"context"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/x/errors"
)

func main() {
	app := iris.New()

	service := new(myService)
	app.Post("/", createHandler(service))
	app.Get("/", listHandler(service))
	app.Delete("/{id:string}", deleteHandler(service))

	app.Listen(":8080")
}

func createHandler(service *myService) iris.Handler {
	return func(ctx iris.Context) {
		// What it does?
		// 1. Reads the request body and binds it to the CreateRequest struct.
		// 2. Calls the service.Create function with the given request body.
		// 3. If the service.Create returns an error, it sends an appropriate error response to the client.
		// 4. If the service.Create returns a response, it sets the status code to 201 (Created) and sends the response as a JSON payload to the client.
		//
		// Useful for create operations.
		errors.Create(ctx, service.Create)
	}
}

func listHandler(service *myService) iris.Handler {
	return func(ctx iris.Context) {
		// What it does?
		// 1. If the 3rd variadic (optional) parameter is empty (not our case here), it reads the request body and binds it to the ListRequest struct,
		// otherwise (our case) it calls the service.List function directly with the given input parameter (empty ListRequest struct value in our case).
		// 2. Calls the service.List function with the ListRequest value.
		// 3. If the service.List returns an error, it sends an appropriate error response to the client.
		// 4. If the service.List returns a response, it sets the status code to 200 (OK) and sends the response as a JSON payload to the client.
		//
		// Useful for get single, fetch multiple and search operations.
		errors.OK(ctx, service.List, ListRequest{})
	}
}

func deleteHandler(service *myService) iris.Handler {
	return func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		// What it does?
		// 1. Calls the service.Delete function with the given input parameter.
		// 2. If the service.Delete returns an error, it sends an appropriate error response to the client.
		// 3.If the service.Delete doesn't return an error then it sets the status code to 204 (No Content) and
		// sends the response as a JSON payload to the client.
		// errors.NoContent(ctx, service.Delete, id)
		// OR:
		// 1. Calls the service.DeleteWithFeedback function with the given input parameter.
		// 2. If the service.DeleteWithFeedback returns an error, it sends an appropriate error response to the client.
		// 3. If the service.DeleteWithFeedback returns true, it sets the status code to 204 (No Content).
		// 4. If the service.DeleteWithFeedback returns false, it sets the status code to 304 (Not Modified).
		//
		// Useful for update and delete operations.
		errors.NoContentOrNotModified(ctx, service.DeleteWithFeedback, id)
	}
}

type (
	myService struct{}

	CreateRequest struct {
		Fullname string
	}

	CreateResponse struct {
		ID        string
		Firstname string
		Lastname  string
	}
)

func (s *myService) Create(ctx context.Context, in CreateRequest) (CreateResponse, error) {
	arr := strings.Split(in.Fullname, " ")
	firstname, lastname := arr[0], arr[1]
	id := "test_id"

	resp := CreateResponse{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
	}
	return resp, nil // , errors.New("create: test error")
}

type ListRequest struct {
}

func (s *myService) List(ctx context.Context, in ListRequest) ([]CreateResponse, error) {
	resp := []CreateResponse{
		{
			ID:        "test-id-1",
			Firstname: "test first name 1",
			Lastname:  "test last name 1",
		},
		{
			ID:        "test-id-2",
			Firstname: "test first name 2",
			Lastname:  "test last name 2",
		},
		{
			ID:        "test-id-3",
			Firstname: "test first name 3",
			Lastname:  "test last name 3",
		},
	}

	return resp, nil //, errors.New("list: test error")
}

func (s *myService) Delete(ctx context.Context, id string) error {
	return nil // errors.New("delete: test error")
}

func (s *myService) DeleteWithFeedback(ctx context.Context, id string) (bool, error) {
	return true, nil // false, errors.New("delete: test error")
}
