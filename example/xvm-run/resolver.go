package main

import (
	"fmt"

	"github.com/xuperchain/xvm/exec"
)

var resolver = exec.MapResolver(map[string]interface{}{
	"go.main.puts": func(ctx exec.Context, sp uint32) uint32 {
		codec := exec.NewCodec(ctx)
		str := codec.GoString(sp + 8)
		fmt.Print(str)
		return 0
	},
	"env._print": func(ctx exec.Context, ptr uint32) uint32 {
		codec := exec.NewCodec(ctx)
		str := codec.CString(ptr)
		fmt.Print(str)
		return 0
	},
	"env._println": func(ctx exec.Context, ptr uint32) uint32 {
		codec := exec.NewCodec(ctx)
		str := codec.CString(ptr)
		fmt.Println(str)
		return 0
	},
	//return is neccessary for xvm
	//but for wasm ok??
	"__wbindgen_placeholder__.__wbindgen_describe": func(ctx exec.Context) uint32 {
		return 0
	},
	"__wbindgen_placeholder__.__wbg_alert_531f0294104c16c5": func(ctx exec.Context, a uint32, b uint32) uint32 {
		return 0
	},
	"__wbindgen_externref_xform__.__wbindgen_externref_table_grow": func(ctx exec.Context, b uint32) uint32 {
		return 0
	},
	"__wbindgen_externref_xform__.__wbindgen_externref_table_set_null": func(ctx exec.Context, b uint32) {
	},
})

//(import "wasi_snapshot_preview1" "fd_write" (func $wasi::lib_generated::wasi_snapshot_preview1::fd_write::ha0aef7cef0a152b0 (type 12)))
//(import "wasi_snapshot_preview1" "environ_sizes_get" (func $__wasi_environ_sizes_get (type 6)))
//(import "wasi_snapshot_preview1" "proc_exit" (func $__wasi_proc_exit (type 2)))
//(import "wasi_snapshot_preview1" "environ_get" (func $__wasi_environ_get (type 6)))
