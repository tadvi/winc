package main

import (
	"fmt"

	"github.com/tadvi/winc"
)

func btnOnClick(arg *winc.Event) {
	//edt.SetCaption("Got you !!!")
	fmt.Println("Button clicked")
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func main() {
	mainWindow := winc.NewForm(nil)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	//none := winc.Shortcut{}

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	fileMn.AddItem("New", winc.NoShortcut)
	editMn := menu.AddSubMenu("Edit")
	delMn := editMn.AddItem("Delete", winc.Shortcut{winc.ModControl, winc.KeyX})
	delAllMn := editMn.AddItem("Delete All", winc.NoShortcut)
	menu.Show()

	delMn.OnClick().Bind(func(e *winc.Event) {
		dlg := winc.NewDialog(mainWindow)
		dlg.Center()
		dlg.Show()
	})

	delAllMn.OnClick().Bind(func(e *winc.Event) {
		dlg := winc.NewDialog(mainWindow)
		dlg.Center()
		dlg.Show()
	})

	toolbar := winc.NewPanel(mainWindow)
	toolbar.SetPos(0, 0)
	toolbar.SetSize(100, 40)

	btnRun := winc.NewIconButton(toolbar)
	btnRun.SetText(" Run")
	btnRun.SetPos(2, 2)
	btnRun.SetSize(98, 38)
	btnRun.SetResIcon(15)

	btnRun.OnClick().Bind(func(e *winc.Event) {
		println("event OnClick")
	})

	//tipRun := winc.NewToolTip(mainWindow)
	//tipRun.AddTool(btnRun, "Run project")

	btnEdit := winc.NewPushButton(toolbar)
	btnEdit.SetText(" Edit")
	btnEdit.SetPos(102, 2)
	btnEdit.SetSize(98, 38)
	btnEdit.SetResIcon(18)

	left := winc.NewMultiEdit(mainWindow)
	left.SetPos(5, 5)
	right := winc.NewMultiEdit(mainWindow)
	right.SetPos(50, 100)

	split := winc.NewVResizer(mainWindow)
	split.SetControl(left, right, winc.Left, 150)

	// --- Add controls to Dock, LoadStateFile and Show window in this order
	mainWindow.Center()
	mainWindow.Show()

	dock := winc.NewSimpleDock(mainWindow)
	//mainWindow.SetLayout(dock)
	dock.Dock(toolbar, winc.Top)
	dock.Dock(left, winc.Left)
	dock.Dock(split, winc.Left)
	dock.Dock(right, winc.Fill)
	// if err := dock.LoadStateFile("layout.json"); err != nil {
	// 	log.Println(err)
	// }

	mainWindow.OnClose().Bind(func(e *winc.Event) {
		dock.SaveStateFile("layout.json") // error gets ignored
		winc.Exit()
	})

	dock.Update()

	winc.RunMainLoop()
	// --- end of Dock and main window management

}
