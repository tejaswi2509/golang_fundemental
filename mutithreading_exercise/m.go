package main

import (
    "fmt"
    "sync"
 "time"
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
    sliceLength := len(slice)

    var wg sync.WaitGroup
    
    wg.Add(sliceLength)

    fmt.Println("Running for loop...")

    for i := 0; i < sliceLength; i++ {
        fmt.Println(i)
        go func(i int) {
            
            defer wg.Done()
            val := slice[i]
	    
            fmt.Printf("i: %v, val: %v\n", i, val)
	    	//time.Sleep(time.Second * 10)

        }(i)
    }
    time.Sleep(time.Second * 20)
    fmt.Println("Doing other stuff")

    wg.Wait()

    fmt.Println("Finished for loop")
}
