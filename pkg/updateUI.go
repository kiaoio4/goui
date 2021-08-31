package pkg

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
)

var countLabel *ui.Label
var count int

func UpdateUI() ui.Control {
	vbContainer := ui.NewVerticalBox()
	vbContainer.SetPadded(true)

	inputGroup := ui.NewGroup("Input")
	inputGroup.SetMargined(true)

	vbInput := ui.NewVerticalBox()
	vbInput.SetPadded(true)

	inputForm := ui.NewForm()
	inputForm.SetPadded(true)

	message := ui.NewEntry()
	message.SetText("Hello World")
	inputForm.Append("What message do you want to show?", message, false)

	showMessageButton := ui.NewButton("Show message")
	clearMessageButton := ui.NewButton("Clear message")

	vbInput.Append(inputForm, false)
	vbInput.Append(showMessageButton, false)
	vbInput.Append(clearMessageButton, false)

	inputGroup.SetChild(vbInput)

	messageGroup := ui.NewGroup("Message")
	messageGroup.SetMargined(true)

	vbMessage := ui.NewVerticalBox()
	vbMessage.SetPadded(true)

	messageLabel := ui.NewLabel("")

	vbMessage.Append(messageLabel, false)

	messageGroup.SetChild(vbMessage)

	countGroup := ui.NewGroup("Counter")
	countGroup.SetMargined(true)

	vbCounter := ui.NewVerticalBox()
	vbCounter.SetPadded(true)

	countLabel = ui.NewLabel(fmt.Sprintf("%d", count))

	vbCounter.Append(countLabel, false)
	countGroup.SetChild(vbCounter)

	vbContainer.Append(inputGroup, false)
	vbContainer.Append(messageGroup, false)
	vbContainer.Append(countGroup, false)

	showMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		messageLabel.SetText(message.Text())
	})

	clearMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		messageLabel.SetText("")
	})

	go counter()

	return vbContainer
}

func counter() {
	for {
		time.Sleep(1 * time.Second)
		count++

		// Update the UI using the QueueMain function
		ui.QueueMain(func() {
			countLabel.SetText(fmt.Sprintf("%d", count))
		})
	}
}
