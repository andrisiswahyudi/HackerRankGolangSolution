package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
    // Write your code here
    var i int32
    for i = 0; i < n; i++ {
        var j int32
        for j = 0; j < n - i -1; j++ {
            fmt.Print(" ")
        }
        //var k int32
        for j = 0; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()  
    }

}