package main

import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
r := make([][]uint8,dx)
    for x := 0; x < dx; x++ {
        r[x] = make([]uint8,dy)
        for y := 0; y < dy; y++ {
            r[x][y] = uint8(math.Pow(float64(y),2)/9+math.Pow(float64(x),float64(2))/7)
        }
    }
	return r
}

func main() {
	pic.Show(Pic)
}
