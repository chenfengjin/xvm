package exec

// #include "wasm-rt.h"
// extern void init_go_trap();
import "C"

import (
	"fmt"
)

//export go_xvm_trap
func go_xvm_trap(code C.wasm_rt_trap_t) {
	switch code {
	case C.WASM_RT_TRAP_OOB:
		panic(TrapOOB)
	case C.WASM_RT_TRAP_INT_OVERFLOW:
		panic(TrapIntOverflow)
	case C.WASM_RT_TRAP_DIV_BY_ZERO:
		panic(TrapDivByZero)
	case C.WASM_RT_TRAP_INVALID_CONVERSION:
		panic(TrapInvalidConvert)
	case C.WASM_RT_TRAP_UNREACHABLE:
		panic(TrapUnreachable)
	case C.WASM_RT_TRAP_CALL_INDIRECT:
		panic(TrapInvalidIndirectCall)
	case C.WASM_RT_TRAP_EXHAUSTION:
		panic(TrapCallStackExhaustion)
	case C.WASM_RT_TRAP_GAS_EXHAUSTION:
		panic(TrapGasExhaustion)
	case C.WASM_RT_TRAP_INVALID_ARGUMENT:
		panic(TrapInvalidArgument)
	default:
		panic(NewTrap(fmt.Sprintf("trap with code:%d", code)))
	}
}

//export xvm_raise
func xvm_raise(msgptr *C.char) {
	msg := C.GoString(msgptr)
	Throw(NewTrap(msg))
}

func init() {
	C.init_go_trap()
}
