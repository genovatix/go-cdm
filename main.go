package main

import (
	"context"
	"github.com/libp2p/go-libp2p"
)

func main() {
	ctx := context.Background()
	h, err := libp2p.NewWithoutDefaults()
}
