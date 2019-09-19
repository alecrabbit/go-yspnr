package spinner

import (
    "container/ring"
    "fmt"
    "io"
    "os"
    "sync"
    "time"
)

func init() {
    fmt.Println("Init")
}

type ColorLevel int

const (
    NoColor            = iota
    Color   ColorLevel = 1 << (2 * 2 * iota)
    Color256
    Truecolor
)

// Spinner struct to hold the provided options
type Spinner struct {
    Interval   time.Duration // Delay is the speed of the indicator
    frames     *ring.Ring      // chars holds the chosen character set
    active     bool          // active holds the state of the spinner
    FinalMSG   string        // string displayed after Stop() is called
    colorLevel ColorLevel
    lock       *sync.RWMutex //
    Writer     io.Writer     // to make testing better, exported so users have access
    stop       chan bool     // stopChan is a channel used to stop the indicator
    HideCursor bool          // hideCursor determines if the cursor is visible
    // lastOutput string                        // last character(set) written
    // color      func(a ...interface{}) string // default color is white
    // enabled  bool          // active holds the state of the spinner
    // Prefix     string                        // Prefix is the text preppended to the indicator
    // Suffix     string                        // Suffix is the text appended to the indicator
}

// New provides a pointer to an instance of Spinner with the supplied options
func New(t int, d time.Duration) *Spinner {
    strings := CharSets[t]
    k := len(strings)
    s := Spinner{
        Interval:   d,
        frames:     ring.New(k),
        lock:       &sync.RWMutex{},
        Writer:     os.Stderr,
        colorLevel: Color256,
        FinalMSG:   "Done!\n",
        stop:       make(chan bool),
    }
    for i := 0; i < k; i++ {
        s.frames.Value = strings[i]
        s.frames = s.frames.Next()
    }
    return &s
}

// IsActive will return whether or not the spinner is currently active
func (s *Spinner) IsActive() bool {
    return s.active
}

func (s *Spinner) getFrame() string {
    s.frames = s.frames.Next()
    return s.frames.Value.(string) + "\x1b[1D"
}

// Start will start the indicator
func (s *Spinner) Start() {
    s.lock.Lock()
    if s.active {
        s.lock.Unlock()
        return
    }
    s.active = true
    s.lock.Unlock()
    ticker := time.NewTicker(s.Interval)
    go func() {
        for {
            select {
            case <-s.stop:
                return
            case <-ticker.C:
                _, _ = fmt.Fprintf(s.Writer, s.getFrame())
            }
        }
    }()
}

// Stop stops the indicator
func (s *Spinner) Stop() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.active {
        s.active = false
        if s.FinalMSG != "" {
            _, _ = fmt.Fprintf(s.Writer, s.FinalMSG)
        }
        s.stop <- true
    }
}
