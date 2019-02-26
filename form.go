/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import (
	"github.com/tadvi/winc/w32"
)

type LayoutManager interface {
	Update()
}

// A Form is main window of the application.
type Form struct {
	ControlBase

	layoutMng LayoutManager
}

func NewForm(parent Controller) *Form {
	fm := new(Form)

	RegClassOnlyOnce("winc_Form")

	fm.isForm = true
	fm.hwnd = CreateWindow("winc_Form", parent, w32.WS_EX_CONTROLPARENT|w32.WS_EX_APPWINDOW, w32.WS_OVERLAPPEDWINDOW)
	fm.parent = parent

	// this might fail if icon resource is not embedded in the binary
	if ico, err := NewIconFromResource(GetAppInstance(), uint16(AppIconID)); err == nil {
		fm.SetIcon(0, ico)
	}

	// This forces display of focus rectangles, as soon as the user starts to type.
	w32.SendMessage(fm.hwnd, w32.WM_CHANGEUISTATE, w32.UIS_INITIALIZE, 0)

	RegMsgHandler(fm)

	fm.SetFont(DefaultFont)
	fm.SetText("Form")
	return fm
}

func (fm *Form) SetLayout(mng LayoutManager) {
	fm.layoutMng = mng
}

// UpdateLayout refresh layout.
func (fm *Form) UpdateLayout() {
	if fm.layoutMng != nil {
		fm.layoutMng.Update()
	}
}

func (fm *Form) NewMenu() *Menu {
	hMenu := w32.CreateMenu()
	if hMenu == 0 {
		panic("failed CreateMenu")
	}
	m := &Menu{hMenu: hMenu, hwnd: fm.hwnd}
	if !w32.SetMenu(fm.hwnd, hMenu) {
		panic("failed SetMenu")
	}
	return m
}

// Public methods
func (fm *Form) Center() {
	sWidth := w32.GetSystemMetrics(w32.SM_CXFULLSCREEN)
	sHeight := w32.GetSystemMetrics(w32.SM_CYFULLSCREEN)
	if sWidth != 0 && sHeight != 0 {
		w, h := fm.Size()
		fm.SetPos((sWidth/2)-(w/2), (sHeight/2)-(h/2))
	}
}

// IconType: 1 - ICON_BIG; 0 - ICON_SMALL
func (fm *Form) SetIcon(iconType int, icon *Icon) {
	if iconType > 1 {
		panic("IconType is invalid")
	}
	w32.SendMessage(fm.hwnd, w32.WM_SETICON, uintptr(iconType), uintptr(icon.Handle()))
}

func (fm *Form) EnableMaxButton(b bool) {
	ToggleStyle(fm.hwnd, b, w32.WS_MAXIMIZEBOX)
}

func (fm *Form) EnableMinButton(b bool) {
	ToggleStyle(fm.hwnd, b, w32.WS_MINIMIZEBOX)
}

func (fm *Form) EnableSizable(b bool) {
	ToggleStyle(fm.hwnd, b, w32.WS_THICKFRAME)
}

func (fm *Form) EnableDragMove(b bool) {
	//fm.isDragMove = b
}

func (fm *Form) EnableTopMost(b bool) {
	tag := w32.HWND_NOTOPMOST
	if b {
		tag = w32.HWND_TOPMOST
	}
	w32.SetWindowPos(fm.hwnd, tag, 0, 0, 0, 0, w32.SWP_NOMOVE|w32.SWP_NOSIZE)
}

func (fm *Form) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_COMMAND:
		if lparam == 0 && w32.HIWORD(uint32(wparam)) == 0 {
			// Menu support.
			actionID := uint16(w32.LOWORD(uint32(wparam)))
			if action, ok := actionsByID[actionID]; ok {
				action.onClick.Fire(NewEvent(fm, nil))
			}
		}

	case w32.WM_KEYDOWN:
		// Accelerator support.
		key := Key(wparam)
		if uint32(lparam)>>30 == 0 {
			// Using TranslateAccelerators refused to work, so we handle them
			// ourselves, at least for now.
			shortcut := Shortcut{ModifiersDown(), key}
			if action, ok := shortcut2Action[shortcut]; ok {
				if action.Enabled() {
					action.onClick.Fire(NewEvent(fm, nil))
				}
			}
		}

	case w32.WM_CLOSE:
		return 0
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
		return 0

	case w32.WM_SIZE, w32.WM_PAINT:
		if fm.layoutMng != nil {
			fm.layoutMng.Update()
		}
	}
	return w32.DefWindowProc(fm.hwnd, msg, wparam, lparam)
}
