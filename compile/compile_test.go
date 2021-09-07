package compile

import (
	"testing"
)

func BenchmarkCompileNativeLibrary(b *testing.B) {
	b.Run("OPT0", func(b *testing.B) {
		cfg := Config{
			Wasm2cPath:   "/Users/chenfengjin/baidu/xuperchain/output/bin/wasm2c",
			Wat2wasmPath: "",
			OptLevel:     0,
		}
		for i := 0; i < b.N; i++ {
			err := CompileNativeLibrary(&cfg, "testdata/counter.so", "testdata/counter.wasm")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("OPT2", func(b *testing.B) {
		cfg := Config{
			Wasm2cPath:   "/Users/chenfengjin/baidu/xuperchain/output/bin/wasm2c",
			Wat2wasmPath: "",
			OptLevel:     2,
		}
		for i := 0; i < b.N; i++ {
			err := CompileNativeLibrary(&cfg, "testdata/counter.so", "testdata/counter.wasm")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkWabt(b *testing.B) {
	b.Run("Default", func(b *testing.B) {
		cfg := &Config{
			Wasm2cPath:   "/Users/chenfengjin/baidu/xuperchain/output/bin/wasm2c",
			Wat2wasmPath: "",
			OptLevel:     2,
		}
		for i := 0; i < b.N; i++ {
			err := CompileCSource(cfg, "testdata/counter.c", "testdata/counter.wasm")
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("New", func(b *testing.B) {
		cfg := &Config{
			Wasm2cPath:   "/Users/chenfengjin/baidu/wabt2/cmake-build-release/wasm2c",
			Wat2wasmPath: "",
			OptLevel:     2,
		}
		for i := 0; i < b.N; i++ {
			err := CompileCSource(cfg, "testdata/counter.c", "testdata/counter.wasm")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
