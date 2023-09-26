package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"crypto-lab-1/task1"
	"crypto-lab-1/task2"
	"crypto-lab-1/task3"
	"crypto-lab-1/task4"
	"crypto-lab-1/task5"
	"crypto-lab-1/task6"
	"crypto-lab-1/task7"
	"crypto-lab-1/task8"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Yurii's Lab")
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

	LegendreAInput := widget.NewEntry()
	LegendrePInput := widget.NewEntry()
	LegendreResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	LegendreForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "a:", Widget: LegendreAInput},
			{Text: "p:", Widget: LegendrePInput},
			{Text: "Result", Widget: LegendreResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
			// parse each entry to big.Int and then pass them to task3.Legendre()
			// and then display the result in a label

			a := new(big.Int)

			// parse a
			a, ok := a.SetString(LegendreAInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse p
			p, ok := new(big.Int).SetString(LegendrePInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// calculate the result
			result := task3.Legendre(a, p)
			LegendreResultLabel.SetText("The result is: " + result.String())
		},
	}

	JacobiInputA := widget.NewEntry()
	JacobiInputN := widget.NewEntry()
	JacobiResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	JacobiForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "a:", Widget: JacobiInputA},
			{Text: "n:", Widget: JacobiInputN},
			{Text: "Result", Widget: JacobiResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
			// parse each entry to big.Int and then pass them to task3.Jacobi()
			// and then display the result in a label

			a := new(big.Int)

			// parse a
			a, ok := a.SetString(JacobiInputA.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse n
			n, ok := new(big.Int).SetString(JacobiInputN.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// calculate the result
			result := task3.JacobiSymbol(a, n)

			stringResult := strconv.Itoa(result)

			JacobiResultLabel.SetText(stringResult)
		},
	}

	PollardRhoInputN := widget.NewEntry()
	PollardRhoResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	PollardRhoForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "n:", Widget: PollardRhoInputN},
			{Text: "Result", Widget: PollardRhoResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {

			if PollardRhoInputN.Text == "" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}
			if PollardRhoInputN.Text == "0" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Zero input")), myWindow)
				return
			}

			// parse each entry to big.Int and then pass them to task3.Jacobi()
			// and then display the result in a label

			n := new(big.Int)

			// parse n
			n, ok := n.SetString(PollardRhoInputN.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// calculate the result
			result := task4.PollardRho(n)

			PollardRhoResultLabel.SetText(result.String())
		},
	}

	BabyStepGiantStepBaseInput := widget.NewEntry()
	BabyStepGiantStepTargetInput := widget.NewEntry()
	BabyStepGiantStepModInput := widget.NewEntry()
	BabyStepGiantStepResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	BabyStepGiantStepForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Base:", Widget: BabyStepGiantStepBaseInput},
			{Text: "Target:", Widget: BabyStepGiantStepTargetInput},
			{Text: "Mod:", Widget: BabyStepGiantStepModInput},
			{Text: "Result", Widget: BabyStepGiantStepResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {

			if BabyStepGiantStepBaseInput.Text == "" || BabyStepGiantStepTargetInput.Text == "" || BabyStepGiantStepModInput.Text == "" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}
			if BabyStepGiantStepBaseInput.Text == "0" || BabyStepGiantStepTargetInput.Text == "0" || BabyStepGiantStepModInput.Text == "0" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Zero input")), myWindow)
				return
			}

			// parse each entry to big.Int and then pass them to task3.Jacobi()
			// and then display the result in a label

			base := new(big.Int)
			target := new(big.Int)
			mod := new(big.Int)

			// parse base
			base, ok := base.SetString(BabyStepGiantStepBaseInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse target
			target, ok = target.SetString(BabyStepGiantStepTargetInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse mod
			mod, ok = mod.SetString(BabyStepGiantStepModInput.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// calculate the result
			result := task5.BabyStepGiantStep(base, target, mod)

			BabyStepGiantStepResultLabel.SetText(result.String())
		},
	}

	CipolloInputN := widget.NewEntry()
	CipolloInputP := widget.NewEntry()

	CipolloResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	CipolloForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "n:", Widget: CipolloInputN},
			{Text: "p:", Widget: CipolloInputP},
			{Text: "Result", Widget: CipolloResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {

			if CipolloInputN.Text == "" || CipolloInputP.Text == "" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}
			if CipolloInputN.Text == "0" || CipolloInputP.Text == "0" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Zero input")), myWindow)
				return
			}

			// parse each entry to big.Int and then pass them to task3.Jacobi()
			// and then display the result in a label

			n := new(big.Int)
			p := new(big.Int)

			// parse n
			n, ok := n.SetString(CipolloInputN.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse p
			p, ok = p.SetString(CipolloInputP.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// calculate the result
			root1, root2, ok := task6.Cipolla(*n, *p)
			if !ok {
				CipolloResultLabel.SetText("No solution")
				return
			}
			CipolloResultLabel.SetText(fmt.Sprintf("Root1: %s\nRoot2: %s", root1.String(), root2.String()))
		},
	}

	MillerRabinInputN := widget.NewEntry()
	MillerRabinInputK := widget.NewEntry()

	MillerRabinResultLabel := widget.NewLabel("") // Create an empty label to display the result.
	MillerRabinForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "n:", Widget: MillerRabinInputN},
			{Text: "k:", Widget: MillerRabinInputK},
			{Text: "Result", Widget: MillerRabinResultLabel},
		},
		SubmitText: "Calculate",
		OnSubmit: func() {
			if MillerRabinInputN.Text == "" || MillerRabinInputK.Text == "" {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse each entry to big.Int and then pass them to task3.Jacobi()
			// and then display the result in a label

			n := new(big.Int)
			// parse n
			n, ok := n.SetString(MillerRabinInputN.Text, 10)
			if !ok {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// parse k
			k, err := strconv.Atoi(MillerRabinInputK.Text)
			if err != nil {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", "Empty input")), myWindow)
				return
			}

			// calculate the result
			isPrime, err := task7.MillerRabin(n, k)
			if err != nil {
				dialog.ShowError(
					errors.New(fmt.Sprintf("Invalid integer input: \n%s", err.Error())), myWindow)
				return
			}
			if isPrime {
				MillerRabinResultLabel.SetText("Prime")
			} else {
				MillerRabinResultLabel.SetText("Composite")
			}
		},
	}

	// Key Generation Section
	pInput := widget.NewEntry()
	qInput := widget.NewEntry()

	nLabel := widget.NewLabel("n:")
	eLabel := widget.NewLabel("e:")
	dLabel := widget.NewLabel("d:")

	generateKeysBtn := widget.NewButton("Generate Keys", func() {
		p, ok := new(big.Int).SetString(pInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid p"), myWindow)
			return
		}

		q, ok := new(big.Int).SetString(qInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid q"), myWindow)
			return
		}

		n, e, d := task8.GenerateKeys(p, q)
		nLabel.SetText("n: " + n.String())
		eLabel.SetText("e: " + e.String())
		dLabel.SetText("d: " + d.String())
	})

	keyGenForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "p:", Widget: pInput},
			{Text: "q:", Widget: qInput},
			{Widget: container.NewVBox(nLabel, eLabel, dLabel)},
		},
		SubmitText: "",
	}
	keyGenBox := container.NewVBox(widget.NewLabel("Key Generation"), keyGenForm, generateKeysBtn)

	// Encryption Section
	messageInput := widget.NewEntry()
	eInput := widget.NewEntry()
	nInput := widget.NewEntry()
	encryptedLabel := widget.NewLabel("Encrypted: ")

	encryptBtn := widget.NewButton("Encrypt", func() {
		m, ok := new(big.Int).SetString(messageInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid Message"), myWindow)
			return
		}

		e, ok := new(big.Int).SetString(eInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid e"), myWindow)
			return
		}

		n, ok := new(big.Int).SetString(nInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid n"), myWindow)
			return
		}

		c := task8.Encrypt(m, e, n)
		encryptedLabel.SetText("Encrypted: " + c.String())
	})

	encryptForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Message:", Widget: messageInput},
			{Text: "e:", Widget: eInput},
			{Text: "n:", Widget: nInput},
		},
		SubmitText: "",
	}
	encryptBox := container.NewVBox(widget.NewLabel("Encryption"), encryptForm, encryptBtn, encryptedLabel)

	// Decryption Section
	ciphertextInput := widget.NewEntry()
	dInput := widget.NewEntry()
	nDecInput := widget.NewEntry()
	decryptedLabel := widget.NewLabel("Decrypted: ")

	decryptBtn := widget.NewButton("Decrypt", func() {
		c, ok := new(big.Int).SetString(ciphertextInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid Ciphertext"), myWindow)
			return
		}

		d, ok := new(big.Int).SetString(dInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid d"), myWindow)
			return
		}

		n, ok := new(big.Int).SetString(nDecInput.Text, 10)
		if !ok {
			dialog.ShowError(errors.New("Invalid n"), myWindow)
			return
		}

		m := task8.Decrypt(c, d, n)
		decryptedLabel.SetText("Decrypted: " + m.String())
	})

	decryptForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Ciphertext:", Widget: ciphertextInput},
			{Text: "d:", Widget: dInput},
			{Text: "n:", Widget: nDecInput},
		},
		SubmitText: "",
	}
	decryptBox := container.NewVBox(widget.NewLabel("Decryption"), decryptForm, decryptBtn, decryptedLabel)

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
		container.NewTabItem("Task 3",
			container.NewVBox(
				widget.NewLabel("Legendre symbol (a/p)"),
				LegendreForm,
				widget.NewSeparator(),
				widget.NewLabel("Jacobi symbol (a/n)"),
				JacobiForm,
			),
		),

		container.NewTabItem("Task 4",
			container.NewVBox(
				widget.NewLabel("Pollard Rho"),
				PollardRhoForm,
			),
		),

		container.NewTabItem("Task 5",
			container.NewVBox(
				widget.NewLabel("Baby Step Giant Step"),
				BabyStepGiantStepForm,
			),
		),

		container.NewTabItem("Task 6",
			container.NewVBox(
				widget.NewLabel("Cipollo"),
				CipolloForm,
			),
		),

		container.NewTabItem("Task 7",
			container.NewVBox(
				widget.NewLabel("Miller Rabin"),
				MillerRabinForm,
			),
		),
		container.NewTabItem("Task 8",
			container.NewVBox(
				keyGenBox,
				widget.NewSeparator(),
				encryptBox,
				widget.NewSeparator(),
				decryptBox,
			),
		),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()

}
