package cli

import (
	"fmt"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/brs"
	"kmschr.com/mc2brs/mca"
)

func Init() {
	ansi.WindowsInitTerminal("Minecraft to Brickadia")
	printIntro()
}

func printIntro() {
	ansi.Print(ansi.BrightGreen, "Minecraft")
	ansi.Print(ansi.BrightWhite, " to ")
	ansi.Println(ansi.BrightRed, "Brickadia")
	fmt.Print("By ")
	ansi.Println(ansi.BrightYellow, "Smallguy")
	ansi.Println(ansi.BrightBlue, "Version 1.0")
}

func ConvertWorld() {
	worldInfo := listWorlds(1)
	ansi.Println(ansi.Yellow, "\nIndexing World...")
	world := mca.NewWorld(worldInfo.Name, worldInfo.Path)

	for {
		if ansi.Dimension() != world.Dimension {
			world = mca.NewWorld(worldInfo.Name, worldInfo.Path)
		}

		world.Overview()
		x, z := world.SelectRegion()
		world.LoadRegion(x, z)
		save := world.ConvertRegion(x, z)
		saveName := fmt.Sprintf("%s_%s_%d_%d", world.Name, world.Dimension, x, z)
		ansi.Println(ansi.BrightYellow, fmt.Sprintf("\nWriting %s.brs...", saveName))
		brs.Write(save, saveName)
		ansi.Println(ansi.BrightBlue, fmt.Sprintf("Wrote %d bricks", len(save.Bricks)))
		mca.PrintUncolored()
		convertAgain := ansi.BasicPrompt("\nConvert another region? (y/n)")

		if convertAgain != "y" && convertAgain != "yes" {
			break
		}
	}

}
