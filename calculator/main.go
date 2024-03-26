package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// 使用fyne开发一个数学计算器(+-*/)

func main() {
	// 1. 编写一个计算器整体的窗体框架
	myApp := app.New()
	myWindow := myApp.NewWindow("计算器")
	display := widget.NewEntry()

	// 2. 定义一些初始值
	display.Text = "0"
	var firstNumber float64
	var operation string

	// 3. 定义数字按钮点击时所执行的函数
	appendNumber := func(num string) {
		display.Text = num
		display.Refresh()
	}

	// 4. 定义运算按钮的回调函数
	doOperation := func(op string) {
		firstNumber, _ = strconv.ParseFloat(display.Text, 64)
		operation = op
	}

	// 5. 定义计算结果时的计算函数
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
			result = firstNumber / secondNumber
			if secondNumber == 0 {
				display.Text = "ERROR"
				return
			}
		}
		display.Text = fmt.Sprintf("%f", result)
		display.Refresh()
	}

	// 6. 第一个清理按钮的回调函数
	clearButton := func() {
		display.Text = "0"
		firstNumber = 0
		operation = ""
		display.Refresh()
	}

	// 定义计算器的所有按钮
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
	buttonClear := widget.NewButton("C", clearButton)

	// 整齐排列按钮
	buttons := container.NewGridWithColumns(4,
		button7, button8, button9, buttonDivide,
		button4, button5, button6, buttonMultiply,
		button1, button2, button3, buttonMinus,
		button0, buttonClear, buttonEquals, buttonPlus,
	)

	myWindow.SetContent(container.NewVBox(
		display,
		// 添加这些按钮到计算器中
		buttons,
	))
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
