package main

import (  
    "fmt"
)

func main(){  
    for combo := range GenerateCombinationsString([]string{"1","2","3","4"}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsInt([]int{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsInt8([]int8{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsInt16([]int16{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsInt32([]int32{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsInt64([]int64{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsUint([]uint{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsUint8([]uint8{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsUint16([]uint16{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsUint32([]uint32{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsUint64([]uint64{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsFloat32([]float32{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsFloat64([]float64{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsComplex64([]complex64{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
    for combo := range GenerateCombinationsComplex128([]complex128{1,2,3,4}, 2) {
        fmt.Println(combo)
    }
}

func GenerateCombinationsString(data []string, length int) <-chan []string {  
    c := make(chan []string)
    go func(c chan []string) {
        defer close(c)
        combosString(c, []string{}, data, length)
    }(c)
    return c
}
func combosString(c chan []string, combo []string, data []string, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []string
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]string, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosString(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsInt(data []int, length int) <-chan []int {  
    c := make(chan []int)
    go func(c chan []int) {
        defer close(c)
        combosInt(c, []int{}, data, length)
    }(c)
    return c
}
func combosInt(c chan []int, combo []int, data []int, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []int
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]int, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosInt(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsInt8(data []int8, length int) <-chan []int8 {  
    c := make(chan []int8)
    go func(c chan []int8) {
        defer close(c)
        combosInt8(c, []int8{}, data, length)
    }(c)
    return c
}
func combosInt8(c chan []int8, combo []int8, data []int8, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []int8
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]int8, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosInt8(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsInt16(data []int16, length int) <-chan []int16 {  
    c := make(chan []int16)
    go func(c chan []int16) {
        defer close(c)
        combosInt16(c, []int16{}, data, length)
    }(c)
    return c
}
func combosInt16(c chan []int16, combo []int16, data []int16, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []int16
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]int16, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosInt16(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsInt32(data []int32, length int) <-chan []int32 {  
    c := make(chan []int32)
    go func(c chan []int32) {
        defer close(c)
        combosInt32(c, []int32{}, data, length)
    }(c)
    return c
}
func combosInt32(c chan []int32, combo []int32, data []int32, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []int32
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]int32, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosInt32(c, newCombo, data, length-1)
    }
}


func GenerateCombinationsInt64(data []int64, length int) <-chan []int64 {  
    c := make(chan []int64)
    go func(c chan []int64) {
        defer close(c)
        combosInt64(c, []int64{}, data, length)
    }(c)
    return c
}
func combosInt64(c chan []int64, combo []int64, data []int64, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []int64
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]int64, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosInt64(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsUint(data []uint, length int) <-chan []uint {  
    c := make(chan []uint)
    go func(c chan []uint) {
        defer close(c)
        combosUint(c, []uint{}, data, length)
    }(c)
    return c
}
func combosUint(c chan []uint, combo []uint, data []uint, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []uint
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]uint, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosUint(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsUint8(data []uint8, length int) <-chan []uint8 {  
    c := make(chan []uint8)
    go func(c chan []uint8) {
        defer close(c)
        combosUint8(c, []uint8{}, data, length)
    }(c)
    return c
}
func combosUint8(c chan []uint8, combo []uint8, data []uint8, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []uint8
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]uint8, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosUint8(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsUint16(data []uint16, length int) <-chan []uint16 {  
    c := make(chan []uint16)
    go func(c chan []uint16) {
        defer close(c)
        combosUint16(c, []uint16{}, data, length)
    }(c)
    return c
}
func combosUint16(c chan []uint16, combo []uint16, data []uint16, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []uint16
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]uint16, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosUint16(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsUint32(data []uint32, length int) <-chan []uint32 {  
    c := make(chan []uint32)
    go func(c chan []uint32) {
        defer close(c)
        combosUint32(c, []uint32{}, data, length)
    }(c)
    return c
}
func combosUint32(c chan []uint32, combo []uint32, data []uint32, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []uint32
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]uint32, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosUint32(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsUint64(data []uint64, length int) <-chan []uint64 {  
    c := make(chan []uint64)
    go func(c chan []uint64) {
        defer close(c)
        combosUint64(c, []uint64{}, data, length)
    }(c)
    return c
}
func combosUint64(c chan []uint64, combo []uint64, data []uint64, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []uint64
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]uint64, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosUint64(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsFloat32(data []float32, length int) <-chan []float32 {  
    c := make(chan []float32)
    go func(c chan []float32) {
        defer close(c)
        combosFloat32(c, []float32{}, data, length)
    }(c)
    return c
}
func combosFloat32(c chan []float32, combo []float32, data []float32, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []float32
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]float32, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosFloat32(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsFloat64(data []float64, length int) <-chan []float64 {  
    c := make(chan []float64)
    go func(c chan []float64) {
        defer close(c)
        combosFloat64(c, []float64{}, data, length)
    }(c)
    return c
}
func combosFloat64(c chan []float64, combo []float64, data []float64, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []float64
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]float64, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosFloat64(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsComplex64(data []complex64, length int) <-chan []complex64 {  
    c := make(chan []complex64)
    go func(c chan []complex64) {
        defer close(c)
        combosComplex64(c, []complex64{}, data, length)
    }(c)
    return c
}
func combosComplex64(c chan []complex64, combo []complex64, data []complex64, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []complex64
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]complex64, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosComplex64(c, newCombo, data, length-1)
    }
}

func GenerateCombinationsComplex128(data []complex128, length int) <-chan []complex128 {  
    c := make(chan []complex128)
    go func(c chan []complex128) {
        defer close(c)
        combosComplex128(c, []complex128{}, data, length)
    }(c)
    return c
}
func combosComplex128(c chan []complex128, combo []complex128, data []complex128, length int) {  
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }
    var newCombo []complex128
    for _, ch := range data {
        newCombo = append(combo, ch)
        //remove this conditional to return all sets of length <=k
        if(length == 1){
            output := make([]complex128, len(newCombo))
            copy(output, newCombo)
            c <- output
        }
        combosComplex128(c, newCombo, data, length-1)
    }
}