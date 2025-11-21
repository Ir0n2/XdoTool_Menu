package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"bufio"
	"time"
	"strings"
)


func main() {
	var file string
	loop: for {
		clear()				
		var ans string

		fmt.Println(Figlet("XdoTool", "slant"))

		fmt.Println("Current file selected: ", file)
		fmt.Println("\n0: Close program\n1: Select File\n2: Create File\n3: Get Mouse Location\n4: Add MouseMove To File\n5: Add Sleep To File\n6: Add Click Into File\n7: Add Single Key Press To File \n8: Add Instruction To Type A String To File\n9: Add Enter Key Press To File\n10: Print File Contents\n11: Delete File Row\n12: Execute File\n13: Execute File In Loop\n14: Execute File In Loop With Delay")
		

		fmt.Scanln(&ans)

		switch ans {
			
			case "0" :
				break loop

			case "1":
				fmt.Println("Current Saved Files:")
                                printCurrentDir()
				fmt.Println("input file name here:")
                                file = selectFile()

			case "2":
				var ans string
				fmt.Println("Current Saved Files:")
				printCurrentDir()
				fmt.Println("new file name here:")
				fmt.Scanln(&ans)
				createNewFile(ans)
			case "3":
                                getMouseLocation()
				fmt.Println("Press ENTER to continue")
                                fmt.Scanln(&ans)

			case "4":
				
                                var x, y string
			        fmt.Println("This will move your mouse to the x and y coordinates you input.")
                                fmt.Println("x goes here:")
                                fmt.Scanln(&x)
				fmt.Println("y goes here:")
                                fmt.Scanln(&y)

				addMouseMoveToFile(file, x, y)
			case "5":
				var num string
				fmt.Println("This will add a break into the order of your operations. You can select how many seconds or do 0.5 for half a second.")
                                fmt.Println("num of seconds goes here:")
                                fmt.Scanln(&num)

				addSleepToFile(file, num)
			case "6":
				var num string
				fmt.Println("Insert Click into file")
				fmt.Println("Buttons  generally  map this way: Left mouse is 1, middle is 2, right is 3, wheel up is 4, wheel down is 5.")
				fmt.Println("Num go here:")
                                fmt.Scanln(&num)
				addClickToFile(file, num)
			case "7":
				var n string
                                fmt.Println("Insert Instruction into file to press a single key")
                                fmt.Println("String go here:")
                                fmt.Scanln(&n)
                                addTypeKeyToFile(file, n)


			case "8":
				var n string
                                fmt.Println("Insert Instruction into file to type a string")
                                fmt.Println("String go here:")
                                fmt.Scanln(&n)
                                addTypeWordToFile(file, n)

			case "9":
				addEnterKeyToFile(file)

			case "10":
				printFile(file)
				fmt.Println("Press ENTER to continue")
				fmt.Scanln(&ans)
			case "11":
                                printFile(file)
                                fmt.Println("Enter which row to remove")
                                var remove int
				fmt.Scanln(&remove)
				removeFromList(file, remove)
				clear()
				printFile(file)
                                fmt.Println("Press ENTER to continue")
				fmt.Scanln(&ans)

			case "12":
				executeFile(file)
			case "13":
				var limit int
				fmt.Println("This will execute your file in a loop, you can select how many times that loop will reiterate.")
				fmt.Println("type the number of iterations your loop should make:")
				fmt.Scanln(&limit)
				executeFileInLoop(file, limit)
			case "14":
				var limit int
				fmt.Println("This will execute your file in a loop, you can select how many times that loop will reiterate, and how many seconds between each iteraction.")
                                fmt.Println("type the number of iterations your loop should make:")
                                fmt.Scanln(&limit)
				var amountOfTime int
                                fmt.Println("type the amount of seconds between iterations of your loop:")
                                fmt.Scanln(&amountOfTime)

                                executeFileInLoopWithDelay(file, limit, amountOfTime)
			}
			


	}
}

