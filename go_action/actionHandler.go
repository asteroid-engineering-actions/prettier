package main

import "context"

type handlerFunc func(actionContext context.Context, actionEvent interface{}) (exitCode int, err error)

func PrettierActionHandler() {}
