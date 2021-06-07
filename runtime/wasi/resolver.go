package wasi

import (
	"fmt"
	"github.com/xuperchain/xvm/exec"
)

var resolver = exec.MapResolver(map[string]interface{}{
	"wasi_unstable.fd_prestat_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 8
	},
	"wasi_unstable.fd_fdstat_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 8
	},
	"wasi_unstable.fd_prestat_dir_name": func(ctx exec.Context, x, y, z uint32) uint32 {
		return 8
	},
	"wasi_unstable.fd_close": func(ctx exec.Context, x uint32) uint32 {
		return 8
	},
	"wasi_unstable.fd_seek": func(ctx exec.Context, x, y, z, w uint32) uint32 {
		return 8
	},
	"wasi_unstable.fd_write": func(ctx exec.Context, x, y, z, w uint32) uint32 {
		return 8
	},
	"wasi_unstable.environ_sizes_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_unstable.environ_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_unstable.args_sizes_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_unstable.args_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_unstable.proc_exit": func(ctx exec.Context, x uint32) uint32 {
		exec.Throw(exec.NewTrap("exit"))
		return 0
	},
	"wasi_snapshot_preview1.fd_prestat_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 8
	},
	"wasi_snapshot_preview1.fd_fdstat_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 8
	},
	"wasi_snapshot_preview1.fd_prestat_dir_name": func(ctx exec.Context, x, y, z uint32) uint32 {
		return 8
	},
	"wasi_snapshot_preview1.fd_close": func(ctx exec.Context, x uint32) uint32 {
		return 8
	},
	"wasi_snapshot_preview1.fd_seek": func(ctx exec.Context, x, y, z, w uint32) uint32 {
		return 8
	},
	"wasi_snapshot_preview1.fd_write": func(ctx exec.Context, x, y, z, w uint32) uint32 {
		fmt.Println("hahahhahahahah")
		return 8
	},
	"wasi_snapshot_preview1.environ_sizes_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_snapshot_preview1.environ_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_snapshot_preview1.args_sizes_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_snapshot_preview1.args_get": func(ctx exec.Context, x, y uint32) uint32 {
		return 0
	},
	"wasi_snapshot_preview1.proc_exit": func(ctx exec.Context, x uint32) uint32 {
		exec.Throw(exec.NewTrap("exit"))
		return 0
	},
})

// NewResolver return exec.Resolver which resolves symbols needed by wasi environment
func NewResolver() exec.Resolver {
	return resolver
}
