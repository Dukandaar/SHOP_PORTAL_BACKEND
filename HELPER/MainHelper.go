package helper

import (
	"runtime"

	"github.com/kataras/iris/v12"
)

func Onit() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores
}

func ServerUp(ctx iris.Context) {
	ctx.HTML("Backend Server Is UP")
}
