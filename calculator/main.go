package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	display := widget.NewEntry()
	display.Text = "0"
	// display.ReadOnly = true
	// display.Alignment = widget.AlignmentTrailing

	var firstNumber float64
	var operation string

	// Function to handle number button click
	appendNumber := func(num string) {
		if display.Text == "0" {
			display.Text = num
		} else {
			display.Text += num
		}
	}

	// Function to handle operation button click
	doOperation := func(op string) {
		firstNumber, _ = strconv.ParseFloat(display.Text, 64)
		operation = op
		display.Text = "0"
		display.Refresh()
	}

	// Function to handle equal button click
	calculate := func() {
		secondNumber, _ := strconv.ParseFloat(display.Text, 64)
		var result float64
		switch operation {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber == 0 {
				display.Text = "Error"
				return
			}
			result = firstNumber / secondNumber
		}
		display.Text = fmt.Sprintf("%f", result)
		display.Refresh()
		fmt.Println(result)
	}

	// Function to handle clear button click
	clear := func() {
		display.Text = "0"
		firstNumber = 0
		operation = ""
		display.Refresh()
	}

	// Buttons
	button0 := widget.NewButton("0", func() { appendNumber("0") })
	button1 := widget.NewButton("1", func() { appendNumber("1") })
	button2 := widget.NewButton("2", func() { appendNumber("2") })
	button3 := widget.NewButton("3", func() { appendNumber("3") })
	button4 := widget.NewButton("4", func() { appendNumber("4") })
	button5 := widget.NewButton("5", func() { appendNumber("5") })
	button6 := widget.NewButton("6", func() { appendNumber("6") })
	button7 := widget.NewButton("7", func() { appendNumber("7") })
	button8 := widget.NewButton("8", func() { appendNumber("8") })
	button9 := widget.NewButton("9", func() { appendNumber("9") })
	buttonPlus := widget.NewButton("+", func() { doOperation("+") })
	buttonMinus := widget.NewButton("-", func() { doOperation("-") })
	buttonMultiply := widget.NewButton("*", func() { doOperation("*") })
	buttonDivide := widget.NewButton("/", func() { doOperation("/") })
	buttonEquals := widget.NewButton("=", calculate)
	buttonClear := widget.NewButton("C", clear)

	buttons := container.NewGridWithColumns(4,
		button7, button8, button9, buttonDivide,
		button4, button5, button6, buttonMultiply,
		button1, button2, button3, buttonMinus,
		button0, buttonClear, buttonEquals, buttonPlus,
	)

	myWindow.SetContent(container.NewVBox(
		display,
		buttons,
	))

	myWindow.ShowAndRun()
}
