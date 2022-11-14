package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    if os.Getenv("DEBUG_ENABLE") == "True" {
        debugOn := true
    }

    f, err := os.Open("cool-log-file.txt")
    check(err)

    scanner := bufio.NewScanner(f)

    line := 1
    connectivity_problems := 0
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), "connection") {
            if debugOn {
                fmt.Println(scanner.Text())
            }
            fmt.Println("connectivity, i KNEW IT %v", line)
            connectivity_problems++
        }

        line++
    }

    fmt.Printf("I found %d connectivity problems\n", connectivity_problems)

}
