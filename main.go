package main

import (
    "fmt"
    "time"

    "github.com/alecrabbit/go-cli-spinner"
    "github.com/mattn/go-runewidth"
    "github.com/rivo/uniseg"
)

func main() {
    i := spinner.Snake3
    fmt.Println()
    fmt.Println()
    printCharSet([]string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙",})
    printCharSet([]string{"■", "□", "▪", "▫"})
    fmt.Println()

    printCharSet([]string{"👪", "뢴", "🇩🇪", "ö", "🏳️‍🌈",})
    printCharSet([]string{"🏳️‍🌈", "🇨🇴", "🇧🇼"})
    // for a, _ := range spinner.CharSets {
    //     printCharSet(spinner.CharSets[a])
    //     fmt.Println()
    // }
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
    s := spinner.New(i, 100*time.Millisecond)
    s.FinalMSG = "Done!\n"
    // s.HideCursor = false
    s.Start()
    l := len(messages)
    for i, m := range messages {
    	s.Erase()
    	fmt.Println(m)
    	fmt.Print(".......")
    	s.Last()
    	s.Message(time.Now().Format("15:04:05"))
        f := float32(i) / float32(l)
        // fmt.Printf("%4.0f%%\n", f)
    	s.Progress(f)
	    time.Sleep(5 * time.Second)
    }
	time.Sleep(1 * time.Second)
	s.Stop()
    fmt.Println("Finished")

}

func printCharSet(aw []string) {
        for _, c := range aw {
            width := runewidth.StringWidth(c)

            fmt.Printf("%s %v ", c, width)
            gr := uniseg.NewGraphemes(c)
            for gr.Next() {
                fmt.Printf("%x ", gr.Runes())
            }
            fmt.Println()
        }
        fmt.Println()
}
