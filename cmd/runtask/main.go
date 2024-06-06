package make

import (
	"fmt"
	"os"

	"github.com/bartdeboer/runtask"
)

func main() {
	if err := runtask.RunTask(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
