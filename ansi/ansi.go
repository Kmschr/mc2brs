package ansi

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"
)

type ColorCode string

const (
	Black         ColorCode = "\x1b[30m"
	Red           ColorCode = "\x1b[31m"
	Green         ColorCode = "\x1b[32m"
	Yellow        ColorCode = "\x1b[33m"
	Blue          ColorCode = "\x1b[34m"
	Magenta       ColorCode = "\x1b[35m"
	Cyan          ColorCode = "\x1b[36m"
	White         ColorCode = "\x1b[37m"
	BrightBlack   ColorCode = "\x1b[90m"
	BrightRed     ColorCode = "\x1b[91m"
	BrightGreen   ColorCode = "\x1b[92m"
	BrightYellow  ColorCode = "\x1b[93m"
	BrightBlue    ColorCode = "\x1b[94m"
	BrightMagenta ColorCode = "\x1b[95m"
	BrightCyan    ColorCode = "\x1b[96m"
	BrightWhite   ColorCode = "\x1b[97m"

	reset = "\x1b[0m"
)

var quitCommands = map[string]bool{
	"quit":  true,
	"q":     true,
	"exit":  true,
	"stop":  true,
	"bye":   true,
	"close": true,
}

var trueFalse = map[string]bool{
	"true":  true,
	"t":     true,
	"false": false,
	"f":     false,
}

var dimensions = map[string]string{
	"overworld": "overworld",
	"o":         "overworld",
	"nether":    "nether",
	"n":         "nether",
	"end":       "end",
	"e":         "end",
}

var scale = 13
var optimize = true
var dimension = "overworld"
var lights = true

func Print(color ColorCode, s string) {
	fmt.Printf("%s%s%s", color, s, reset)
}

func Println(color ColorCode, s string) {
	fmt.Printf("%s%s%s\n", color, s, reset)
}

func Sprint(color ColorCode, s string) string {
	return fmt.Sprintf("%s%s%s", color, s, reset)
}

func PrintRGB(rgb color.RGBA, s string) {
	fmt.Printf("\x1b[38;2;%d;%d;%dm%s%s", rgb.R, rgb.G, rgb.B, s, reset)
}

func PrintlnRGB(rgb color.RGBA, s string) {
	fmt.Printf("\x1b[38;2;%d;%d;%dm%s%s\n", rgb.R, rgb.G, rgb.B, s, reset)
}

func SprintRGB(rgb color.RGBA, s string) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s%s", rgb.R, rgb.G, rgb.B, s, reset)
}

func Color(rgb color.RGBA) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)
}

func RGB(r int, g int, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func BasicPrompt(msg string) string {
	Println(BrightYellow, msg)
	fmt.Print("-> ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return ""
	}
	line = simplifyInput(line)
	if quitCommands[line] {
		os.Exit(0)
	}

	args := strings.Split(line, " ")
	if len(args) == 0 || args[0] == "" {
		return ""
	}

	if strings.HasPrefix("help", simplifyInput(args[0])) {
		Help()
		return BasicPrompt(msg)
	}
	if strings.HasPrefix("convert", simplifyInput(args[0])) {
		return "convert"
	}
	if strings.HasPrefix("scale", simplifyInput(args[0])) {
		if len(args) != 2 {
			Println(BrightRed, "Invalid scale")
			return BasicPrompt(msg)
		}
		newScale, err := strconv.Atoi(args[1])
		if err != nil || newScale <= 0 || newScale > 16 {
			Println(BrightRed, "Invalid scale")
			return BasicPrompt(msg)
		}
		scale = newScale
		Println(BrightBlue, fmt.Sprintf("Scale set to %d", scale))
		return BasicPrompt(msg)
	}

	if strings.HasPrefix("optimize", simplifyInput(args[0])) {
		if len(args) != 2 {
			Println(BrightRed, "Invalid optimization option")
			return BasicPrompt(msg)
		}
		opt, exists := trueFalse[args[1]]
		if !exists {
			Println(BrightRed, "Invalid optimization option")
			return BasicPrompt(msg)
		}
		optimize = opt
		Println(BrightBlue, fmt.Sprintf("Optimization set to %v", optimize))
		return BasicPrompt(msg)
	}

	if strings.HasPrefix("dimension", simplifyInput(args[0])) {
		if len(args) != 2 {
			Println(BrightRed, "Invalid dimension option")
			return BasicPrompt(msg)
		}
		opt, exists := dimensions[args[1]]
		if !exists {
			Println(BrightRed, "Invalid dimension option")
			return BasicPrompt(msg)
		}
		dimension = opt
		Println(BrightBlue, fmt.Sprintf("Dimension set to %s", dimension))
		return BasicPrompt(msg)
	}

	if strings.HasPrefix("lights", simplifyInput(args[0])) {
		if len(args) != 2 {
			Println(BrightRed, "Invalid lighting option")
			return BasicPrompt(msg)
		}
		light, exists := trueFalse[args[1]]
		if !exists {
			Println(BrightRed, "Invalid lighting option")
			return BasicPrompt(msg)
		}
		lights = light
		Println(BrightBlue, fmt.Sprintf("Lighting set to %v", lights))
		return BasicPrompt(msg)
	}

	return line
}

func Help() {
	Println(BrightBlue, "\nCommands:")
	Print(BrightWhite, "c convert\t")
	Println(BrightGreen, "select a world to convert")
	Print(BrightWhite, "q quit\t\t")
	Println(BrightGreen, "exits program at any point")
	Print(BrightWhite, "h help\t\t")
	Println(BrightGreen, "displays info about commands")
	Print(BrightWhite, "s scale\t ")
	Print(BrightCyan, "x\t")
	Println(BrightGreen, "set the size of converted blocks (1-16), default 13")
	Print(BrightWhite, "o optimize ")
	Print(BrightCyan, "t/f\t")
	Println(BrightGreen, "enable optimization to reduce brick count, default true")
	Print(BrightWhite, "d dimension ")
	Print(BrightCyan, "overworld/nether/end\t")
	Println(BrightGreen, "which dimension to load regions from, default overworld")
	Print(BrightWhite, "l lights ")
	Print(BrightCyan, "t/f\t")
	Println(BrightGreen, "generate light components for light emitting blocks, default true")
}

func Scale() int {
	return scale
}

func Optimize() bool {
	return optimize
}

func Dimension() string {
	return dimension
}

func Lights() bool {
	return lights
}

func simplifyInput(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func Quit() {
	f := bufio.NewWriter(os.Stdout)
	f.Flush()
	Println(White, "\nPress any key to quit...")
	in := bufio.NewReader(os.Stdin)
	_, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}
