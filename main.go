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

	// Resets ball starting position
	if b.posX >= screenWidth-b.radius || b.posX <= 0+b.radius {
		b.posX = screenWidth / 2
		b.posY = screenHeight / 2
	}

	if b.calculatePlayerCollision(p1, p2) {
		// Reverses direction if ball hits something
		b.vX *= -1
	}

	if b.calculateWallCollision() {
		fmt.Println("wall hit")
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

func (b *Ball) calculatePlayerCollision(p1 Player, p2 Player) bool {
	p1HitboxYStart := p1.posY
	p1hitboxYEnd := p1.posY + playerHeight
	p1hitboxX := p1.posX + playerWidth + float64(b.radius)

	p2HitboxYStart := p2.posY
	p2hitboxYEnd := p2.posY + playerHeight
	p2hitboxX := p2.posX - float64(b.radius)
	if b.posY > float32(p2HitboxYStart) && b.posY < float32(p2hitboxYEnd) && b.posX > float32(p2hitboxX) {
		return true
	}

	if b.posY > float32(p1HitboxYStart) && b.posY < float32(p1hitboxYEnd) && b.posX < float32(p1hitboxX) {
		return true
	}

	return false
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
	p2 := Player{screenWidth - playerWidth, 0, color.White}
	ball := Ball{posX: screenWidth / 2, posY: screenHeight / 2, vX: 2, vY: 2, radius: 3, color: color.White}
	if err := ebiten.RunGame(&Game{p1: p1, p2: p2, ball: ball}); err != nil {
		log.Fatal(err)
	}
}
