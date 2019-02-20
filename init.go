package winc

import (
	"github.com/tadvi/winc/w32"
)

func init() {
	gControllerRegistry = make(map[w32.HWND]Controller)
	gRegisteredClasses = make([]string, 0)

	var si w32.GdiplusStartupInput
	si.GdiplusVersion = 1
	w32.GdiplusStartup(&si, nil)
}
