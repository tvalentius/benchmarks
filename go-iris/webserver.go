// go get -v -u github.com/kataras/iris
// go build webserver.go
// ./webserver 1337

package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/iris/context"
  "os"
)

func main() {
  app := iris.Default()

  app.Get("/", func(ctx context.Context) {
    ctx.WriteString("Halo Bali!")
  })

  app.Run(iris.Addr(":"+string(os.Args[1])))
}
