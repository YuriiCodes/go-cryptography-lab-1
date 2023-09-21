package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"crypto-lab-1/task1"
	"crypto-lab-1/task2"
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
	eulerPhiResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	eulerPhiForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "EulerPhi:", Widget: eulerPhiInput},
			{Text: "Result", Widget: eulerPhiResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
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
			eulerPhiResultLabel.SetText(result)
		},
	}

	mobiusInput := widget.NewEntry()
	mobiusResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	mobiusForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Mobius:", Widget: mobiusInput},
			{Text: "Result", Widget: mobiusResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
			// Retrieve the text from the Entry widget and update the result label.
			parserMobiusInput, err := strconv.Atoi(mobiusInput.Text)
			if err != nil {
				// write 'invalid input' to the result label:
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", err)), myWindow)
			}
			fmt.Println("Mobius input: ", parserMobiusInput)
			fmt.Println("Mobius result: ", task1.Mobius(parserMobiusInput))

			result := "The result is: " + strconv.Itoa(task1.Mobius(parserMobiusInput))
			mobiusResultLabel.SetText(result)
		},
	}

	LCMInput := widget.NewEntry()
	LCMResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	LCMForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "LCM:", Widget: LCMInput},
			{Text: "Result", Widget: LCMResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
			// Parse an array of integers from the entry widget - they will be separated by commas.
			// and then pass them to task1.FindLCM()
			if LCMInput.Text == "" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			inputs, err := task1.ParseIntArray(LCMInput.Text)
			if len(inputs) == 0 {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", err)), myWindow)
				return
			}
			if err != nil {
				// write 'invalid input' to the result label:
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid array of integer input: \n%s", err)), myWindow)
				return
			}
			LCMResultLabel.SetText("The result is: " + strconv.Itoa(task1.FindLCM(inputs)))
			fmt.Println("LCM input: ", inputs)
			fmt.Println("LCM result: ", task1.FindLCM(inputs))
		},
	}

	Eq1xInput := widget.NewEntry()
	Eq1mInput := widget.NewEntry()
	Eq2xInput := widget.NewEntry()
	Eq2mInput := widget.NewEntry()
	Eq3xInput := widget.NewEntry()
	Eq3mInput := widget.NewEntry()
	EqResultLabel := widget.NewLabel("") // Create an empty label to display the result.

	EqForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "eq1x:", Widget: Eq1xInput},
			{Text: "eq1m:", Widget: Eq1mInput},
			{Text: "eq2x:", Widget: Eq2xInput},
			{Text: "eq2m:", Widget: Eq2mInput},
			{Text: "eq3x:", Widget: Eq3xInput},
			{Text: "eq3m:", Widget: Eq3mInput},
			{Text: "Result", Widget: EqResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
			// parse each entry to big.Int and then pass them to task2.SolveSystem()
			// and then display the result in a label

			eq1x := new(big.Int)

			// parse eq1x
			eq1x, ok := eq1x.SetString(Eq1xInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse eq1m
			eq1m, ok := new(big.Int).SetString(Eq1mInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse eq2x
			eq2x, ok := new(big.Int).SetString(Eq2xInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse eq2m
			eq2m, ok := new(big.Int).SetString(Eq2mInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse eq3x
			eq3x, ok := new(big.Int).SetString(Eq3xInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse eq3m
			eq3m, ok := new(big.Int).SetString(Eq3mInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// solve the system
			congruences := [][2]*big.Int{
				{eq1x, eq1m},
				{eq2x, eq2m},
				{eq3x, eq3m},
			}
			x := task2.SolveCRT(congruences)
			EqResultLabel.SetText("The result is: " + x.String())
		},
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Task 1",
			container.NewVBox(
				eulerPhiForm,
				widget.NewSeparator(),
				mobiusForm,
				widget.NewSeparator(),
				LCMForm,
			),
		),
		container.NewTabItem("Task 2",
			container.NewVBox(
				EqForm,
			),
		),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
