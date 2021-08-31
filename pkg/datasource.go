package pkg

import (
	"goui/util"
	"time"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var messageLabel *ui.Label
var timeLabel *ui.Label
var message string
var timemessage string
var messagecount int
var status *ui.ColorButton
var excute bool
var start *ui.DateTimePicker
var stop *ui.DateTimePicker

func MakeDataChoosersPage(mainwin *ui.Window) ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	starttime := ui.NewHorizontalBox()
	starttime.SetPadded(true)
	Separator_label_l := ui.NewLabel("Start Time ")
	starttime.Append(Separator_label_l, false)
	start = ui.NewDateTimePicker()
	starttime.Append(start, false)

	stoptime := ui.NewHorizontalBox()
	stoptime.SetPadded(true)
	Separator_label_2 := ui.NewLabel("Sop  Time ")
	stoptime.Append(Separator_label_2, false)
	stop = ui.NewDateTimePicker()
	stoptime.Append(stop, false)

	timeCounter := ui.NewVerticalBox()
	timeCounter.SetPadded(true)

	timeLabel = ui.NewLabel(timemessage)

	timeCounter.Append(timeLabel, false)

	vbox.Append(starttime, false)
	vbox.Append(stoptime, false)
	vbox.Append(timeCounter, false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	hbox.Append(ui.NewVerticalSeparator(), false)

	// path grid
	grid := ui.NewGrid()
	grid.SetPadded(true)

	// button grid
	buttongrid := ui.NewGrid()
	buttongrid.SetPadded(true)

	countGroup := ui.NewGroup("Counter")
	countGroup.SetMargined(true)

	vbCounter := ui.NewVerticalBox()
	vbCounter.SetPadded(true)

	messageLabel = ui.NewLabel(message)

	vbCounter.Append(messageLabel, false)
	countGroup.SetChild(vbCounter)

	vbox.Append(grid, false)
	vbox.Append(buttongrid, false)
	vbox.Append(countGroup, false)

	// Get Origin path
	inputForm := ui.NewForm()
	inputForm.SetPadded(true)

	originmessage := ui.NewEntry()
	originmessage.SetText("/home/user/rawdata/")
	inputForm.Append("Origin Path:", originmessage, false)

	grid.Append(inputForm,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	// Get Dst path
	inputForm = ui.NewForm()
	inputForm.SetPadded(true)

	dstmessage := ui.NewEntry()
	dstmessage.SetText("/home/user/Desktop/")
	inputForm.Append("Dst Path:    ", dstmessage, false)

	grid.Append(inputForm,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	//Status Color
	status = ui.NewColorButton()
	status.SetColor(1, 1, 1, 1)

	// Start migration button
	button := ui.NewButton("Start")
	button.OnClicked(func(*ui.Button) {
		command := `cp -r ` + originmessage.Text() + " " + dstmessage.Text()
		// command := `ping www.baidu.com`
		go func(m string) {
			util.ExcueteCommand(command, m)
		}(util.Message)
		status.SetColor(0, 1, 0, 1)
		excute = true
		ui.MsgBox(mainwin,
			"Start copy production environment data!",
			command,
		)
	})
	buttongrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	// Stop migration button
	button = ui.NewButton("Stop")
	button.OnClicked(func(*ui.Button) {
		if excute {
			go func(m string) {
				util.ExcuteStopCommand()
			}(util.Message)
			status.SetColor(1, 0, 0, 1)
			excute = false
		} else {
			status.SetColor(1, 1, 1, 1)
			excute = false
		}

		ui.MsgBoxError(mainwin,
			"Stop copy production environment data",
			dstmessage.Text())
	})
	buttongrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	buttongrid.Append(status,
		2, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	go collectMessage()

	return hbox
}

func collectMessage() {
	for {
		time.Sleep(1 * time.Second)
		messagecount++
		if util.ProcessDone && excute {
			status.SetColor(0, 0, 1, 1)
			excute = false
		}

		ui.QueueMain(func() {
			messageLabel.SetText(util.Message + "\n" + time.Now().UTC().Format("2006-01-02 15:04:05.999"))
		})

		ui.QueueMain(func() {
			timeLabel.SetText("Begin: " + start.Time().Format("2006-01-02 15:04:05") + "\nEnd:    " + stop.Time().Format("2006-01-02 15:04:05"))
		})
	}
}
