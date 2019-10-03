package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/davecgh/go-spew/spew"
    // "github.com/pkg/profile"

    "github.com/alecrabbit/go-cli-spinner"
    "github.com/alecrabbit/go-cli-spinner/color"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    // profile.CPUProfile, profile.MemProfile, profile.TraceProfile, profile.MutexProfile
    // defer profile.Start(profile.TraceProfile, profile.ProfilePath("./profiling")).Stop()
    // defer profile.Start(profile.CPUProfile, profile.ProfilePath("./profiling")).Stop()
    // defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath("./profiling")).Stop()

    messages := map[int]string{
        0:  "Starting",
        3:  "Initializing",
        6:  "Gathering data",
        9:  "Processing",
        16: "Processing",
        25: "Processing",
        44: "Processing",
        60: "Processing",
        79: "Processing",
        82: "Still processing",
        90: "Be patient",
        95: "Almost there",
    }

    s, err := spinner.New(
        spinner.Interval(100),
        // spinner.ColorLevel(color.TNoColor),
        spinner.ColorLevel(color.TColor256),
        spinner.ProgressFormat("[%4s]"), // [  7%]
        spinner.MessageFormat("(%s)"),   // (message)
        // spinner.Format("%s "),         //
    )
    if err != nil {
        log.Fatal(err)
    }
    s.FinalMessage = "Done!\n"
    s.HideCursor = true
    s.Reversed = true
    s.Prefix = " " // spinner prefix

    // spew.Dump(s)
    spew.Dump(nil)
    fmt.Println()

    fmt.Print("..................................")
    // Start spinner
    s.Start()
    for i := 0; i <= 100; i++ {
        // Doing some work 1
        time.Sleep(duration())
        if m, ok := messages[i]; ok {
            // Printing execution message
            {
                s.Erase() // optional if you're absolutely sure that your messages are longer
                fmt.Println(m)
                fmt.Print("..................................") // string to show that spinner can be used in inline mode
                s.Current()                                     // Write current frame to output(optional - for smooth amination)
            }

        }
        // Simulating spinner message
        if i > 10 && i < 20 || i > 40 && i < 60 {
            s.Message("") // Sometimes there are no messages
        } else {
            s.Message(fmt.Sprintf("Status message at %s", time.Now().Format("15:04:05")))
        }
        // Doing some work 2
        time.Sleep(duration())
        // Simulating spinner progress
        if i > 50 && i < 70 {
            s.Progress(0)
        } else {
            s.Progress(float32(i) / float32(100)) // float32 0..1
        }

    }
    // Stop spinner
    s.Stop()
    time.Sleep(500 * time.Millisecond)
}

func duration() time.Duration {
    return (50 + time.Duration(rand.Intn(50))) * time.Millisecond
}
