package chaos

const jitterPercent = 80

var lastTendency = true

func GenerateJitter(ms float64, percent float64) float64 {
	normalizedPercent := percent / 100.0
	return RandomFloat64(0, ms*normalizedPercent)
}

func BuildLatencyTendency() bool {
	lastTendency = !lastTendency
	return lastTendency
}
