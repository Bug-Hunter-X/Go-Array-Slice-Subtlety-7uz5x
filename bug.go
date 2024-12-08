```go
func main() {


        var a [10]int
        for i := 0; i < 10; i++ {
                a[i] = i
        }
        fmt.Println(a)
        b := a[:]
        b[0] = 100
        fmt.Println(a)
        fmt.Println(b)

}```