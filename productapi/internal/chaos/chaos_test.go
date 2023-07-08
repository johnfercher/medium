package chaos

import (
	"fmt"
	"testing"
)

func TestGeneratePositiveJitter(t *testing.T) {
	for i := 0; i < 10; i++ {
		value := GenerateJitter(100, 50)
		fmt.Println(value)
	}
}

func TestSleep(t *testing.T) {
	fmt.Println("endpoint1")

	for i := 0; i < 10; i++ {
		value := Sleep(100, "endpoint1")
		fmt.Println(value)
	}

	fmt.Println("endpoint2")

	for i := 0; i < 10; i++ {
		value := Sleep(100, "endpoint2")
		fmt.Println(value)
	}
}
