package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"math/rand"
	"time"
)

type Board [4][4]uint16

func NewBoard() *Board {
	return &Board{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
}

func StringMul(s string, n int) string {
	r := ""
	for n > 0 {
		r += s
		n--
	}
	return r
}

func (B *Board) MoveL() {
	happened := true
	for happened {
		happened = false
		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				if x != 0 {
					if B[x-1][y] == 0 && B[x][y] != 0 {
						B[x-1][y] = B[x][y]
						B[x][y] = 0
						happened = true
					} else if B[x-1][y] == B[x][y] && B[x-1][y] != 0 {
						B[x-1][y] *= 2
						B[x][y] = 0
						happened = true
					}
				}
			}
		}
	}
}

func (B *Board) MoveR() {
	happened := true
	for happened {
		happened = false
		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				if x != 3 {
					if B[x+1][y] == 0 && B[x][y] != 0 {
						B[x+1][y] = B[x][y]
						B[x][y] = 0
						happened = true
					} else if B[x+1][y] == B[x][y] && B[x+1][y] != 0 {
						B[x+1][y] *= 2
						B[x][y] = 0
						happened = true
					}
				}
			}
		}
	}
}
func (B *Board) MoveU() {
	happened := true
	for happened {
		happened = false
		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				if y != 0 {
					if B[x][y-1] == 0 && B[x][y] != 0 {
						B[x][y-1] = B[x][y]
						B[x][y] = 0
						happened = true
					} else if B[x][y-1] == B[x][y] && B[x][y-1] != 0 {
						B[x][y-1] *= 2
						B[x][y] = 0
						happened = true
					}
				}
			}
		}
	}
}
func (B *Board) MoveD() {
	happened := true
	for happened {
		happened = false
		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				if y != 3 {
					if B[x][y+1] == 0 && B[x][y] != 0 {
						B[x][y+1] = B[x][y]
						B[x][y] = 0
						happened = true
					} else if B[x][y+1] == B[x][y] && B[x][y+1] != 0 {
						B[x][y+1] *= 2
						B[x][y] = 0
						happened = true
					}
				}
			}
		}
	}
}

func (B *Board) Render() {

	renderElements := []ui.Drawable{}

	// calculate, where the number has to places

	x2 := 0
	y2 := 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			pg := widgets.NewParagraph()
			ns := fmt.Sprint(B[x][y])
			if ns == "2" {
				pg.Text = "\n    2"
			} else if ns == "4" {
				pg.Text = "\n    4"
			} else if ns == "8" {
				pg.Text = "\n    8"
			} else if ns == "16" {
				pg.Text = "\n    16"
			} else if ns == "32" {
				pg.Text = "\n    32"
			} else if ns == "64" {
				pg.Text = "\n    64"
			} else if ns == "128" {
				pg.Text = "\n   128"
			} else if ns == "256" {
				pg.Text = "\n   256"
			} else if ns == "512" {
				pg.Text = "\n   512"
			} else if ns == "1024" {
				pg.Text = "\n  1024"
			} else if ns == "2048" {
				pg.Text = "\n  2048"
			}
			pg.TextStyle = ui.NewStyle(ui.ColorMagenta)
			pg.SetRect(x2, y2, x2+15, y2+5)
			x2 += 15
			renderElements = append(renderElements, pg)
		}
		y2 += 5
		x2 = 0
	}

	ui.Render(renderElements...)
}

func Randomizer() {
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Nanosecond * 1)
	}
}

func (B *Board) Spawn() *Board {
	type Point struct {
		x int
		y int
	}
	freeFields := []Point{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if B[x][y] == 0 {
				freeFields = append(freeFields, Point{x, y})
			}
		}
	}
	if len(freeFields) == 0 {
		func() {
			for {
				r := widgets.NewParagraph()
				r.SetRect(0, 0, 20, 5)
				ui.Render(r)
				renderBoard = false
			}
		}()
	}
	fp := freeFields[rand.Intn(len(freeFields))]
	cx := fp.x
	cy := fp.y
	B[cx][cy] = []uint16{2, 4}[rand.Intn(2)] // spawn 2 or 4
	return B
}

func (B *Board) SpawnSpec(x uint16) *Board {
	type Point struct {
		x int
		y int
	}
	freeFields := []Point{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if B[x][y] == 0 {
				freeFields = append(freeFields, Point{x, y})
			}
		}
	}
	if len(freeFields) == 0 {
		renderBoard = false
	}
	fp := freeFields[rand.Intn(len(freeFields))]
	cx := fp.x
	cy := fp.y
	B[cx][cy] = []uint16{x, x}[rand.Intn(2)] // spawn 2 or 4
	return B
}

var renderBoard = true

func main() {
	go Randomizer()
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	b := NewBoard()
	b.SpawnSpec(2).SpawnSpec(2)
	go func() {
		for {
			if renderBoard {
				b.Render()
			} else {
				for y := 0; y < 4; y++ {
					for x := 0; x < 4; x++ {
						if b[x][y] == 2048 {
							r := widgets.NewParagraph()
							r.Text = "You won."
							r.SetRect(0, 0, 20, 5)
							ui.Render(r)
						}
					}
				}
				r := widgets.NewParagraph()
				r.Text = "You lost."
				r.SetRect(0, 0, 20, 5)
				ui.Render(r)
			}
		}
	}()

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			moved := true
			if e.ID == "<Escape>" {
				return
			} else if e.ID == "<Down>" {
				// move down
				b.MoveD()
			} else if e.ID == "<Up>" {
				// move up
				b.MoveU()
			} else if e.ID == "<Left>" {
				// move left
				b.MoveL()
			} else if e.ID == "<Right>" {
				// move right
				b.MoveR()
			} else {
				moved = false
			}
			if moved {
				b.Spawn()
			}
		}
	}
}
