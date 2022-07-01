package main

import "fmt"

type Pen interface {
	Write()
}

func Write(p Pen) {
	p.Write()
}

type NewPen struct{}

func (n *NewPen) Write() {
	fmt.Println("New Pen")
}

type OldPen struct{}

func (o *OldPen) OldWrite() {
	fmt.Println("Old Pen")
}

type PenAdapter struct {
	*OldPen
}

func (p *PenAdapter) Write() {
	p.OldWrite()
}

func main() {
	Write(&NewPen{}) // New Pen
	// Write(&OldPen{}) // Compile error
	Write(&PenAdapter{}) // Old Pen
}
