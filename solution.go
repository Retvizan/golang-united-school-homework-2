package square

import (
	"math"
)

const pi = math.Pi

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myType int
const SidesTriangle myType = 3
const SidesSquare myType = 4
const SidesCircle myType = 0


func CalcSquare(sideLen float64, sidesNum myType) float64 {
	switch sidesNum {
	case 0:
		return pi * math.Pow(sideLen, 2)
	case 3:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
