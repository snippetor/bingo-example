package east_wall

import (
	"github.com/snippetor/bingo/node"
	"github.com/snippetor/bingo/timer"
	"time"
	"github.com/snippetor/bingo/rpc"
	"math"
	"github.com/snippetor/bingo-example/log"
)

type EastWall struct {
	node.Model
}

func (w *EastWall) OnInit() {
	timer.NewTimer(5*time.Second, func(v ...interface{}) {
		log.Logger.D("echo start...")
		args := rpc.Args{}
		args.Put("arg1", "arg1")
		args.PutBool("arg2", true)
		args.PutFloat32("arg3", 1.33333)
		args.PutInt("arg4", 1111)
		args.PutInt64("arg5", math.MaxInt64)
		w.CallRPCMethodNoReturn("northwall", "OnEcho", args)
	}, nil).Run()
}
