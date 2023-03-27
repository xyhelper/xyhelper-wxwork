package main

import (
	_ "xyhelper-wxwork/internal/packed"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

	_ "xyhelper-wxwork/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"xyhelper-wxwork/internal/cmd"
)

func main() {
	// gres.Dump()
	cmd.Main.Run(gctx.New())
}
