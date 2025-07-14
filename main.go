package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func CreateBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func CreateBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}

func main() {
	fmt.Println("Start")
	color.Blue("Start")
}
