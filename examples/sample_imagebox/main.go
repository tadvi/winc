package main

import (
	"fmt"

	"github.com/tadvi/winc"
)

func btnOnClick(arg *winc.Event) {
	fmt.Println("Button clicked")
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func main() {
	mainWindow := winc.NewForm(nil)
	dock := winc.NewSimpleDock(mainWindow)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	fileMn.AddItem("New", winc.NoShortcut)
	editMn := menu.AddSubMenu("Edit")
	cutMn := editMn.AddItem("Cut", winc.Shortcut{winc.ModControl, winc.KeyX})
	copyMn := editMn.AddItem("Copy", winc.NoShortcut)
	pasteMn := editMn.AddItem("Paste", winc.NoShortcut)
	menu.Show()
	copyMn.SetCheckable(true)
	copyMn.SetChecked(true)
	pasteMn.SetEnabled(false)

	cutMn.OnClick().Bind(func(e *winc.Event) {
		println("cut click")
	})

	imlistTb := winc.NewImageList(16, 16)
	imlistTb.AddResIcon(10)
	imlistTb.AddResIcon(12)
	imlistTb.AddResIcon(15)

	toolbar := winc.NewToolbar(mainWindow)
	toolbar.SetImageList(imlistTb)
	addBtn := toolbar.AddButton("Add", 1)

	toolbar.AddSeparator()
	runBtn := toolbar.AddButton("Run", 2)

	toolbar.Show()

	runBtn.OnClick().Bind(func(e *winc.Event) {
		println("runBtn click")
	})

	scroll := winc.NewScrollView(mainWindow)
	imgv := winc.NewImageViewBox(scroll)
	scroll.SetChild(imgv)

	addBtn.OnClick().Bind(func(e *winc.Event) {
		if filePath, ok := winc.ShowOpenFileDlg(mainWindow,
			"Select EDI X12 file", "All files (*.*)|*.*", 0, ""); ok {

			if err := imgv.DrawImageFile(filePath); err != nil {
				winc.Errorf(mainWindow, "Error: %s", err)
			}
			scroll.Invalidate(true)
		}
	})

	dock.Dock(toolbar, winc.Top) // toolbars always dock to the top
	dock.Dock(scroll, winc.Fill)

	mainWindow.Center()
	mainWindow.Show()
	dock.Update()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
}
