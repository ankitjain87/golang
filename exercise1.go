
// Exercise: Loops and Functions
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(x/2)
    for i := 0; i < 10; i++ {
        if y := z - (math.Pow(z, 2))/ 2 * z; y == z {
            return z
        } else {
            z = y
        }
        
        fmt.Println(z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(4))
}
