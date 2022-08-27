package entity

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Radius float32
	X      int32
	Y      int32
	Color  rl.Color
	SpeedX int32
	SpeedY int32
}

func (ball *Ball) Draw() {
	rl.DrawCircle(ball.X, ball.Y, ball.Radius, ball.Color)
}

func (ball *Ball) SetPosition(x int32, y int32) {
	ball.X = x
	ball.Y = y
}

func (ball *Ball) SetRadius(radius float32) {
	ball.Radius = radius
}

func (ball *Ball) SetColor(color rl.Color) {
	ball.Color = color
}

func (ball *Ball) PrintValues() {
	fmt.Println(ball)
}

func (ball *Ball) CheckUpdateBounds() {
	if ball.Y > int32(rl.GetScreenHeight()) {
		ball.Y = int32(rl.GetScreenHeight())
		ball.SpeedY *= -1
	}

	if ball.Y <= 0 {
		ball.Y = 0
		ball.SpeedY *= -1
	}
}

func (ball *Ball) Reset() {
	ball.X, ball.Y = int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2)
	ball.SpeedX, ball.SpeedY = 150, 150
}
