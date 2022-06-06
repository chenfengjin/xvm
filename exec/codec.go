package exec

import (
	"encoding/binary"
	"fmt"
)

var (
	trapNilMemory = NewTrap("code has no memory")
)

// TrapInvalidAddress is the trap raised when encounter an invalid address
type TrapInvalidAddress uint32

// Reason implements Trap interface
func (t TrapInvalidAddress) Reason() string {
	return fmt.Sprintf("invalid address:0x%x", uint32(t))
}

// Codec helps encoding and decoding data between wasm code and go code
type Codec struct {
	mem []byte
}

var (
	// TrapOOB is raised when memory access out of bound
	TrapOOB = NewTrap("memory access out of bound")
	// TrapIntOverflow is raised when math overflow
	TrapIntOverflow = NewTrap("integer overflow on divide or truncation")
	// TrapDivByZero is raised when divide by zero
	TrapDivByZero = NewTrap("integer divide by zero")
	// TrapInvalidConvert is raised when convert from NaN to integer
	TrapInvalidConvert = NewTrap("conversion from NaN to integer")
	// TrapUnreachable is raised when unreachable instruction executed
	TrapUnreachable = NewTrap("unreachable instruction executed")
	// TrapInvalidIndirectCall is raised when run invalid call_indirect instruction
	TrapInvalidIndirectCall = NewTrap("invalid call_indirect")
	// TrapCallStackExhaustion is raised when call stack exhausted
	TrapCallStackExhaustion = NewTrap("call stack exhausted")
	// TrapGasExhaustion is raised when runnning out of gas limit
	TrapGasExhaustion = NewTrap("run out of gas limit")
	// TrapInvalidArgument is raised when running function with invalid argument
	TrapInvalidArgument = NewTrap("invalid function argument")
)

// Trap 用于表示虚拟机运行过程中的错误，中断虚拟机的运行
type Trap interface {
	Reason() string
}

// TrapError 用于包装一个Trap到Error
type TrapError struct {
	Trap Trap
}

func (t *TrapError) Error() string {
	return fmt.Sprintf("trap error:%s", t.Trap.Reason())
}

// Throw 用于抛出一个Trap
func Throw(trap Trap) {
	panic(trap)
}

// ThrowError throws an error as an trap
func ThrowError(err error) {
	Throw(NewTrap(err.Error()))
}

// ThrowMessage throws a string message as an trap
func ThrowMessage(msg string) {
	Throw(NewTrap(msg))
}

// CaptureTrap 用于捕获潜在的Trap，如果是其他panic则不会捕获
func CaptureTrap(err *error) {
	ret := recover()
	if ret == nil {
		return
	}
	trap, ok := ret.(Trap)
	if ok {
		*err = &TrapError{
			Trap: trap,
		}
		return
	}
	panic(ret)
}

type stringTrap struct {
	reason string
}

func (s *stringTrap) Reason() string {
	return s.reason
}

// NewTrap returns a trap with the given reason
func NewTrap(reason string) Trap {
	return &stringTrap{
		reason,
	}
}

// TrapSymbolNotFound is raised when resolving symbol failed
type TrapSymbolNotFound struct {
	Module string
	Name   string
}

// Reason implements Trap interface
func (s *TrapSymbolNotFound) Reason() string {
	return fmt.Sprintf("%s.%s can't be resolved", s.Module, s.Name)
}

// TrapFuncSignatureNotMatch is raised when calling function signature is not matched
type TrapFuncSignatureNotMatch struct {
	Module string
	Name   string
}

// Reason implements Trap interface
func (s *TrapFuncSignatureNotMatch) Reason() string {
	return fmt.Sprintf("%s.%s not match with host signature", s.Module, s.Name)
}

const (
	// MaxGasLimit is the maximum gas limit
	MaxGasLimit = 0xFFFFFFFF
)

type ErrFuncNotFound struct {
	Name string
}

func (e *ErrFuncNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Name)
}
func (e *ErrFuncNotFound) Is(err error) bool {
	err1, ok := err.(*ErrFuncNotFound)
	return ok && err1.Name == e.Name
}

// ContextConfig configures an execution context
type ContextConfig struct {
	GasLimit int64
}

// DefaultContextConfig returns the default configuration of ContextConfig
func DefaultContextConfig() *ContextConfig {
	return &ContextConfig{
		GasLimit: MaxGasLimit,
	}
}

// Context hold the context data when running a wasm instance
type Context interface {
	Exec(name string, param []int64) (ret int64, err error)
	GasUsed() int64
	ResetGasUsed()
	Memory() []byte
	StaticTop() uint32
	SetUserData(key string, value interface{})
	GetUserData(key string) interface{}
	Release()
}

type Code interface {
	NewContext(cfg *ContextConfig) (ictx Context, err error)
	Release()
}

// NewCodec instances a Codec, if memory of ctx is nil, trapNilMemory will be raised
func NewCodec(ctx Context) Codec {
	mem := ctx.Memory()
	if mem == nil {
		Throw(trapNilMemory)
	}

	return Codec{
		mem: mem,
	}
}

// Bytes returns memory region starting from addr, limiting by length
func (c Codec) Bytes(addr, length uint32) []byte {
	if addr+length >= uint32(len(c.mem)) {
		Throw(TrapInvalidAddress(addr + length))
	}
	return c.mem[addr : addr+length]
}

// Uint32 decodes memory[addr:addr+4] to uint32
func (c Codec) Uint32(addr uint32) uint32 {
	buf := c.Bytes(addr, 4)
	return binary.LittleEndian.Uint32(buf)
}

// SetUint32 set val to memory[addr:addr+4]
func (c Codec) SetUint32(addr uint32, val uint32) {
	buf := c.Bytes(addr, 4)
	binary.LittleEndian.PutUint32(buf, val)
}

// Uint64 decodes memory[addr:addr+8] to uint64
func (c Codec) Uint64(addr uint32) uint64 {
	buf := c.Bytes(addr, 8)
	return binary.LittleEndian.Uint64(buf)
}

// GoBytes decodes Go []byte start from sp
func (c Codec) GoBytes(sp uint32) []byte {
	addr := c.Uint64(sp)
	length := c.Uint64(sp + 8)
	return c.Bytes(uint32(addr), uint32(length))
}

// GoString decodes Go string start from sp
func (c Codec) GoString(sp uint32) string {
	return string(c.GoBytes(sp))
}

// String decodes memory[addr:addr+length] to string
func (c Codec) String(addr, length uint32) string {
	return string(c.Bytes(addr, length))
}

// CString decodes a '\x00' terminated c style string
func (c Codec) CString(addr uint32) string {
	if addr == 0 {
		Throw(TrapInvalidAddress(addr))
	}
	mem := c.mem
	var i = int(addr)
	for ; i < len(mem) && mem[i] != '\x00'; i++ {
	}
	return string(mem[addr:i])
}
