package main

import "testing"

func TestSum(t *testing.T) {

    //test for sum
    returnSum := add(5,5)

    if returnSum != 10 {
        t.Error("Sum failed, expect 10")
    }
}