package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"bufio"
	"time"
	"strings"
	"strconv"
)
//This is the var that if true will automatically put the fucking sleep after every command. You can also set how long the sleep should be for
var autoSleep = false
var sleepAmount = "0.5"
func main() {
	var file string
	loop: for {
		clear()				
		var ans string

		fmt.Println(Figlet("XdoTool", "slant"))
		fmt.Println("auto sleep set to:", autoSleep)
		fmt.Println("Current file selected: ", file)
		fmt.Println("\n0: Close program\n1: Select File\n2: Create File\n3: Get Mouse Location\n4: Add MouseMove To File\n5: Add Sleep To File\n6: Add Click Into File\n7: Add Single Key Press To File \n8: Add Type Instruction To File\n9: Add Enter Key Press To File\n10: Add Any XDoTool Command\n10: Print File Contents\n11: Delete File Row\n12: Execute File\n13: Print File Contents\n14: Toggle Auto Sleep")
		

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
				out := getMouseLocation()
				fmt.Printf("Mouse Location is %s\n", out)
				fmt.Println("Press ENTER to continue")
                                fmt.Scanln(&ans)

			case "4":
				fmt.Println("Would you like to use advanced mode? Y/N?")
				fmt.Println("Selecting N will have you continue to normal mode.")
				fmt.Scanln(&ans)
				if ans == "Y" {
					addMouseMovementAdvanced(file)
				} else {
                                	var x, y string
			        	fmt.Println("This will move your mouse to the x and y coordinates you input.")
                                	fmt.Println("x goes here:")
                                	fmt.Scanln(&x)
					fmt.Println("y goes here:")
                                	fmt.Scanln(&y)

				addMouseMoveToFile(file, x, y)
				}
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
				fmt.Println("If you want to type all the user passed arguments just type \"$*\" \n(This is good for a script you want to type ai output with)")
                                fmt.Println("String go here:")
                                fmt.Scanln(&n)
                                addTypeWordToFile(file, n)

			case "9":
				addEnterKeyToFile(file)

			case "10":
				anyCommandBalls(file)
			case "13":
				printFile(file)
				fmt.Println("Press ENTER to continue")
				fmt.Scanln(&ans)
			case "11":
                               	/*//loop through contents of file line by line
				fileLines := fileLines(file)
				for i := 0; i <= (len(fileLines) - 1); i++{
					fmt.Println(i, ":", fileLines[i])
				}
				fmt.Println("Your file incriments from 0. You can press N to back out.")
                                fmt.Println("Enter which row to remove")
                                //var remove int
				var remove string //get some user input on either which row to remove, or whether to back out
				fmt.Scanln(&remove)
				if remove == "n"  || remove == "N"{
					break
				}
				//since we didn't back out we have no need for our num answer to be a string so we make it an int
				remove1, err := strconv.ParseInt(remove, 6, 12)
				check(err)
    				//fmt.Println(int1)
				removeFromList(file, int(remove1))
				clear()
				//print what the file looks like now and let our user look at it.
				printFile(file)
                                fmt.Println("Press ENTER to continue")
				fmt.Scanln(&ans)*/
				removeRowFromCurrentFile(file)
			case "12":
				executeFileMenu(file)
			/*case "13":
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

                                executeFileInLoopWithDelay(file, limit, amountOfTime)*/
			case "14":
				setAutoSleep()
			case "16":
				//anyCommandBalls(file)
			}
			


	}
}

func removeRowFromCurrentFile(file string) {
	var ans string
	//loop through contents of file line by line
                                fileLines := fileLines(file)
                                for i := 0; i <= (len(fileLines) - 1); i++{
                                        fmt.Println(i, ":", fileLines[i])
                                }
                                fmt.Println("Your file incriments from 0. You can press N to back out.")
                                fmt.Println("Enter which row to remove")
                                //var remove int
                                var remove string //get some user input on either which row to remove, or whether to back out
                                fmt.Scanln(&remove)
                                if remove == "n"  || remove == "N"{
                                        return
                                }
                                //since we didn't back out we have no need for our num answer to be a string so we make it an int
                                remove1, err := strconv.ParseInt(remove, 6, 12)
                                check(err)
                                //fmt.Println(int1)
                                removeFromList(file, int(remove1))
                                clear()
                                //print what the file looks like now and let our user look at it.
                                printFile(file)
                                fmt.Println("Press ENTER to continue")
                                fmt.Scanln(&ans)

}

