package exec

// #include "xvm.h"
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/xuperchain/xvm/pointer"
)

//export xvm_resolve_func
func xvm_resolve_func(env unsafe.Pointer, module, name *C.char) C.wasm_rt_func_handle_t {
	r := pointer.Restore(uintptr(env)).(*resolverBridge)
	moduleStr, nameStr := C.GoString(module), C.GoString(name)
	key := moduleStr + ":" + nameStr

	idx := r.funcmap[key]
	if idx != 0 {
		return C.wasm_rt_func_handle_t(uintptr(idx))
	}
	if r.resolver == nil {
		Throw(&TrapSymbolNotFound{
			Module: moduleStr,
			Name:   nameStr,
		})
	}

	f, ok := r.resolver.ResolveFunc(moduleStr, nameStr)
	if !ok {
		Throw(&TrapSymbolNotFound{
			Module: moduleStr,
			Name:   nameStr,
		})
	}
	r.funcs = append(r.funcs, importFunc{
		module: moduleStr,
		name:   nameStr,
		body:   f,
	})
	idx = len(r.funcs) - 1
	r.funcmap[key] = idx
	return C.wasm_rt_func_handle_t(uintptr(idx))
}

//export xvm_resolve_global
func xvm_resolve_global(env unsafe.Pointer, module, name *C.char) C.int64_t {
	r := pointer.Restore(uintptr(env)).(*resolverBridge)
	moduleStr, nameStr := C.GoString(module), C.GoString(name)
	value, ok := r.resolver.ResolveGlobal(moduleStr, nameStr)
	if !ok {
		Throw(&TrapSymbolNotFound{
			Module: moduleStr,
			Name:   nameStr,
		})
	}
	return C.int64_t(value)
}

//export xvm_call_func
func xvm_call_func(env unsafe.Pointer, handle C.wasm_rt_func_handle_t,
	ctxptr *C.xvm_context_t, params *C.uint32_t, paramLen C.uint32_t) C.uint32_t {
	r := pointer.Restore(uintptr(env)).(*resolverBridge)
	idx := int(uintptr(handle))
	if idx <= 0 || idx >= len(r.funcs) {
		Throw(NewTrap(fmt.Sprintf("bad func idx %d", idx)))
	}
	f := r.funcs[idx]
	args := make([]uint32, paramLen)
	for i := range args {
		args[i] = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(params)) + uintptr(i*4)))
	}
	// TODO: 因为context字段是Context的第一个字段，可以强转，希望后续go的内存布局不会变化
	// FIXME: cgo应该不会有问题，如果有问题可以使用pointer package来转换
	ctx := (*aotContext)(unsafe.Pointer(ctxptr))
	ret, ok := applyFuncCall(ctx, f.body, args)
	if !ok {
		Throw(&TrapFuncSignatureNotMatch{
			Module: f.module,
			Name:   f.name,
		})
	}
	return C.uint32_t(ret)
}
