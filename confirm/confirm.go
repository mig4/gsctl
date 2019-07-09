package confirm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Ask asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func Ask(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", color.YellowString(s))

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		switch strings.ToLower(strings.TrimSpace(response)) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		}
	}
}