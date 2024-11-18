package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//Make the window anda  resize
	myApp := app.New()
	myWindow := myApp.NewWindow("Triangulo de Pitagoras")
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.SetIcon(theme.FyneLogo())
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	tr := Triangulo{0, 0, 0}
	//Add to the border pane all borders
	//Top - NAME
	title := canvas.NewText("Triangulo de Pitagoras", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	top := container.NewCenter(title)
	//Right Image
	img := canvas.NewImageFromFile("pitagoras.jpg")
	lblType := widget.NewLabel("Tipo Desconocido")
	img.SetMinSize(fyne.NewSize(250, 400))
	lblType.TextStyle = fyne.TextStyle{Bold: true}
	lblType.Alignment = fyne.TextAlignCenter

	right := container.NewVBox(img, lblType)
	// Left - gtidpane to labels
	lado1 := canvas.NewText("Lado1: ", color.White)
	lado2 := canvas.NewText("Lado 2: ", color.White)
	lado3 := canvas.NewText("Lado 3: ", color.White)
	lado1.TextStyle = fyne.TextStyle{Bold: true}
	lado2.TextStyle = fyne.TextStyle{Bold: true}
	lado3.TextStyle = fyne.TextStyle{Bold: true}
	left := container.NewGridWithRows(3, lado1, lado2, lado3)
	// Mid - gridPane to Entrys
	date1 := widget.NewEntry()
	date1.SetPlaceHolder("Lado 1 :")
	date2 := widget.NewEntry()
	date2.SetPlaceHolder("Lado 2 :")
	date3 := widget.NewEntry()
	date3.SetPlaceHolder("Lado 3 :")
	middle := container.NewGridWithRows(3, date1, date2, date3)
	//Bottom - grind pane to buttons
	actualizar := widget.NewButton("Actualizar Medidas", func() {
		val1, _ := strconv.ParseFloat(date1.Text, 32)
		val2, _ := strconv.ParseFloat(date2.Text, 32)
		val3, _ := strconv.ParseFloat(date3.Text, 32)
		tr.SetLado1(val1)
		tr.SetLado2(val2)
		tr.SetLado3(val3)

	})
	calcular := widget.NewButton("Calcular", func() {
		lblType.SetText("El triangulo es: " + tr.GetTipo())
	})
	help := widget.NewButton("Help", func() {
		dialog.ShowInformation("HElP", "Calcula Los Lados", myWindow)
	})
	bottom := container.NewGridWithColumns(3, actualizar, calcular, help)
	//Set content
	content := container.NewBorder(top, bottom, left, right, middle)

	//Display the window
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
