package cli

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"

	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/mca/nbt"
)

type WorldInfo struct {
	Path      string
	Name      string
	Edition   string
	Directory fs.FileInfo
}

func listWorlds(edition int) WorldInfo {
	var world WorldInfo
	if edition == 1 {
		world = userSelectWorldJava()
	} else if edition == 2 {
		world = listWorldsBedrock()
	}
	return world
}

func userSelectWorldJava() WorldInfo {
	savesPath := minecraftJavaPath() + "\\saves"
	worldFiles := getWorldsJava(savesPath)

	ansi.Println(ansi.BrightBlue, "\nWorlds:")
	for i, f := range worldFiles {
		versionName, compatible := readVersionInfo(savesPath + "\\" + f.Name())
		var color ansi.ColorCode
		if compatible {
			color = ansi.BrightGreen
		} else {
			color = ansi.BrightRed
		}

		s := fmt.Sprintf("%v) %v %s", i+1, f.Name(), ansi.Sprint(color, versionName))
		if i%2 == 0 {
			ansi.Println(ansi.BrightWhite, s)
		} else {
			ansi.Println(ansi.White, s)
		}
	}

	selected := 0
	for selected == 0 {
		input := ansi.BasicPrompt("Select world by number:")
		i, err := strconv.Atoi(input)
		if err != nil || i > len(worldFiles) || i <= 0 {
			ansi.Println(ansi.Red, "Invalid world selected")
			continue
		}
		selected = i
	}

	var world WorldInfo
	world.Directory = worldFiles[selected-1]
	world.Edition = "Java"
	world.Name = world.Directory.Name()
	world.Path = savesPath + "\\" + world.Name

	return world
}

func readVersionInfo(path string) (string, bool) {
	levelDat, err := os.Open(path + "\\level.dat")
	if err != nil {
		return "?", false
	}
	rootNBT, err := nbt.ReadGzip(levelDat, false)
	if err != nil {
		return "?", false
	}
	dataNBT, exists := rootNBT.Compound()["Data"]
	if !exists {
		return "?", false
	}
	versionNBT, exists := dataNBT.Compound()["Version"]
	if !exists {
		return "?", false
	}
	nameNBT, exists := versionNBT.Compound()["Name"]
	if !exists {
		return "?", false
	}
	name := nameNBT.String()
	versionNumNBT, exists := dataNBT.Compound()["version"]
	if !exists {
		return name, false
	}
	dataVersionNBT, exists := dataNBT.Compound()["DataVersion"]
	if !exists {
		return name, false
	}
	version := versionNumNBT.Int()
	dataVersion := dataVersionNBT.Int()
	var versionCompatible bool
	if dataVersion < 1952 || version != 19133 {
		versionCompatible = false
	} else {
		versionCompatible = true
	}
	return name, versionCompatible
}

func getWorldsJava(savesPath string) []fs.FileInfo {
	files, err := ioutil.ReadDir(savesPath)
	if err != nil {
		ansi.Println(ansi.Red, "Error reading save files\n")
		ansi.Println(ansi.Red, err.Error())
		ansi.Quit()
	}
	return files
}

func listWorldsBedrock() WorldInfo {
	savesPath := minecraftBedrockPath() + "\\LocalState\\games\\com.mojang\\minecraftWorlds"
	worlds, err := ioutil.ReadDir(savesPath)
	if err != nil {
		ansi.Quit()
	}

	fmt.Println("Select World:")
	for i, world := range worlds {
		namefile := savesPath + "\\" + world.Name() + "\\levelname.txt"
		name := readNameFile(namefile)

		fmt.Printf("%v) %v\n", i+1, name)
	}

	var world WorldInfo
	return world
}

// Bedrock edition keeps world names hidden in a file
func readNameFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()

}
