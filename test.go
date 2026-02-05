package main

import (
	"sort"
	"io"
	"strings"
	"context"
	"fmt"
	"os"
	"os/exec"
	"google.golang.org/genai"
	"log"
	"time"
	"math/rand"
	"bufio"
	"path/filepath"
)

func main() {

	//	fmt.Println(geminiAnalyzeImage("/home/zerocool/Pictures/haroldHappy.jpeg", "Describe this image in a few words."))
	var ans string
	looper: for {
		clear()
		fmt.Println(Figlet("Automata AI pipeline maker", "slant"))
		fmt.Println("1. Basic Chat Bot Setup\n2. ai caption/automatic post pipeline menu\n3. analyzeImage Text \n4. screenshot test\n0. to quit")
		fmt.Scanln(&ans)

		switch ans {

		case "1":
			clear()
			fmt.Println("Here are all sh files in current Dir: ")
			printCurrentDir()
			fileForTypingShit := userInput("enter file for typing in chat:")
			clear()
			fmt.Println("Now Enter a prompt, Here's an example:")
			fmt.Println("Read the discord messages on screen and respond to themm accordingly, whatever you respond with will be said in chat. If you say continue, it will not type what you say. You are the user with the name 'ThePumpkinEater', if you are the last to respond say 'continue' otherwise respond normally like you are part of the conversation")
			prompt := userInput("enter the prompt here: ")
			clear()
			basicChatBotSetup(fileForTypingShit, prompt)

			//basicChatBotLoop("/home/zerocool/automata/scriptForTypingOnDiscord.sh")
		case "2":
				captionImageAndPostMenu()
		case "3":
			 fmt.Println(geminiAnalyzeImage("/home/zerocool/Pictures/paulrudd.jpg", "analyze this image"))
			 fmt.Scanln(&ans)
		case "0":
			break looper

		}
	}
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

func captionImageAndPostMenu() {
	clear()
	var ans string

	wenis:	for {

	fmt.Println("Here is the menu for automating captioning/posting pictures.")
	fmt.Println("1. Return back to menu\n2. ai caption/automatic post pipeline test\n3. one automatically captioned post per day\n4. ai captioning/posting pipeline with user set delay in minutes\n")

	//var ans string
        
	
                
		
		fmt.Scanln(&ans)

		switch ans {
			
		case "2":
                        clear()
                        fmt.Println("Here are all sh files in current Dir: ")
                        printCurrentDir()
                        fmt.Println("Please note this posting bot should be used for posting to instagram/facebook or reddit (maybe youtube?)")
                        fileForPostingShit := userInput("enter file for posting to website:")
                        clear()
                        fmt.Println("Now Enter a prompt, Here's an example:")
                        fmt.Println("analyze the picture and caption it for instagram, just say the caption and maybe relevant hashtags.")
                        prompt := userInput("enter the prompt here: ")

                        fmt.Println("Now Enter image path, Here's an example:")
                        fmt.Println("/home/zerocool/Pictures/badassSkeleton.png")
                        filePath := userInput("enter the file path for the image here: ")

                        clear()
                        basicCaptionImageAndPostItSetup(fileForPostingShit, filePath, prompt)
                        //screenshot("screenshot.png")
                case "3":
                        clear()
                        fmt.Println("THIS MAKES 1 POST A DAY AT A CERTAIN TIME IN A LOOP")
                        fmt.Println("Here are all sh files in current Dir: ")
                        printCurrentDir()
                        fmt.Println("Please note this posting bot should be used for posting to instagram/facebook or reddit (maybe youtube?)")
                        fileForPostingShit := userInput("enter file for posting to website:")
                        clear()
                        fmt.Println("Now Enter a prompt, Here's an example:")
                        fmt.Println("analyze the picture and caption it for instagram, just say the caption and maybe relevant hashtags.")
                        prompt := userInput("enter the prompt here: ")

                        fmt.Println("Now Enter folder for images, Here's an example:")
                        fmt.Println("/home/zerocool/Pictures/postQueue")
                        filePath := userInput("enter the file path here: ")
                        var hour, minute int
                        fmt.Println("Set the hour (24 hours 24 being 12 pm, 1 being 1 am)")
                        fmt.Scanln(&hour)
                        fmt.Println("and the minutes")
                        fmt.Scanln(&minute)
                        clear()
                        runScheduler(filePath, fileForPostingShit, prompt, hour, minute)
		
		case "4":
                        clear()
                        fmt.Println("Here are all sh files in current Dir: ")
                        printCurrentDir()
                        fmt.Println("Please note this posting bot should be used for posting to instagram/facebook or reddit (maybe youtube?)")
                        fileForPostingShit := userInput("enter file for posting to website:")
                        clear()
                        fmt.Println("Now Enter a prompt, Here's an example:")
                        fmt.Println("analyze the picture and caption it for instagram, just say the caption and maybe relevant hashtags.")
                        prompt := userInput("enter the prompt here: ")

                        fmt.Println("Now Enter folder path for your images, Here's an example:")
                        fmt.Println("/home/zerocool/Pictures/postQueue")
                        filePath := userInput("enter the file path for the folder here: ")
                        var num int
                        fmt.Println("how long should the delay be? (type an integer. this will be how many minutes between each post)")
                        fmt.Scanln(&num)
                        clear()
                        postFolderInOrder(filePath, fileForPostingShit, prompt, num)

		case "0":
			break wenis
			
		}
	
	}

}

//reminder to make an automatic posting/captioning system that will sequentially post all the photos in a folder throughout a given time span at a given rate.
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
func check(err error) {

        if err != nil {
                log.Println("Error occurred:", err)
                return
        }

}

//takes user input and returns it as a string
func userInput(userPrompt string) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userPrompt)
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)

}

func CopyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    if err != nil {
        return err
    }

    return nil
}

func postFolderInOrder(folder, script, prompt string, num int) {
    files, err := getImages(folder)
    if err != nil {
        log.Fatal(err)
    }

    sortImages(files, "name") // or modtime/random

    for _, file := range files {
        fmt.Println("Posting:", file)
        basicCaptionImageAndPostItSetup(script, file, prompt)

        // delete or move after posting
        os.Remove(file)

        // optional delay between posts
        time.Sleep(time.Duration(num) * time.Minute)
    }
}


//post once a day at a certain time
func runScheduler(folder, script, prompt string, hour, minute int) {
    for {
        waitUntil(hour, minute)

        files, err := getImages(folder)
        if err != nil {
            fmt.Println("Error reading images:", err)
            continue
        }

        if len(files) == 0 {
            fmt.Println("No images to post, waiting for tomorrow")
            continue
        }

        sortImages(files, "name")

        file := files[0] // next image
        fmt.Println("Posting:", file)

        basicCaptionImageAndPostItSetup(script, file, prompt)
        os.Remove(file)

        // After posting, sleep until next day loop
    }
}

func getImages(folder string) ([]string, error) {
    entries, err := os.ReadDir(folder)
    if err != nil {
        return nil, err
    }

    var files []string
    for _, entry := range entries {
        if !entry.IsDir() {
            name := entry.Name()
            // filter by image types
            if strings.HasSuffix(name, ".png") ||
               strings.HasSuffix(name, ".jpg") ||
               strings.HasSuffix(name, ".jpeg") {
                files = append(files, filepath.Join(folder, name))
            }
        }
    }

    return files, nil
}


func waitUntil(hour, minute int) {
    now := time.Now()
    next := time.Date(
        now.Year(), now.Month(), now.Day(),
        hour, minute, 0, 0, now.Location(),
    )

    if next.Before(now) {
        next = next.Add(24 * time.Hour)
    }

    time.Sleep(next.Sub(now))
}
func sortImages(files []string, mode string) {
    switch mode {
    case "name":
        sort.Strings(files)

    case "modtime":
        sort.Slice(files, func(i, j int) bool {
            infoI, _ := os.Stat(files[i])
            infoJ, _ := os.Stat(files[j])
            return infoI.ModTime().Before(infoJ.ModTime())
        })

    case "random":
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(files), func(i, j int) {
            files[i], files[j] = files[j], files[i]
        })
    }
}


//Here's the basic script for a posting bot pipeline
func basicCaptionImageAndPostItSetup(scriptToPostShit, imageFile, aiPrompt string) {
        fmt.Println("Starting bot script")
        //for our botscript to find the image, the image needs to be in the expected folder, so we're going to copy it there, then delete it after we post it!
	/*dst := fmt.Sprintf("/home/zerocool/Pictures/posts/%s", imageFile)
	*/

	  // imageFile is a full path. Keep it as source:
    src := imageFile

    // Extract just the filename ONLY
    filename := filepath.Base(imageFile)

    // Build correct destination path
    dst := filepath.Join("/home/zerocool/Pictures/posts", filename)

    fmt.Println("SRC:", src)
    fmt.Println("DST:", dst)

    	err := CopyFile(imageFile, dst)
    	if err != nil {
    	    log.Fatalf("File copy failed: %v", err)
    	}
    	log.Println("File copied successfully!")
        os.Remove("screenshot.png")
        //screenshot("screenshot.png")
        //fmt.Println("screenshotted")
        //output string of ai response
        output := geminiAnalyzeImage(imageFile, aiPrompt)
        fmt.Println("analyzed by gemini")
        //this should be our caption baby
	fmt.Println(output)
        /*if strings.Contains(output, "continue") {
                fmt.Println("continueing")
                continue
        }*/
        //go already has a remove file func so i dont need to make my own.
        //os.Remove("screenshot.png")
        //fmt.Println("screen unshotted")
        executeBashScript(scriptToPostShit, output)
        fmt.Println("executed script")
       	os.Remove(dst)
	fmt.Println("tied up loose ends!")
}


