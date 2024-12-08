```go
package main

import (

        "fmt"

)

func main() {

        var a [10]int
        for i := 0; i < 10; i++ {
                a[i] = i
        }
        fmt.Println(a)

        // Create a deep copy of the array
        b := make([]int, len(a))
        copy(b, a[:])

        b[0] = 100
        fmt.Println(a)
        fmt.Println(b)

}```