package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	keys []ebiten.Key
	p1   Player
	p2   Player
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

func (b *Ball) move(p1 Player, p2 Player) {
	if b.posX == screenWidth/2 && b.posY == screenHeight/2 {
		b.vX = 1
		b.vY = 0
	}

	if b.vX >= 1 && b.vY == 0 {
		b.posX += b.vX
	}
	
	// Resets ball starting position
	if b.posX >= screenWidth-b.radius || b.posX <= 0 + b.radius {
		b.posX = screenWidth/2
		b.posY = screenHeight/2
	}

	b.calculateCollision(p1, p2)
}

func (b *Ball) calculateCollision(p1 Player, p2 Player) {
	// Math is wrong here
	hitboxYStart := p2.posY 
	hitboxYEnd := p2.posY+playerHeight 
	hitboxX := p2.posX
	if b.posY > float32(hitboxYStart) && b.posY < float32(hitboxYEnd)  && b.posX > float32(hitboxX) {
		fmt.Println("hit")
	}
}

type Player struct {
	posX  float64
	posY  float64
	color color.Color
}

func (p *Player) up() {
	if p.posY > 0 && p.posY <= screenHeight-playerHeight {
		p.posY -= 1
	}
}

func (p *Player) down() {
	if p.posY >= 0 && p.posY < screenHeight-playerHeight {
		p.posY += 1
	}
}

func drawPlayer(screen *ebiten.Image, p Player) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(playerScaleWidth, playerScaleHeight)
	op.GeoM.Translate(p.posX, p.posY)

	playerImage := ebiten.NewImage(imageWidth, imageHeight)
	playerImage.Fill(p.color)
	screen.DrawImage(playerImage, op)
}

func drawBall(screen *ebiten.Image, b Ball) {
	vector.DrawFilledCircle(screen, b.posX, b.posY, b.radius, b.color, true)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.ball.move(g.p1, g.p2)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Pong")

	// Player1
	drawPlayer(screen, g.p1)

	// Player2
	drawPlayer(screen, g.p2)

	// Ball
	drawBall(screen, g.ball)

	for _, keyPress := range g.keys {
		switch keyPress {
		case ebiten.KeyW:
			g.p1.up()
		case ebiten.KeyS:
			g.p1.down()
		case ebiten.KeyUp:
			g.p2.up()
		case ebiten.KeyDown:
			g.p2.down()
		}

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Pong")
	p1 := Player{0, 0, color.White}
	p2 := Player{screenWidth-playerWidth, 0, color.White}
	ball := Ball{posX: screenWidth / 2, posY: screenHeight / 2, radius: 3, color: color.White}
	if err := ebiten.RunGame(&Game{p1: p1, p2: p2, ball: ball}); err != nil {
		log.Fatal(err)
	}
}
