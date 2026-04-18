package main
import (
    "fmt"
    "github.com/Andrii-Brazhaniuk/lab1-tooling/internal"
)

func main() {
    sum := internal.Add(40, 2)
    fmt.Printf("Sum: %d\n", sum)

    div, err := internal.Divide(10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Division: %d\n", div)
    }
}
