package main

import (
	"fmt"
	"time"

	"github.com/alecrabbit/go-yspnr/spinner"
)

func main() {
	fmt.Println("Started")
	s := spinner.New(1, 120*time.Millisecond)
	fmt.Print("   .")
	s.Start()
	time.Sleep(4*time.Second)
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