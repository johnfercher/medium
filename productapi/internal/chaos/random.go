package chaos

import "math/rand"

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomBool() bool {
	return rand.Int()%2 == 0
}
