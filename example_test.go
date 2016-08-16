package cqrs

import (
	"context"
	"fmt"

	"github.com/rogerclotet/cqrs/argument"
	"github.com/rogerclotet/cqrs/command"
	"github.com/rogerclotet/cqrs/query"
)

func ExampleRegisterAndHandleCommandAndQuery() {
	cr, _ := command.NewRegistry(
		command.NewRegisteredCommand("increment", incrementHandler()),
	)

	qr, _ := query.NewRegistry(
		query.NewRegisteredQuery("value", valueHandler()),
	)

	ctx := context.TODO()
	_ = cr.Handle(ctx, command.New("increment", argument.Arguments{}))
	_ = cr.Handle(ctx, command.New("increment", argument.Arguments{}))

	v, _ := qr.Handle(ctx, query.New("value", argument.Arguments{}))
	fmt.Println(v)

	// Output: 2
}

var value int

func incrementHandler() command.HandlerFunc {
	return func(_ context.Context, _ argument.Arguments) error {
		value++
		return nil
	}
}

func valueHandler() query.HandlerFunc {
	return func(_ context.Context, _ argument.Arguments) (interface{}, error) {
		return value, nil
	}
}
