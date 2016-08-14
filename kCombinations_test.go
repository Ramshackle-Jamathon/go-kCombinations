package kCombinations_test

import (
	"testing"
	"fmt"

	"github.com/Ramshackle-Jamathon/kCombinations"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    for combo := range kCombinations.GenerateCombinationsString([]string{"1","2","3","4"}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsInt([]int{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsInt8([]int8{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsInt16([]int16{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsInt32([]int32{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsInt64([]int64{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsUint([]uint{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsUint8([]uint8{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsUint16([]uint16{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsUint32([]uint32{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsUint64([]uint64{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsFloat32([]float32{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsFloat64([]float64{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsComplex64([]complex64{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
    for combo := range kCombinations.GenerateCombinationsComplex128([]complex128{1,2,3,4}, 2) {
        fmt.Println(combo)
        assert.NotNil(t, combo)
    }
}