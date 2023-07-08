package chaos

import (
	"math/rand"
	"time"
)

const jitterPercent = 80

var derivativeMap = make(map[string]float64)
var tendencyMap = make(map[string]bool)
var lastTendency = true

func Sleep(ms float64, key string) (appliedSleep float64) {
	derivative, ok := derivativeMap[key]
	if !ok {
		derivative = 0
		tendencyMap[key] = lastTendency
		lastTendency = !lastTendency
	}

	positiveDerivative := GetDerivativeWithTendency(key)
	jitter := GenerateJitter(ms, jitterPercent)
	//fmt.Printf("jitter: %f\n", jitter)

	if positiveDerivative {
		derivative = (derivative + jitter) / 2.0
	} else {
		derivative = (derivative - jitter) / 2.0
	}

	//fmt.Printf("derivative: %f\n", derivative)

	derivativeMap[key] = derivative

	msDerivative := ms + derivative

	time.Sleep(time.Millisecond * time.Duration(msDerivative))
	return msDerivative
}

func GetDerivativeWithTendency(key string) bool {
	derivative := RandomBool()
	if derivative != tendencyMap[key] {
		return RandomBool()
	}

	return derivative
}

func GenerateJitter(ms float64, percent float64) float64 {
	normalizedPercent := percent / 100.0
	return RandomFloat64(0, ms*normalizedPercent)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomBool() bool {
	return rand.Int()%2 == 0
}
