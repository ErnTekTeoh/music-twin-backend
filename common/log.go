package common

import (
	"context"
	"fmt"
)

const (
	Debug int = iota
	Error
	Info
	Data
)

func log(ctx context.Context, level int, msg string) {
	fmt.Println(msg)
}

func LogWithError(ctx context.Context, msg string) {
	log(ctx, Error, msg)
}

func LogWithData(ctx context.Context, msg string) {
	log(ctx, Data, msg)
}

func LogWithDebug(ctx context.Context, msg string) {
	log(ctx, Debug, msg)
}

func LogWithErrorRequiringAction(ctx context.Context, msg string) {
	log(ctx, Error, "‼️ "+msg)
}
