package router

import (
	"awesomeProject/router/handlers"
	"context"
	"net/http"
)

var path = "/do-things"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Routes(ctx context.Context) []route {
	return []route{
		{
			Name:        "DoThings",
			Method:      http.MethodPost,
			Pattern:     path,
			HandlerFunc: handlers.DoThings(ctx),
		},
	}
}
