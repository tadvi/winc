package winc

import (
	"syscall"

	"github.com/tadvi/winc/w32"
)

//Private global variables.
var (
	gAppInstance        w32.HINSTANCE
	gControllerRegistry map[w32.HWND]Controller
	gRegisteredClasses  []string
)

//Public global variables.
var (
	GeneralWndprocCallBack = syscall.NewCallback(generalWndProc)
	DefaultFont            *Font
)
