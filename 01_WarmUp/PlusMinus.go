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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var positif, negatif, nol int
    l := len(arr)
    
    
    for _, i := range arr {
        if i > 0 {
            positif++
        } else if i < 0 {
            negatif++
        } else {
            nol++
        }
        
    }
    a := float64(positif)/float64(l)
    b := float64(negatif)/float64(l)
    c := float64(nol)/float64(l)
    
    fmt.Printf("%.6f\n", a)
    fmt.Printf("%.6f\n", b)
    fmt.Printf("%.6f\n", c)
}