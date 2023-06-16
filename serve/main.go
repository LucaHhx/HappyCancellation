package main

import (
	"github.com/LucaHhx/nano"
	"github.com/LucaHhx/nano/serialize/json"
	"net/http"
	"serve/service"
)

func main() {
	nano.Listen(":4321",
		nano.WithIsWebsocket(true),
		nano.WithClientAddr(":4321"),
		nano.WithWSPath("/nano"),
		nano.WithComponents(service.Services),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
		nano.WithDebugMode(),
	)
}
