package ver

import (
	"fmt"
	"runtime"
)

const myver = "0.1"

func GetVersion(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, myver, runtime.Version())
}
