package main

import (
	"errors"
	"fmt"
	"strconv"

	"crypto-lab-1/task1"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")
	myWindow.Resize(fyne.NewSize(800, 600))

	eulerPhiInput := widget.NewEntry()
	resultLabel := widget.NewLabel("") // Create an empty label to display the result.

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Enter N for EulerPhi", Widget: eulerPhiInput},
			{Text: "Result", Widget: resultLabel},
			{Widget: widget.NewButton("Calculate", func() {
				// Retrieve the text from the Entry widget and update the result label.
				parserEulerPhiInput, err := strconv.Atoi(eulerPhiInput.Text)
				if err != nil {
					// write 'invalid input' to the result label:
					dialog.ShowError(
						errors.New(fmt.Sprintf("Invalid integer input: \n%s", err)), myWindow)
				}
				fmt.Println("EulerPhi input: ", parserEulerPhiInput)
				fmt.Println("EulerPhi result: ", task1.EulerPhi(parserEulerPhiInput))

				result := "The result is: " + strconv.Itoa(task1.EulerPhi(parserEulerPhiInput))
				resultLabel.SetText(result)
			})},
		},
		SubmitText: "Calculate",
		OnSubmit:   nil, // You can add a callback here if needed.
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Task 1",
			container.NewVBox(
				form,
			),
		),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
