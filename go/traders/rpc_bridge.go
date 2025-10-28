package traders

/*
#cgo LDFLAGS: -L../../rust/target/release -lai_trading_core
#include "../../rust/target/release/ai_trading_core.h"
*/
import "C"
import "fmt"

func InitBridge() {
	fmt.Println("Connecting to Rust trading core...")
	C.init_core()
}
