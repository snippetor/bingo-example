package sourth_wall

import (
	"github.com/snippetor/bingo/node"
	"github.com/snippetor/bingo/rpc"
	"github.com/snippetor/bingo-example/log"
)

type SouthWall struct {
	node.Model
}

func (w *SouthWall) OnInit() {
	w.RegisterRPCMethod("OnEcho", w.OnEcho)
}

func (w *SouthWall) OnEcho(ctx *rpc.Context) *rpc.Result {
	log.Logger.D("SouthWall receive echo message...")
	log.Logger.D("%s", ctx.Args)
	w.CallRPCMethodNoReturn("westwall", "OnEcho", ctx.Args)
	return nil
}
