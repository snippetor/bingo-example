package west_wall

import (
	"github.com/snippetor/bingo/node"
	"github.com/snippetor/bingo/rpc"
	"github.com/snippetor/bingo-example/log"
)

type WestWall struct {
	node.Model
}

func (w *WestWall) OnInit() {
	w.RegisterRPCMethod("OnEcho", w.OnEcho)
}

func (w *WestWall) OnEcho(ctx *rpc.Context) *rpc.Result {
	log.Logger.D("WestWall receive echo message...")
	log.Logger.D("%s", ctx.Args)
	return nil
}