//Here's the basic script for a chat bot
func basicChatBotSetup(scriptToTypeShit, aiPrompt string) {
        fmt.Println("Starting bot script")
	for {
	os.Remove("screenshot.png")	
        screenshot("screenshot.png")
        fmt.Println("screenshotted")
        //output string of ai response
        output := geminiAnalyzeImage("screenshot.png", aiPrompt)
        fmt.Println("analyzed by gemini")
        fmt.Println(output)
        if strings.Contains(output, "continue") {
                fmt.Println("continueing")
                continue
        }
        //go already has a remove file func so i dont need to make my own.
        os.Remove("screenshot.png")
        fmt.Println("screen unshotted")
        executeBashScript(scriptToTypeShit, output)
        fmt.Println("executed script")
        }
}

func randomSleep() {
    // Seed the RNG once
    rand.Seed(time.Now().UnixNano())

    // Random minutes between 1 and 5
    mins := rand.Intn(5) + 1 // 1..5
    dur := time.Duration(mins) * time.Minute

    fmt.Println("Sleeping for:", dur)
    time.Sleep(dur)
}
//make a pipeline for posting stuff to instagram.
// ai look at / caption pic -> feed teh caption to script we use to post it -> money

//make a basic chat bot like you wood for discord, you know the kind, feed gemini a pic of chat + prompt how to deal with responses and make sure it knows to respond like it's texting on discord, because it is lol. then send the ai response on discord, using the automation script.
func basicChatBotLoop(scriptToTypeShit string) {
	for {
	os.Remove("screenshot.png")
	screenshot("screenshot.png")
	fmt.Println("screenshotted")
	//output string of ai response
	output := geminiAnalyzeImage("screenshot.png", "Read the discord messages on screen and respond to themm accordingly, whatever you respond with will be said in chat. If you say continue, it will not type what you say. You are the user with the name 'ThePumpkinEater', if you are the last to respond say 'continue' otherwise respond normally like you are part of the conversation, try to respond like a human. Make fun of Anon-Monroe, tell him how you're more tech savvy than he is. He's a skid who only uses kali linux and hak5 tools. Can't write his own code, etc")
	fmt.Println("analyzed by gemini")
	fmt.Println(output)
	if strings.Contains(output, "continue") {
		fmt.Println("continueing")
		continue
	}
	//go already has a remove file func so i dont need to make my own.
	os.Remove("screenshot.png")
	fmt.Println("screen unshotted")
	executeBashScript(scriptToTypeShit, output)
	fmt.Println("executed script")
	randomSleep()
	}
}

func executeBashScript(scriptPath, shitToType string) {
	/*command := fmt.Sprintf("bash %s", scriptPath)
        cmd := exec.Command(command)
        cmd.Stdout = os.Stdout
        cmd.Run()*/

	cmd := exec.Command("bash", scriptPath, shitToType)
    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println(string(output))
}


//clear screen
func clear() {

        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}
//save screenshot at filepath
func screenshot(filepath string) {
	/*command := fmt.Sprintf("scrot %s", filepath)
        cmd := exec.Command(command)
        cmd.Stdout = os.Stdout
        cmd.Run()*/
	cmd := exec.Command("scrot", filepath)
    	cmd.Run()

}

func geminiAnalyzeImage(path, input string) string {
    ctx := context.Background()

    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  "AIzaSyDllaKsIjJsRYvGtKt9g-QT0tEiP0VVNiQ",
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Println("client error:", err)
        return "(client error)"
    }

    uploadedFile, err := client.Files.UploadFromPath(ctx, path, nil)
    if err != nil {
        log.Println("upload error:", err)
        return "(upload error)"
    }

    parts := []*genai.Part{
        genai.NewPartFromText(input),
        genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType),
    }

    contents := []*genai.Content{
        genai.NewContentFromParts(parts, genai.RoleUser),
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        contents,
        nil,
    )
    if err != nil {
        log.Println("gemini error:", err)
        return "(gemini api error)"
    }

    // Defensive checks to prevent nil pointer access
    if result == nil {
        return "(nil response)"
    }
    if len(result.Candidates) == 0 {
        return "(no candidates)"
    }
    if result.Candidates[0].Content == nil {
        return "(empty content)"
    }
    if len(result.Candidates[0].Content.Parts) == 0 {
        return "(empty parts)"
    }

    // Now safe
    return result.Text()
}


/*
// uses gemini to analyze screenshot (so we can what's on screen.)
func geminiAnalyzeImage(path, input string) string {
	/*
	  ctx := context.Background()
	  client, err := genai.NewClient(ctx, nil)
	  if err != nil {
	      log.Fatal(err)
	  }
	*//*

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  "AIzaSyDllaKsIjJsRYvGtKt9g-QT0tEiP0VVNiQ",
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	uploadedFile, _ := client.Files.UploadFromPath(ctx, path, nil)

	parts := []*genai.Part{
		genai.NewPartFromText(input),
		genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		contents,
		nil,
	)

	return result.Text()
}*/
