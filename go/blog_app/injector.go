//go:build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(
		//tag
		NewGreeter,
		NewMessage
	)
	return Event{}
}
