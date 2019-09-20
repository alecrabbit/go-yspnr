package main

import (
    "fmt"
    "time"

    "github.com/mattn/go-runewidth"

    "github.com/alecrabbit/go-yspnr/spinner"
)

func main() {
    i := spinner.Arrows03

    for a, _ := range spinner.CharSets {
        printCharSet(a)
        fmt.Println()
    }
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
    s := spinner.New(i, 500*time.Millisecond)
    s.FinalMSG = "Done!\n"
    s.HideCursor = false
    s.Start()
    for _, m := range messages {
    	s.Erase()
    	fmt.Println(m)
    	fmt.Print(".......")
    	s.Last()
    	s.Message(time.Now().Format("15:04:05"))
	    time.Sleep(5 * time.Second)
    }
	time.Sleep(1 * time.Second)
	s.Stop()
    fmt.Println("Finished")

}

func printCharSet(n int) {
        var widths []int
        for _, c := range spinner.CharSets[n] {
            width := runewidth.StringWidth(c)
            widths = append(widths, width)
            fmt.Printf("%s %v \n", c, width)
        }
        fmt.Println()
}
