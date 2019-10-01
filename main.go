package main

import (
    "flag"
    "fmt"
    "log"
    "math/rand"
    "os"
    "runtime"
    "runtime/pprof"
    "time"

    "github.com/davecgh/go-spew/spew"

    "github.com/alecrabbit/go-cli-spinner"
    "github.com/alecrabbit/go-cli-spinner/color"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        runtime.SetCPUProfileRate(100000) // Needed here this app is too fast
        fmt.Println("SetCPUProfileRate set")
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer fmt.Println("CPU profiling: stop")
        defer pprof.StopCPUProfile()
    }

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

    s, _ := spinner.New(
        spinner.Interval(100),
        // spinner.ColorLevel(color.TNoColor),
        spinner.ColorLevel(color.TColor256),
        spinner.ProgressFormat("[%4s]"), // [  7%]
        // spinner.MessageFormat("%s "), // (message)
    )
    s.FinalMessage = "Done!\n"
    // s.HideCursor = false
    s.Reversed = true
    s.Prefix = " " // spinner prefix

    // spew.Dump(s)
    spew.Dump(nil)
    fmt.Println()

    fmt.Print(".........")
    // Start spinner
    s.Start()
    // for _, m := range messages {
    l := len(messages)
    for i, m := range messages {
        // Doing some work 1
        time.Sleep(duration())
        // Printing execution message
        {
            s.Erase() // optional if you're absolutely sure that your messages are longer
            fmt.Println(m)
            fmt.Print("..................................") // string to show that spinner can be used in inline mode
            s.Current()                                     // Write current frame to output(optional - for smooth amination)
        }
        // Simulating spinner message
        if rand.Intn(16) > 9 {
            s.Message("") // Sometimes there are no messages
        } else {
            s.Message(fmt.Sprintf("Message at %s", time.Now().Format("15:04:05")))
        }
        // Doing some work 2
        time.Sleep(duration())
        // Simulating spinner progress
        s.Progress(float32(i) / float32(l)) // float32 0..1
    }
    time.Sleep(1 * time.Second)
    // Stop spinner
    s.Stop()
    time.Sleep(1 * time.Second)

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close()
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        defer fmt.Println("Memory profiling: stop")
    }
}

func duration() time.Duration {
    return (200 + time.Duration(rand.Intn(1600))) * time.Millisecond
}
