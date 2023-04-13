package service

import "github.com/LucaHhx/nano/component"

var (
	Services = &component.Components{}
	serve    = NewServe()
)

func init() {
	Services.Register(serve)
}
