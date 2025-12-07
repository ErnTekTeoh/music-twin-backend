package module

import (
	"context"
	"fmt"
	"testing"
)

func TestSearchTrack(t *testing.T) {
	res, err := SearchTrack(context.Background(), "drip", "babymonster")
	fmt.Println(res)
	fmt.Println(err)
}
