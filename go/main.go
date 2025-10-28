package main

import (
	"fmt"
	"ai-trading-arena/go/traders"
)

func main() {
	fmt.Println("🧠 AI Trading Arena starting...")
	traders.InitBridge()
	fmt.Println("✅ Go ↔ Rust bridge initialized")
}