func executeFileMenu(file string) {
	

	nigger: for {
			var ans string
			clear()
			fmt.Println("Current file selected: ", file)
			fmt.Println("This is the menu for executing scripts.\n0. Exit\n1. execute file normally\n2. execute file in for loop\n3.execute file in delayed for loop")
			fmt.Scanln(&ans)
			switch ans {
			
			case "1":
                                executeFile(file)
                        case "2":
                                var limit int
                                fmt.Println("This will execute your file in a loop, you can select how many times that loop will reiterate.")
                                fmt.Println("type the number of iterations your loop should make:")
                                fmt.Scanln(&limit)
                                executeFileInLoop(file, limit)
                        case "3":
                                var limit int
                                fmt.Println("This will execute your file in a loop, you can select how many times that loop will reiterate, and how many seconds between each iteraction.")
                                fmt.Println("type the number of iterations your loop should make:")
                                fmt.Scanln(&limit)
                                var amountOfTime int
                                fmt.Println("type the amount of seconds between iterations of your loop:")
                                fmt.Scanln(&amountOfTime)

                                executeFileInLoopWithDelay(file, limit, amountOfTime)
	
			case "0":
				break nigger
			}
	}

}

//type ANY command to xdotool to do this I think we'll just print out the list of xdotool commands and a link to the docs. We'll let the user input be all fancy and do whitespaces and shit so the user can just type their command and add it to the file plus the autosleep should work. so swag!
func anyCommandBalls(file string) {
	clear()
	printXdotoolHelp()
	fmt.Println("https://manpages.ubuntu.com/manpages/trusty/man1/xdotool.1.html")
	fmt.Println("Here is a list of all the xdotool commands available and a link to the docs in the ubuntu manpages. \nPlease type N to back out or type the xdotool command you wish to use. \n(please note you won't need to preface every command with xdotool, the tool does this already.)")
	balls := userInput("Type Your Command Here: ")
	if balls == "n" || balls == "N" {
		return
	}
	content := fmt.Sprintf("xdotool %s", balls)
	writeToFileNewLine(file, content)	

}
func printXdotoolHelp() {
        cmd := exec.Command("xdotool", "help")
        cmd.Stdout = os.Stdout
        cmd.Run()


}

//takes user input and returns it as a string
func userInput(userPrompt string) string {

        reader := bufio.NewReader(os.Stdin)
        fmt.Print(userPrompt)
        text, _ := reader.ReadString('\n')

        return strings.TrimSpace(text)

}


func setAutoSleepAmount() {
	
	
	fmt.Println("Type the amount of seconds you want the program to pause between actions. 0.5 for half a second, 1 is 1 second. You get the Idea.")
	fmt.Scanln(&sleepAmount)


}

func setAutoSleep() {

	var ans string
        if autoSleep == false {

		fmt.Println("This is the setting to turn on auto sleep. Auto sleep, automatically puts the sleep instruction in your bash file after every instruction.\nSo you won't need to put it in yourself! Hit Y to enable auto sleep or N to back out.")
		fmt.Scanln(&ans)
		if ans == "N" || ans == "n" {
			return
		}
		autoSleep = true
		setAutoSleepAmount()
	} else {
		fmt.Println("This is the setting to turn off auto sleep. Auto sleep, automatically puts the sleep instruction in your bash file after every instruction.\nSo you won't need to put it in yourself! Hit Y to disable auto sleep or N to back out.")
		fmt.Scanln(&ans)
                if ans == "N"  || ans == "n"{
                        return
                }
                autoSleep = false

	}
}

func addMouseMovementAdvanced(file string) {
	var ans string
	out := getMouseLocation()
        fmt.Printf("Mouse Location is %s\n", out)
	fmt.Println("REMEMBER: This will add the instruction to move the cursor to wherever your mouse is CURRENTLY")
	balls := strings.Split(out, ":")
	fmt.Println("Press ENTER to continue or N to back out")
        
	//balls1 is x, 
	//balls2 is y
	x := strings.Trim(balls[1], "y")
	y := strings.Trim(balls[2], "screen")
	//fmt.Println(balls[1])
	//fmt.Println(balls[2])

	fmt.Scanln(&ans)
	if ans == "n" || ans == "N" {
		return
	}
	//fmt.Println(balls)
	addMouseMoveToFile(file, x, y)
	//fmt.Scanln(&ans)

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
        command := fmt.Sprintf("xdotool key Return")
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

	if autoSleep == true {
	
		command := fmt.Sprintf("sleep %s", sleepAmount)
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
        	check(err)
        	newLine := command
        	_, err = fmt.Fprintln(f, newLine)
        	if err != nil {
        	        fmt.Println(err)
        	        f.Close()
        	        return
        	}
        	err = f.Close()
        	check(err)

	}

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
/*
func moveMouse(x, y string) {

	out, err := exec.Command("xdotool", "mouse move", x, y).Output()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Mouse Location is %s\n", out)


}
*/
func getMouseLocation() string {

	out, err := exec.Command("xdotool", "getmouselocation").Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Mouse Location is %s\n", out)
	return string(out)
}

