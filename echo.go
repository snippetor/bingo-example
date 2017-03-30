package main

import (
	"github.com/snippetor/bingo"
	"github.com/snippetor/bingo-example/east-wall"
	"github.com/snippetor/bingo-example/north-wall"
	"github.com/snippetor/bingo-example/south-wall"
	"github.com/snippetor/bingo-example/west-wall"
)

func main() {
	bingo.BindNodeModel("eastwall", &east_wall.EastWall{})
	bingo.BindNodeModel("northwall", &north_wall.NorthWall{})
	bingo.BindNodeModel("southwall", &sourth_wall.SouthWall{})
	bingo.BindNodeModel("westwall", &west_wall.WestWall{})
	bingo.Run()
}
