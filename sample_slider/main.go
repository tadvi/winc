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
	//mainWindow.SetLayout(dock)

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
	runBtn := toolbar.AddButton("Run Now Fast", 2)
	toolbar.Show()

	runBtn.OnClick().Bind(func(e *winc.Event) {
		println("runBtn click")
	})

	dock.Dock(toolbar, winc.Top) // toolbars always dock to the top
	//dock.Dock(tree, winc.Fill)

	slide := winc.NewSlider(mainWindow)
	slide.SetPos(10, 50)
	slide.OnScroll().Bind(func(e *winc.Event) {
		println(slide.Value())
	})

	addBtn.OnClick().Bind(func(e *winc.Event) {
		println("addBtn click")
		slide.SetValue(30)
	})

	//track.SetRange(0, 100)
	//track.SetValue(20)

	mainWindow.Center()
	mainWindow.Show()
	dock.Update()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
}
