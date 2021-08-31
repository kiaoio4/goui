package main

import (
	"goui/pkg"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func SetupUI() {
	mainwin := ui.NewWindow("KIGA Data Governance", 920, 600, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Disk Usage", pkg.Table(mainwin))
	tab.SetMargined(0, true)

	tab.Append("Data Migration", pkg.MakeDataChoosersPage(mainwin))
	tab.SetMargined(1, true)

	tab.Append("Basic Controls", pkg.MakeBasicControlsPage())
	tab.SetMargined(2, true)

	tab.Append("UpdateUI", pkg.UpdateUI())
	tab.SetMargined(3, true)

	tab.Append("Component", pkg.Component())
	tab.SetMargined(4, true)

	tab.Append("Histogram", pkg.Histogram())
	tab.SetMargined(5, true)

	tab.Append("Numbers and Lists", pkg.MakeNumbersPage())
	tab.SetMargined(6, true)

	mainwin.Show()
}
