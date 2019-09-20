package main

import (
    "fmt"
    "time"

    "github.com/alecrabbit/go-yspnr/spinner"
)

func main() {
    messages := []string{
        "Initializing",
        "Starting",
        "Long message: this message continues further",
        "Gathering data",
        "Short",
        "Short",
        "Processing",
        "Process",
        "Processing",
    }
    fmt.Println("Open > ")
    s := spinner.New(1, 120*time.Millisecond)
    s.FinalMSG = "Done!\n"
    s.Start()
    for _, m := range messages {
    	s.Erase()
    	fmt.Println(m)
    	fmt.Print(".......")
    	s.Last()
    	s.Message(time.Now().Format("2006-01-02 15:04:05"))
    	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	    time.Sleep(1 * time.Second)
    }
	time.Sleep(1 * time.Second)
	s.Stop()
    fmt.Println("Finished")

}

// 	fmt.Print("      .")
//	frame := "🕐"
//	message := " BOLD "
//	w := runewidth.StringWidth(frame + message)
//	fmt.Printf("%v\x1b[1m%v\x1b[0m\x1b[%vD", frame, message, w)
//
//	// time.Sleep(3 * time.Second)
//	fmt.Println()
//	fmt.Println(w)
//
//	fmt.Println(spinner.NoColor, spinner.Color, spinner.Color256, spinner.Truecolor)
//
//
//	time.Sleep(3 * time.Second)
//	ticker.Stop()
//	fmt.Println("Ticker stopped")
//	time.Sleep(2 * time.Second)
//
//	done <- true
