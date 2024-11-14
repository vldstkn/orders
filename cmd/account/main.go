package main

import "orders/internal/app"

func main() {
	accountApp := app.NewAccountApp()
	accountApp.Run()

}