func executeFileInLoopWithDelay(file string, something int, balls int) {


        for i := 0; i < something; i++ {
		executeFile(file)
        time.Sleep(time.Duration(balls) * time.Second)
        }

}

func executeFileInLoop(file string, something int) {


	for i := 0; i < something; i++ {
        executeFile(file)
	
	}

}

func executeFile(file string) {

	cmd := exec.Command("bash", file)
        cmd.Stdout = os.Stdout
        cmd.Run()


}
//guess what that fucking does
func printCurrentDir() {
    entries, err := os.ReadDir("./")
    check(err)

    for _, e := range entries {
            if strings.Contains(e.Name(), ".sh") {
                fmt.Println(e.Name())
                } else {
                        continue
                }
    }
}


func removeFromList(file string, remove int) {

	
	fileLines := fileLines(file)

	kill := os.Remove(file)
        create, err := os.Create(file)
        defer create.Close()
        check(kill)
        check(err)
	

	copyArray := make([]string, len(fileLines))
        for i, value := range fileLines {
        	
		if i == remove {
			fmt.Println(value)
			continue
		} else {

			
			copyArray[i] = value
			writeToFileNewLine(file, copyArray[i])
		}
    	}
	
	//return copyArray

}

//declares file lines, which makes an array out of the current list in the text file. It's the back bone of this whole operation.
func fileLines(f string) []string {

	filePath := f
        readFile, err := os.Open(filePath)
        check(err)

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string


	        for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }

        readFile.Close()

	return fileLines
}

func Figlet(text string, font string) (string) {
	// Run `figlet` with the specified text and font
	cmd := exec.Command("figlet", "-f", font, text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(output)
}

func printFile(file string) {
	cmd := exec.Command("cat", file)
        cmd.Stdout = os.Stdout
        cmd.Run()


}
func clear() {

        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}

func addEnterKeyToFile(file string) {
        command := fmt.Sprintf("xdotool type Return")
        writeToFileNewLine(file, command)

}


func addTypeKeyToFile(file, word string) {
        command := fmt.Sprintf("xdotool key %s", word)
        writeToFileNewLine(file, command)

}


func addTypeWordToFile(file, word string) {
        command := fmt.Sprintf("xdotool type %s", word)
        writeToFileNewLine(file, command)

}

func addClickToFile(file, num string) {
        command := fmt.Sprintf("xdotool click  %s", num)
        writeToFileNewLine(file, command)

}

func selectFile() string {

	var file string
        fmt.Println("file name here:")
        fmt.Scanln(&file)

	return file

}

func addSleepToFile(file, num string) {
        command := fmt.Sprintf("sleep %s", num)
        writeToFileNewLine(file, command)

}


func addMouseMoveToFile(file, x, y string) {
	command := fmt.Sprintf("xdotool mousemove %s %s", x, y)
	writeToFileNewLine(file, command)

}

func createNewFile(path string) {

	f, err := os.Create(path)
    	check(err)
	defer f.Close()

}

func writeToFileNewLine(path, content string) {

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
        check(err)
        newLine := content
        _, err = fmt.Fprintln(f, newLine)
        if err != nil {
                fmt.Println(err)
                f.Close()
                return
        }
        err = f.Close()
        check(err)

}

func check(err error) {

	if err != nil {
                log.Println("Error occurred:", err)  
    		return
        }

}

func writeToFile(path, content string) {
	err := os.WriteFile(path, []byte(content), 0666)
	/*if err != nil {
		log.Fatal(err)
	}*/
	check(err)

}

func moveMouse(x, y string) {

	out, err := exec.Command("xdotool", "mouse move", x, y).Output()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Mouse Location is %s\n", out)


}

func getMouseLocation() {

	out, err := exec.Command("xdotool", "getmouselocation").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mouse Location is %s\n", out)

}

