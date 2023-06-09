package applescript

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "osascript_runner.h"
import "C"
import (
	"fmt"
	"time"
)

func Run(s string) (string, error) {
	time.Sleep(time.Millisecond * 10)

	var resultMessage, errorMessage *C.char
	rc := C.run_osascript(C.CString(s), &resultMessage, &errorMessage)

	if rc == 0 {
		if resultMessage != nil {
			return C.GoString(resultMessage), nil
		} else {
			return "", nil
		}
	}

	return "", fmt.Errorf(C.GoString(errorMessage))
}
