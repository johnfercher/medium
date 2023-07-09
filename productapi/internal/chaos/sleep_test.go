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
