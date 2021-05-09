package cli

import (
	"fmt"
	"os"

	"kmschr.com/mc2brs/ansi"
)

func roamingAppDataPath() string {
	return os.Getenv("APPDATA")
}

func localAppDataPath() string {
	return os.Getenv("LOCALAPPDATA")
}

func filePathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func minecraftJavaPath() string {
	minecraftPath := roamingAppDataPath() + "\\.minecraft"

	fmt.Print("\nLooking for ")
	ansi.Print(ansi.Green, "Minecraft ")
	ansi.Print(ansi.BrightRed, "Java Edition ")
	fmt.Println("at:")
	ansi.Println(ansi.Yellow, minecraftPath)

	installed, _ := filePathExists(minecraftPath)
	if !installed {
		ansi.Println(ansi.Red, "Minecraft installation not found, try drag and drop method instead for specific worlds/regions")
		ansi.Quit()
	}

	return minecraftPath
}

func minecraftBedrockPath() string {
	minecraftPath := localAppDataPath() + "\\Packages\\Microsoft.MinecraftUWP_8wekyb3d8bbwe"

	fmt.Print("Looking for ")
	ansi.Print(ansi.Green, "Minecraft ")
	ansi.Print(ansi.BrightBlack, "Bedrock Edition ")
	fmt.Println("at:")
	ansi.Println(ansi.Yellow, minecraftPath)

	installed, _ := filePathExists(minecraftPath)
	if !installed {
		ansi.Println(ansi.Red, "Minecraft installation not found, try drag and drop method instead for specific worlds/regions")
		ansi.Quit()
	}

	return minecraftPath
}
