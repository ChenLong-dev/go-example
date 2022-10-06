package callbackplugin

import (
	"context"
	"fmt"
)

type Params struct {
	Name string
}

type Result struct {
	Age int
}

type Callback struct {
	Name string
	Age  int
}

func newCallback() (*Callback, error) {
	return &Callback{
		Name: "Sangfor",
		Age:  20,
	}, nil
}

func (c *Callback) CallBack(ctx context.Context, params Params, result *Result) error {
	fmt.Println("params:", params)
	if c.Name == params.Name {
		result.Age = c.Age
	} else {
		result.Age = 0
	}
	return nil
}
