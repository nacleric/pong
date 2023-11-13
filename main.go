package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	keys []ebiten.Key
	ball Ball
}

type Ball struct {
	posX   float32
	posY   float32
	vX     float32
	vY     float32
	radius float32
	color  color.Color
}

func (b *Ball) move() {
	// Move ball right
	if b.vX > 0 {
		b.posX += b.vX
	}

	// Move ball left
	if b.vX < 0 {
		b.posX += b.vX
	}

	// Move ball up
	if b.vY < 0 {
		b.posY += b.vY
	}

	// Move ball down
	if b.vY > 0 {
		b.posY += b.vY
	}
	
	// side wall collission
	if b.posX >= screenWidth-b.radius || b.posX <= 0+b.radius {
		b.vX *= -1
	}

	if b.calculateWallCollision() {
		b.vY *= -1
	}
}

func (b *Ball) calculateWallCollision() bool {
	// Top Wall
	if b.posY <= 0 && b.posX >= 0 && b.posX <= screenWidth {
		return true
	}
	// Bottom Wall
	if b.posY >= screenHeight && b.posX >= 0 && b.posX <= screenWidth {
		return true
	}
	return false
}

type Player struct {
	posX  float64
	posY  float64
	color color.Color
}

func drawBall(screen *ebiten.Image, b Ball) {
	vector.DrawFilledCircle(screen, b.posX, b.posY, b.radius, b.color, true)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.ball.move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Pong")

	// Ball
	drawBall(screen, g.ball)
}

var count int

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Pong")
	ball := Ball{posX: screenWidth / 2, posY: screenHeight / 2, vX: 2, vY: 2, radius: 10, color: color.White}
	if err := ebiten.RunGame(&Game{ball: ball}); err != nil {
		log.Fatal(err)
	}
}
