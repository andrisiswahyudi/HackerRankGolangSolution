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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    // Write your code here
    var min, max = arr[0], arr[0]
    var sum int64
    for _, i:= range arr {
        if i < min {
            min = i
        }
        if i > max {
            max = i
        }
        sum += int64(i)
    }   
    fmt.Println(sum - int64(max), sum - int64(min))

}