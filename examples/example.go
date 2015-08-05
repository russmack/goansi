package main

import (
	"fmt"
	"github.com/russmack/goansi"
)

func main() {

	/*
	 * Fluent interface
	 */

	// Create anstring, add style later.
	a1 := goansi.NewAnstring("I am blue.")
	fmt.Println(a1.Blue().Underline())

	// Create anstring and apply style at the same time.
	a2 := goansi.Anstring("I am white on cyan.").Gray().Bold().BgCyan()
	fmt.Println(a2)

	// Create anstring, style, and print inline.
	s := "I am blinky black on green."
	fmt.Println(goansi.Anstring(s).Black().BgGreen().Underline().BlinkSlow())

	/*
	 * Create style, then render string.
	 */

	//style := goansi.NewAnstyle().Yellow().BgRed()
	style := goansi.NewStyle().Red().BgYellow()
	fmt.Println(style.Render("I am Iron Man!"))

	/* TODO: maybe.
	 * Nested methods
	 */
	//fmt.Println(goansi.Yellow(goansui.BgBlack("I am iron man!")))

	/*
	 * Utility
	 */

	// Print the ansi escape codes.
	fmt.Println("Raw:", a2.Raw())
}
