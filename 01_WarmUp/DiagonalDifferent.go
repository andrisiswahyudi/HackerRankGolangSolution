package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
    var d1, d2 int32 = 0, 0
    n := len(arr)
    var diff float64
    
    for i, _ := range arr {
        d1 += arr[i][i]
        d2 += arr[i][n - i -1]
        diff = math.Abs(float64(d1 - d2))
    }
    return int32(diff)
}
