package exec

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/xuperchain/xvm/compile"
)

func withCode(t testing.TB, watCode string, r Resolver, optLevel int, f func(code Code)) {
	tmpdir, err := ioutil.TempDir("", "xvm-exec-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	wasmpath := filepath.Join(tmpdir, "wasm.wasm")
	libpath := filepath.Join(tmpdir, "wasm.so")
	cfg := &compile.Config{
		Wasm2cPath:   "../compile/wabt/build/wasm2c",
		Wat2wasmPath: "../compile/wabt/build/wat2wasm",
		OptLevel:     optLevel,
	}

	err = compile.CompileWatSource(cfg, wasmpath, watCode)
	if err != nil {
		t.Fatal(err)
	}

	err = compile.CompileNativeLibrary(cfg, libpath, wasmpath)
	if err != nil {
		t.Fatal(err)
	}
	code, err := NewAOTCode(libpath, r)
	if err != nil {
		t.Fatal(err)
	}
	f(code)
	code.Release()
}

func BenchmarkAotCode(b *testing.B) {
	b.Run("TestOpt0", func(b *testing.B) {
		withCode(b, "testdata/sum.wat", nil, 0, func(code Code) {
			ctx, err := code.NewContext(&ContextConfig{GasLimit: MaxGasLimit})
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ret, err := ctx.Exec("run", []int64{})
				if err != nil {
					b.Fatal(err)
				}
				_ = ret
			}
		})
	})

	b.Run("TestOpt1", func(b *testing.B) {
		withCode(b, "testdata/sum.wat", nil, 1, func(code Code) {
			ctx, err := code.NewContext(&ContextConfig{GasLimit: MaxGasLimit})
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ret, err := ctx.Exec("run", []int64{})
				if err != nil {
					b.Fatal(err)
				}
				_ = ret
			}
		})
	})
	b.Run("TestOpt2", func(b *testing.B) {
		withCode(b, "testdata/sum.wat", nil, 2, func(code Code) {
			ctx, err := code.NewContext(&ContextConfig{GasLimit: MaxGasLimit})
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ret, err := ctx.Exec("run", []int64{})
				if err != nil {
					b.Fatal(err)
				}
				_ = ret
			}
		})
	})
	b.Run("TestOpt3", func(b *testing.B) {
		withCode(b, "testdata/sum.wat", nil, 3, func(code Code) {
			ctx, err := code.NewContext(&ContextConfig{GasLimit: MaxGasLimit})
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ret, err := ctx.Exec("run", []int64{})
				if err != nil {
					b.Fatal(err)
				}
				_ = ret
			}
		})
	})
}
