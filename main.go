package main

import (
    "fmt"
    "time"

    "github.com/alecrabbit/go-cli-spinner"
)

func main() {
    messages := []string{
        "Starting",
        "Initializing",
        "Gathering data",
        "Checking data",
        "Checking weather",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Almost there",
        "Be patient",
    }

    s := spinner.New(spinner.Snake3, 100*time.Millisecond)
    s.FinalMessage = "Done!\n"
    // s.HideCursor = false

    s.Start()
    l := len(messages)
    for i, m := range messages {
        // Printing execution message
        {
            s.Erase()
            fmt.Println(m)
            s.Last()
        }
        // Simulating spinner message
        s.Message(fmt.Sprintf("Message at %s", time.Now().Format("15:04:05")))
        // Simulating spinner progress
        f := float32(i) / float32(l)
        s.Progress(f)
        // Doing some work
        time.Sleep(1 * time.Second)
    }
    time.Sleep(3 * time.Second)
    s.Stop()
    // fmt.Println("Finished")
}
