package main

import (
	"github.com/LucaHhx/nano"
	"github.com/LucaHhx/nano/serialize/json"
	"net/http"
	"serve/serve"
)

func main() {
	nano.Listen(":8080",
		nano.WithIsWebsocket(true),
		nano.WithClientAddr(":8080"),
		nano.WithWSPath("/nano"),
		nano.WithComponents(serve.Services),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
		nano.WithDebugMode(),
	)
}
