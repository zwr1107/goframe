package main

import (
	_ "goframe/internal/packed"

	_ "goframe/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
