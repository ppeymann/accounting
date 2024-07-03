package pkg

import (
	"fmt"
	"runtime"
)

func getSchemaPath(mod string) string {

	if runtime.GOOS == "windows" {
		return fmt.Sprintf(".\\schemas\\%s", mod)
	}
	return fmt.Sprintf("./schemas/%s", mod)
}
