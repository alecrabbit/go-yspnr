package main

import (
    "fmt"
    "regexp"
    "time"
)

const exp = `\x1b[[][^A-Za-z]*[A-Za-z]`

func main() {
    re := regexp.MustCompile(`a(x*)b`)
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

    r := regexp.MustCompile(exp)
    str := "\x1b[1mlolo\x1b[0m\x1b[4D"
    fmt.Print(str)
    time.Sleep(1*time.Second)
    fmt.Println("")
    s := r.ReplaceAllString(str, "")
    fmt.Print(s)
    time.Sleep(1*time.Second)
    fmt.Println("")
    fmt.Println(len(str))
    fmt.Println(len(s))
}
