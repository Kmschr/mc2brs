package ansi

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// Enable Windows virtual terminal sequences for console control without safe access to Console API
func WindowsInitTerminal(title string) {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32
	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	windowsSetTitleAndIcon(title)
}

// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences#window-title
func windowsSetTitleAndIcon(title string) {
	fmt.Printf("\x1b]0;%s\x07", title)
}
