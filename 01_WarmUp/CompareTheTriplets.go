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
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
    // Write your code here
    var scoreAlice int32 = 0
    var scoreBob int32 = 0
    
    
    for i, _ := range a {
        if a[i] > b[i] {
            scoreAlice += 1
        } else if a[i] < b[i] {
            scoreBob += 1
        }
    }
    return [] int32 {scoreAlice, scoreBob}

}

