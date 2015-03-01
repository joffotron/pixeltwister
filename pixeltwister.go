package main

import (
    "brutalbits.com/pixeltwister/server"
    "fmt"
)

func main() {
    fmt.Println("Pixeltwister started!")
    go server.Start()
    select {} //Sleep
}
