package main

import (
	entity "github.com/adithyavhebbar/pong/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 800
	windowHeight = 600
	title        = "Pong"
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, title)

	rl.SetTargetFPS(60)

	ballX, ballY := int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2)

	ball := entity.Ball{
		X:      ballX,
		Y:      ballY,
		Radius: 7,
		SpeedX: 150,
		SpeedY: 150,
		Color:  rl.RayWhite,
	}
	paddleWidth := int32(10)
	paddleHeight := int32(100)
	paddleSpeed := int32(500)

	leftPaddleX := 30 + (paddleWidth / 2)
	leftPaddleY := int32(rl.GetScreenHeight() / 2)
	leftPaddle := entity.Paddle{}
	leftPaddle.SetValues(leftPaddleX, leftPaddleY, paddleWidth, paddleHeight, paddleSpeed, rl.RayWhite)

	rightPaddleX := rl.GetScreenWidth() - 30 - int(paddleWidth/2)
	rightPaddleY := int32(rl.GetScreenHeight() / 2)

	rightPaddle := entity.Paddle{}

	rightPaddle.SetValues(int32(rightPaddleX), rightPaddleY, paddleWidth, paddleHeight, paddleSpeed, rl.RayWhite)

	winnerText := ""

	for !rl.WindowShouldClose() {
		ball.X += int32(float32(ball.SpeedX) * rl.GetFrameTime())
		ball.Y += int32(float32(ball.SpeedY) * rl.GetFrameTime())
		ball.CheckUpdateBounds()

		rl.BeginDrawing()

		rl.DrawFPS(10, 10)

		rl.ClearBackground(rl.Black)

		// HandleInteraction(leftPaddle, rightPaddle)
		if rl.IsKeyDown(rl.KeyUp) {
			rightPaddle.MoveUp()
		}

		if rl.IsKeyDown(rl.KeyDown) {
			rightPaddle.MoveDown()
		}

		if rl.IsKeyDown(rl.KeyW) {
			leftPaddle.MoveUp()
		}

		if rl.IsKeyDown(rl.KeyS) {
			leftPaddle.MoveDown()
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			ball.Reset()
			winnerText = ""
		}

		if HandleCollision(ball, leftPaddle) {
			if ball.SpeedX < 0 {
				ball.SpeedX = int32(float32(ball.SpeedX) * -1.1)
				// ball.SpeedY = (ball.Y - leftPaddle.Y) / (leftPaddle.Height / 2) * ball.SpeedX
			}
		}

		if HandleCollision(ball, rightPaddle) {
			if ball.SpeedX > 0 {
				ball.SpeedX = int32(float32(ball.SpeedX) * -1.1)
				// ball.SpeedY = (ball.Y - rightPaddle.Y) / (rightPaddle.Height / 2) * -ball.SpeedX
			}
		}

		if ball.X > int32(rl.GetScreenWidth()) {
			fontSize := 30
			winnerText = "Left Player Won"
			size := rl.MeasureText(winnerText, int32(fontSize))
			rl.DrawText(winnerText, int32(rl.GetScreenWidth()/2-int(size)/2), int32(rl.GetScreenHeight()/2), int32(fontSize), rl.Yellow)
		}

		if ball.X < 0 {
			fontSize := 30
			winnerText = "Right Player Won"
			size := rl.MeasureText(winnerText, int32(fontSize))
			rl.DrawText(winnerText, int32(rl.GetScreenWidth()/2-int(size)/2), int32(rl.GetScreenHeight()/2), int32(fontSize), rl.Yellow)
		}

		ball.Draw()
		leftPaddle.Draw()
		rightPaddle.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func HandleCollision(ball entity.Ball, paddle entity.Paddle) bool {
	return rl.CheckCollisionCircleRec(rl.Vector2{X: float32(ball.X), Y: float32(ball.Y)}, ball.Radius, paddle.GetRect())
}
