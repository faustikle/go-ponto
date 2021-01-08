package errorhandler

import (
	"fmt"
	"os"
)

func Handler(err error) {
	fmt.Println(err)
	os.Exit(1)
}
