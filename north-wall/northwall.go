package north_wall

import (
	"github.com/snippetor/bingo/node"
	"github.com/snippetor/bingo/rpc"
	"github.com/snippetor/bingo-example/log"
)

type NorthWall struct {
	node.Model
}

func (w *NorthWall) OnInit() {
	w.RegisterRPCMethod("OnEcho", w.OnEcho)
}

func (w *NorthWall) OnEcho(ctx *rpc.Context) *rpc.Result {
	log.Logger.D("NorthWall receive echo message...")
	log.Logger.D("%s", ctx.Args)
	w.CallRPCMethodNoReturn("southwall", "OnEcho", ctx.Args)
	return nil
}
