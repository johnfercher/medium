package chaos

import (
	"fmt"
	"testing"
)

func TestRandomFloat64(t *testing.T) {
	// Act
	for i := 0; i < 100; i++ {
		value := RandomFloat64(0, 100)
		fmt.Println(value)
	}
}
