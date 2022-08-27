package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Struct representing paddle object
// x and y represent the center point of the paddle
type Paddle struct {
	Width  int32
	Height int32
	X      int32 // center of the paddle x
	Y      int32 // center of the paddle y
	color  rl.Color
	speed  int32
}

type Point struct {
	X int32
	Y int32
}

func (paddle *Paddle) Draw() {
	topLeft := paddle.GetTopLeft()
	rl.DrawRectangle(topLeft.X, topLeft.Y, paddle.Width, paddle.Height, paddle.color)
}

// x and y represent the center of paddle

func (paddle *Paddle) SetValues(x, y, width, height int32, speed int32, color rl.Color) {
	paddle.X = x
	paddle.Y = y
	paddle.Width = width
	paddle.Height = height
	paddle.speed = speed
	paddle.color = color
}

func (paddle *Paddle) GetTopLeft() Point {
	return Point{
		X: paddle.X - (paddle.Width / 2),
		Y: paddle.Y - (paddle.Height / 2),
	}
}

func (paddle *Paddle) MoveUp() {
	paddle.Y -= int32(float32(paddle.speed) * rl.GetFrameTime())
	paddle.checkUpdateBounds()
}

func (paddle *Paddle) MoveDown() {
	paddle.Y += int32(float32(paddle.speed) * rl.GetFrameTime())
	paddle.checkUpdateBounds()
}

func (paddle *Paddle) checkUpdateBounds() {
	if paddle.Y-(paddle.Height/2) < 0 {
		paddle.Y = paddle.Height / 2
	}
	if paddle.Y+(paddle.Height/2) > int32(rl.GetScreenHeight()) {
		paddle.Y = int32(rl.GetScreenHeight()) - (paddle.Height / 2)
	}
}

func (paddle *Paddle) GetRect() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(paddle.X - (paddle.Width / 2)),
		Y:      float32(paddle.Y - (paddle.Height / 2)),
		Width:  float32(paddle.Width),
		Height: float32(paddle.Height),
	}
}

func (paddle *Paddle) Reset() {
	// paddle.SetValues((rightPaddleX), rightPaddleY, paddleWidth, paddleHeight, paddleSpeed, rl.RayWhite)
}
