package pkg

import (
	"goui/util"
	_ "image/png"
	"sort"
	"strings"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/spf13/cast"
)

var (
	key      []string
	mountded = "/run/media/"
	dstpath  = "/mydata/data"
	tablemap = map[string][]string{}
	tablekey = map[int][]string{}
	command  = `df -h |sed '1d'|sed 's/%//g'|awk '{printf $1"|"$2"|"$3"|"$4"|"$5"|"$6"|"localdates"\n"}'`
)

type modelHandler struct {
	mainwin  *ui.Window
	tablemap map[string][]string
	tablekey map[int][]string
}

func newModelHandler(mainwin *ui.Window) *modelHandler {
	m := new(modelHandler)
	m.mainwin = mainwin
	return m
}

func (mh *modelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""), // column 0 text 	 FileSystem
		ui.TableString(""), // column 1 text 	 Size
		ui.TableString(""), // column 2 text 	 Used
		ui.TableString(""), // column 3 text 	 Avail
		ui.TableString(""), // column 4 button   Mount Operation
		ui.TableInt(0),     // column 5 progress Disk Used Percent
	}
}

func (mh *modelHandler) NumRows(m *ui.TableModel) int {
	return len(key)
}

func (mh *modelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	sort.Strings(key)

	switch column {
	case 0:
		return ui.TableString(key[row])
	case 1:
		return ui.TableString(tablemap[key[row]][0])
	case 2:
		return ui.TableString(tablemap[key[row]][1])
	case 3:
		return ui.TableString(tablemap[key[row]][2])
	case 4:
		if !strings.Contains(tablemap[key[row]][4], mountded) && !strings.Contains(tablemap[key[row]][4], dstpath) {
			return ui.TableString("isMounted")
		}
		return ui.TableString(tablemap[key[row]][4])
	case 5:
		use := strings.Replace(tablemap[key[row]][3], "%", "", -1)
		return ui.TableInt(cast.ToInt(use))
	}
	panic("unreachable")
}

func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	if column == 4 {
		// Mount Operation
		if strings.Contains(tablemap[key[row]][4], mountded) { // Contains "/run/media"
			button := ui.NewButton("Message Box")
			window := ui.NewWindow("Mount", 300, 300, false)

			// Mount Disk
			mountcommand := `mount ` + key[row] + " " + dstpath
			util.MountCommand(mountcommand)

			box := ui.NewVerticalBox()
			box.Append(button, false)
			window.SetChild(box)
			ui.MsgBox(window,
				"This is disk mounts message box.",
				"Mounted on: "+mountcommand)

			tablemap[key[row]][4] = dstpath

		} else if strings.Contains(tablemap[key[row]][4], dstpath) {
			button := ui.NewButton("Message Box")
			window := ui.NewWindow("Mount", 300, 300, false)
			box := ui.NewVerticalBox()
			box.Append(button, false)
			window.SetChild(box)

			unmountcommand := `umount -lf ` + dstpath
			util.MountCommand(unmountcommand)

			ui.MsgBox(window,
				"This is disk mounts message box.",
				"UnMounted on: "+unmountcommand)
		}
	}
}

func Table(mainwin *ui.Window) ui.Control {

	vbContainer := ui.NewVerticalBox()
	vbContainer.SetPadded(true)

	buttonGroup := ui.NewGroup("Operation")
	buttonGroup.SetMargined(false)

	vbButtonBox := ui.NewVerticalBox()
	vbButtonBox.SetPadded(true)

	refreshButton := ui.NewButton("Refresh")
	vbButtonBox.Append(refreshButton, false)
	buttonGroup.SetChild(vbButtonBox)

	// table := SetTable(mainwin)
	util.DiskCommand(command, tablemap, tablekey)
	key = util.Getkey(tablemap)
	mh := newModelHandler(mainwin)
	mh.tablemap = tablemap
	mh.tablekey = tablekey

	model := ui.NewTableModel(mh)

	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
		RowBackgroundColorModelColumn: 3,
	})

	table.AppendTextColumn("Filesystem",
		0, ui.TableModelColumnNeverEditable, nil)

	table.AppendTextColumn("Size",
		1, ui.TableModelColumnAlwaysEditable, nil)

	table.AppendTextColumn("Used",
		2, ui.TableModelColumnAlwaysEditable, nil)

	table.AppendTextColumn("Avail",
		3, ui.TableModelColumnAlwaysEditable, nil)

	table.AppendButtonColumn("Mounted On",
		4, ui.TableModelColumnAlwaysEditable)

	table.AppendProgressBarColumn("Use",
		5)

	refreshButton.OnClicked(func(button *ui.Button) {
		tablemap = map[string][]string{}
		tablekey = map[int][]string{}

		util.DiskCommand(command, tablemap, tablekey)
		key = util.Getkey(tablemap)

		ui.MsgBox(mainwin,
			"Refresh DiskUsage",
			"OK!")
	})

	vbContainer.Append(buttonGroup, false)
	vbContainer.Append(table, true)
	return vbContainer
}

func SetTable(mainwin *ui.Window) ui.Control {

	util.DiskCommand(command, tablemap, tablekey)
	key = util.Getkey(tablemap)
	mh := newModelHandler(mainwin)
	model := ui.NewTableModel(mh)

	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
		RowBackgroundColorModelColumn: 3,
	})

	table.AppendTextColumn("Filesystem",
		0, ui.TableModelColumnNeverEditable, nil)

	table.AppendTextColumn("Size",
		1, ui.TableModelColumnAlwaysEditable, nil)

	table.AppendTextColumn("Used",
		2, ui.TableModelColumnAlwaysEditable, nil)

	table.AppendTextColumn("Avail",
		3, ui.TableModelColumnAlwaysEditable, nil)

	table.AppendButtonColumn("Mounted On",
		4, ui.TableModelColumnAlwaysEditable)

	table.AppendProgressBarColumn("Use",
		5)
	return table
}
