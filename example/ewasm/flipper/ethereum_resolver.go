package main

import (
	"fmt"

	"github.com/xuperchain/xvm/exec"
)

var (

	// ???
	ethResolver = exec.MapResolver{
		"ethereum.storageStore": func(ctx exec.Context, a int, b int) {
			fmt.Println("storageStore")
		},
		"ethereum.getCallValue": func(ctx exec.Context) {
			fmt.Println("getCallValue")

		},
		"ethereum.getCodeSize": func(ctx exec.Context) uint32 {
			fmt.Println("getCodeSize")

			return 0
		},
		"ethereum.codeCopy": func(ctx exec.Context, a, b, c uint32) {
			fmt.Println("codeCopy")
		},
		"ethereum.revert": func(ctx exec.Context, a, b uint32) {
			fmt.Println("revert")
		},

		"ethereum.finish": func(ctx exec.Context, a, b uint32) {
			fmt.Println("revert")
		},
		"seal0.seal_get_storage": func(ctx exec.Context, a, b, c uint32) uint32 {
			fmt.Println("seal_get_storage")
			return 0
		},
		"seal0.seal_set_storage": func(ctx exec.Context, a, b, c uint32) {
			fmt.Println("seal_get_storage")
		},
		// TODO return value
		"seal0.seal_input": func(ctx exec.Context, a, b uint32) uint32 {
			fmt.Println(ctx.Memory()[a : a+10])
			fmt.Println(ctx.Memory()[b : b+10])

			fmt.Println("seal input")
			return 128
		},

		"seal0.seal_return": func(ctx exec.Context, a, b, c uint32) uint32 {
			fmt.Println("seal_return")
			return 0
		},
		"seal0.seal_value_transferred": func(ctx exec.Context, a, b uint32) {
			fmt.Println("seal_value_transferred")
		},
	}
)

func newEthereumResolver() exec.Resolver {
	return ethResolver

}
