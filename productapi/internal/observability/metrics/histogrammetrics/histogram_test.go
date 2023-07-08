package histogrammetrics

import (
	"fmt"
	"testing"
)

func TestGetDefaultBucket(t *testing.T) {
	// Act
	buckets := GetDefaultBucket()

	// Assert
	fmt.Println(buckets)
}
