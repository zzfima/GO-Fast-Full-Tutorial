package main

import (
	no_goroutines "9_goroutines/cmd_no_goroutines"
	yes_goroutines "9_goroutines/cmd_yes_goroutines"
)

func main() {
	no_goroutines.DoNoGoroutines()
	yes_goroutines.DoYesGoroutines()
}
