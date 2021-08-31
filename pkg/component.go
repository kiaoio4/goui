package pkg

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/spf13/cast"
)

var (
	radiomap = map[int]string{
		0: "Golang",
		1: "C++",
		2: "Python",
		3: "Other",
	}
)

func Component() ui.Control {
	input := ui.NewEntry()
	input.SetText("this is input element")
	input.LibuiControl()

	processbar := ui.NewProgressBar()
	process := 0
	processbar.SetValue(process)

	messageLabel := ui.NewLabel("Output")

	showMessageButton := ui.NewButton("Show message")
	showMessageButton.SetText("Search")
	// showMessageButton.Enable()
	// showMessageButton.Show()
	// showMessageButton.Visible()

	spinbox := ui.NewSpinbox(50, 150)
	spinbox.SetValue(55)
	slider := ui.NewSlider(0, 100)
	slider.SetValue(10)

	// Separator_div.Append(messageLabel, true)
	combobox := ui.NewCombobox()
	combobox.Append("select one")
	combobox.Append("select two")
	combobox.Append("select three")
	combobox.SetSelected(2)
	checkbox1 := ui.NewCheckbox(radiomap[0])
	checkbox1.SetChecked(true)
	checkbox2 := ui.NewCheckbox(radiomap[1])
	checkbox3 := ui.NewCheckbox(radiomap[2])
	checkbox4 := ui.NewCheckbox(radiomap[3])
	checkbox_div := ui.NewHorizontalBox()
	checkbox_div.Append(checkbox1, true)
	checkbox_div.Append(checkbox2, true)
	checkbox_div.Append(checkbox3, true)
	checkbox_div.Append(checkbox4, true)

	radio := ui.NewRadioButtons()
	radio.Append(radiomap[0])
	radio.Append(radiomap[1])
	radio.Append(radiomap[2])
	radio.Append(radiomap[3])

	checkbox_div.SetPadded(true)
	Separator := ui.NewHorizontalSeparator()
	Separator_label_l := ui.NewLabel("left")
	Separator_label_r := ui.NewLabel("right")
	Separator_div := ui.NewHorizontalBox()
	Separator_div.Append(Separator_label_l, true)
	Separator_div.Append(Separator, false)
	Separator_div.Append(Separator_label_r, true)
	Separator_div.SetPadded(true)
	datetimepicker := ui.NewDateTimePicker()

	// Controller function
	showMessageButton.OnClicked(func(*ui.Button) {
		// Update the UI directly as it is called from the main thread
		// messageLabel.SetText(input.Text())
		process++
		processbar.SetValue(process)
		messageLabel.SetText("Hello, " + input.Text() + "! :" + radiomap[cast.ToInt(radio.Selected())])
	})

	//-----------------Set a single child to a  new group.------------

	container1 := ui.NewGroup("input")
	container1.SetChild(input)
	containre00 := ui.NewGroup("output")
	containre00.SetChild(messageLabel)
	container0 := ui.NewGroup("search")
	container0.SetChild(showMessageButton)
	container2 := ui.NewGroup("spinbox")
	container2.SetChild(spinbox)
	container3 := ui.NewGroup("slider")
	container3.SetChild(slider)
	container4 := ui.NewGroup("processbar")
	container4.SetChild(processbar)
	container5 := ui.NewGroup("checkbox")
	container5.SetChild(checkbox_div)
	container6 := ui.NewGroup("radio")
	container6.SetChild(radio)
	container7 := ui.NewGroup("combobox")
	container7.SetChild(combobox)
	container8 := ui.NewGroup("Separator")
	container8.SetChild(Separator_div)
	container9 := ui.NewGroup("datetimepicker")
	container9.SetChild(datetimepicker)

	//------垂直排列的容器---------
	div := ui.NewVerticalBox()
	//------水平排列的容器
	boxs_0 := ui.NewHorizontalBox()
	boxs_0.SetPadded(true)
	boxs_0.Append(containre00, true)

	boxs_1 := ui.NewHorizontalBox()
	boxs_1.Append(container1, true)
	boxs_1.Append(container0, true)
	boxs_1.Append(container2, true)

	boxs_1.SetPadded(false)
	boxs_2 := ui.NewHorizontalBox()
	boxs_2.Append(container3, true)
	boxs_2.Append(container4, true)

	boxs_3 := ui.NewHorizontalBox()
	boxs_3.Append(container5, true)
	boxs_3.Append(container6, true)

	boxs_4 := ui.NewHorizontalBox()
	boxs_4.Append(container7, true)
	boxs_4.Append(container8, true)

	div.Append(boxs_1, true)
	div.Append(boxs_0, true)
	div.Append(boxs_2, true)
	div.Append(boxs_3, true)
	div.Append(boxs_4, true)
	div.Append(container9, true)
	div.SetPadded(false)

	return div
}
