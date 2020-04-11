// Autogenerated by mkruntimetypes.go, DO NOT EDIT.

package compiler

// This file contains definitions for runtime types and functions, so that the
// runtime package can be compiled independently of other packages.

import (
	"go/token"
	"go/types"
	"strconv"
)

// getRuntimeType constructs a new runtime type with the given name. The types
// constructed here must match the types in the runtime package.
func (c *compilerContext) getRuntimeType(name string) types.Type {
	if c.program != nil {
		return c.program.ImportedPackage("runtime").Type(name).Type()
	}
	if typ, ok := c.runtimeTypes[name]; ok {
		// This type was already created.
		return typ
	}

	typeName := types.NewTypeName(token.NoPos, c.runtimePkg, name, nil)
	named := types.NewNamed(typeName, nil, nil)

	// Make sure recursive types are only defined once.
	c.runtimeTypes[name] = named

	var fieldTypes []types.Type
	switch name {
	case "_defer":
		fieldTypes = []types.Type{
			types.Typ[types.Uintptr], // callback
			types.NewPointer(named),  // next
		}
	case "_string":
		fieldTypes = []types.Type{
			types.NewPointer(types.Typ[types.Byte]), // ptr
			types.Typ[types.Uintptr],                // length
		}
	case "stringIterator":
		fieldTypes = []types.Type{
			types.Typ[types.Uintptr], // byteindex
		}
	case "hashmap":
		fieldTypes = []types.Type{
			types.NewPointer(named),        // next
			types.Typ[types.UnsafePointer], // buckets
			types.Typ[types.Uintptr],       // count
			types.Typ[types.Uint8],         // keySize
			types.Typ[types.Uint8],         // valueSize
			types.Typ[types.Uint8],         // bucketBits
		}
	case "hashmapBucket":
		fieldTypes = []types.Type{
			types.NewArray(types.Typ[types.Uint8], 8), // tophash
			types.NewPointer(named),                   // next
		}
	case "hashmapIterator":
		fieldTypes = []types.Type{
			types.Typ[types.Uintptr],                            // bucketNumber
			types.NewPointer(c.getRuntimeType("hashmapBucket")), // bucket
			types.Typ[types.Uint8],                              // bucketIndex
		}
	case "channel":
		fieldTypes = []types.Type{
			types.Typ[types.Uintptr],                                 // elementSize
			types.Typ[types.Uintptr],                                 // bufSize
			types.Typ[types.Uint8],                                   // state
			types.NewPointer(c.getRuntimeType("channelBlockedList")), // blocked
			types.Typ[types.Uintptr],                                 // bufHead
			types.Typ[types.Uintptr],                                 // bufTail
			types.Typ[types.Uintptr],                                 // bufUsed
			types.Typ[types.UnsafePointer],                           // buf
		}
	case "channelBlockedList":
		taskType := types.NewNamed(types.NewTypeName(token.NoPos, c.taskPkg, "Task", nil), nil, nil)
		fieldTypes = []types.Type{
			types.NewPointer(named),                               // next
			types.NewPointer(taskType),                            // t
			types.NewPointer(c.getRuntimeType("chanSelectState")), // s
			types.NewSlice(named),                                 // allSelectOps
		}
	case "chanSelectState":
		fieldTypes = []types.Type{
			types.NewPointer(c.getRuntimeType("channel")), // ch
			types.Typ[types.UnsafePointer],                // value
		}
	case "_interface":
		fieldTypes = []types.Type{
			types.Typ[types.Uintptr],       // typecode
			types.Typ[types.UnsafePointer], // value
		}
	case "interfaceMethodInfo":
		fieldTypes = []types.Type{
			types.NewPointer(types.Typ[types.Uint8]), // signature
			types.Typ[types.Uintptr],                 // funcptr
		}
	case "typecodeID":
		fieldTypes = []types.Type{
			types.NewPointer(named),  // references
			types.Typ[types.Uintptr], // length
		}
	case "structField":
		fieldTypes = []types.Type{
			types.NewPointer(c.getRuntimeType("typecodeID")), // typecode
			types.NewPointer(types.Typ[types.Uint8]),         // name
			types.NewPointer(types.Typ[types.Uint8]),         // tag
			types.Typ[types.Bool],                            // embedded
		}
	case "typeInInterface":
		fieldTypes = []types.Type{
			types.NewPointer(c.getRuntimeType("typecodeID")),          // typecode
			types.NewPointer(c.getRuntimeType("interfaceMethodInfo")), // methodSet
		}
	case "funcValue":
		fieldTypes = []types.Type{
			types.Typ[types.UnsafePointer], // context
			types.Typ[types.Uintptr],       // id
		}
	case "funcValueWithSignature":
		fieldTypes = []types.Type{
			types.Typ[types.Uintptr],                         // funcPtr
			types.NewPointer(c.getRuntimeType("typecodeID")), // signature
		}
	default:
		panic("could not find runtime type: runtime." + name)
	}

	// Create the named struct type.
	var fields []*types.Var
	for i, t := range fieldTypes {
		// Field name doesn't matter: this type is only used to create a LLVM
		// struct type which don't have field names.
		fields = append(fields, types.NewField(token.NoPos, nil, "field"+strconv.Itoa(i), t, false))
	}
	named.SetUnderlying(types.NewStruct(fields, nil))
	return named
}

// getRuntimeFuncType constructs a new runtime function signature with the given
// name. The function signatures constructed here must match the functions in
// the runtime package.
func (c *compilerContext) getRuntimeFuncType(name string) *types.Signature {
	var params []*types.Var
	addParam := func(name string, typ types.Type) {
		params = append(params, types.NewParam(token.NoPos, c.runtimePkg, name, typ))
	}
	var results []*types.Var
	addResult := func(typ types.Type) {
		results = append(results, types.NewParam(token.NoPos, c.runtimePkg, "", typ))
	}
	switch name {
	case "_panic":
		addParam("message", types.NewInterfaceType(nil, nil))
	case "_recover":
		addResult(types.NewInterfaceType(nil, nil))
	case "nilPanic":
	case "lookupPanic":
	case "slicePanic":
	case "chanMakePanic":
	case "negativeShiftPanic":
	case "stringEqual":
		addParam("x", types.Typ[types.String])
		addParam("y", types.Typ[types.String])
		addResult(types.Typ[types.Bool])
	case "stringLess":
		addParam("x", types.Typ[types.String])
		addParam("y", types.Typ[types.String])
		addResult(types.Typ[types.Bool])
	case "stringConcat":
		addParam("x", c.getRuntimeType("_string"))
		addParam("y", c.getRuntimeType("_string"))
		addResult(c.getRuntimeType("_string"))
	case "stringFromBytes":
		addParam("x", types.NewStruct([]*types.Var{
			types.NewField(token.NoPos, nil, "ptr", types.NewPointer(types.Typ[types.Byte]), false),
			types.NewField(token.NoPos, nil, "len", types.Typ[types.Uintptr], false),
			types.NewField(token.NoPos, nil, "cap", types.Typ[types.Uintptr], false),
		}, nil))
		addResult(c.getRuntimeType("_string"))
	case "stringToBytes":
		addParam("x", c.getRuntimeType("_string"))
		addResult(types.NewStruct([]*types.Var{
			types.NewField(token.NoPos, nil, "ptr", types.NewPointer(types.Typ[types.Byte]), false),
			types.NewField(token.NoPos, nil, "len", types.Typ[types.Uintptr], false),
			types.NewField(token.NoPos, nil, "cap", types.Typ[types.Uintptr], false),
		}, nil))
	case "stringFromRunes":
		addParam("runeSlice", types.NewSlice(types.Typ[types.Rune]))
		addResult(c.getRuntimeType("_string"))
	case "stringToRunes":
		addParam("s", types.Typ[types.String])
		addResult(types.NewSlice(types.Typ[types.Rune]))
	case "stringFromUnicode":
		addParam("x", types.Typ[types.Rune])
		addResult(c.getRuntimeType("_string"))
	case "stringNext":
		addParam("s", types.Typ[types.String])
		addParam("it", types.NewPointer(c.getRuntimeType("stringIterator")))
		addResult(types.Typ[types.Bool])
		addResult(types.Typ[types.Int])
		addResult(types.Typ[types.Rune])
	case "complex64div":
		addParam("n", types.Typ[types.Complex64])
		addParam("m", types.Typ[types.Complex64])
		addResult(types.Typ[types.Complex64])
	case "complex128div":
		addParam("n", types.Typ[types.Complex128])
		addParam("m", types.Typ[types.Complex128])
		addResult(types.Typ[types.Complex128])
	case "sliceAppend":
		addParam("srcBuf", types.Typ[types.UnsafePointer])
		addParam("elemsBuf", types.Typ[types.UnsafePointer])
		addParam("srcLen", types.Typ[types.Uintptr])
		addParam("srcCap", types.Typ[types.Uintptr])
		addParam("elemsLen", types.Typ[types.Uintptr])
		addParam("elemSize", types.Typ[types.Uintptr])
		addResult(types.Typ[types.UnsafePointer])
		addResult(types.Typ[types.Uintptr])
		addResult(types.Typ[types.Uintptr])
	case "sliceCopy":
		addParam("dst", types.Typ[types.UnsafePointer])
		addParam("src", types.Typ[types.UnsafePointer])
		addParam("dstLen", types.Typ[types.Uintptr])
		addParam("srcLen", types.Typ[types.Uintptr])
		addParam("elemSize", types.Typ[types.Uintptr])
		addResult(types.Typ[types.Int])
	case "alloc":
		addParam("size", types.Typ[types.Uintptr])
		addResult(types.Typ[types.UnsafePointer])
	case "trackPointer":
		addParam("ptr", types.Typ[types.UnsafePointer])
	case "printbool":
		addParam("b", types.Typ[types.Bool])
	case "printint8":
		addParam("n", types.Typ[types.Int8])
	case "printuint8":
		addParam("n", types.Typ[types.Uint8])
	case "printint16":
		addParam("n", types.Typ[types.Int16])
	case "printuint16":
		addParam("n", types.Typ[types.Uint16])
	case "printint32":
		addParam("n", types.Typ[types.Int32])
	case "printuint32":
		addParam("n", types.Typ[types.Uint32])
	case "printint64":
		addParam("n", types.Typ[types.Int64])
	case "printuint64":
		addParam("n", types.Typ[types.Uint64])
	case "printfloat32":
		addParam("v", types.Typ[types.Float32])
	case "printfloat64":
		addParam("v", types.Typ[types.Float64])
	case "printcomplex64":
		addParam("c", types.Typ[types.Complex64])
	case "printcomplex128":
		addParam("c", types.Typ[types.Complex128])
	case "printstring":
		addParam("s", types.Typ[types.String])
	case "printspace":
	case "printnl":
	case "printptr":
		addParam("ptr", types.Typ[types.Uintptr])
	case "printmap":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
	case "printitf":
		addParam("msg", types.NewInterfaceType(nil, nil))
	case "hashmapMake":
		addParam("keySize", types.Typ[types.Uint8])
		addParam("valueSize", types.Typ[types.Uint8])
		addParam("sizeHint", types.Typ[types.Uintptr])
		addResult(types.NewPointer(c.getRuntimeType("hashmap")))
	case "hashmapLen":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addResult(types.Typ[types.Int])
	case "hashmapNext":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("it", types.NewPointer(c.getRuntimeType("hashmapIterator")))
		addParam("key", types.Typ[types.UnsafePointer])
		addParam("value", types.Typ[types.UnsafePointer])
		addResult(types.Typ[types.Bool])
	case "hashmapStringGet":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.Typ[types.String])
		addParam("value", types.Typ[types.UnsafePointer])
		addParam("valueSize", types.Typ[types.Uintptr])
		addResult(types.Typ[types.Bool])
	case "hashmapStringSet":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.Typ[types.String])
		addParam("value", types.Typ[types.UnsafePointer])
	case "hashmapStringDelete":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.Typ[types.String])
	case "hashmapInterfaceGet":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.NewInterfaceType(nil, nil))
		addParam("value", types.Typ[types.UnsafePointer])
		addParam("valueSize", types.Typ[types.Uintptr])
		addResult(types.Typ[types.Bool])
	case "hashmapInterfaceSet":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.NewInterfaceType(nil, nil))
		addParam("value", types.Typ[types.UnsafePointer])
	case "hashmapInterfaceDelete":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.NewInterfaceType(nil, nil))
	case "hashmapBinaryGet":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.Typ[types.UnsafePointer])
		addParam("value", types.Typ[types.UnsafePointer])
		addParam("valueSize", types.Typ[types.Uintptr])
		addResult(types.Typ[types.Bool])
	case "hashmapBinarySet":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.Typ[types.UnsafePointer])
		addParam("value", types.Typ[types.UnsafePointer])
	case "hashmapBinaryDelete":
		addParam("m", types.NewPointer(c.getRuntimeType("hashmap")))
		addParam("key", types.Typ[types.UnsafePointer])
	case "tryChanSelect":
		addParam("recvbuf", types.Typ[types.UnsafePointer])
		addParam("states", types.NewSlice(c.getRuntimeType("chanSelectState")))
		addResult(types.Typ[types.Uintptr])
		addResult(types.Typ[types.Bool])
	case "chanMake":
		addParam("elementSize", types.Typ[types.Uintptr])
		addParam("bufSize", types.Typ[types.Uintptr])
		addResult(types.NewPointer(c.getRuntimeType("channel")))
	case "chanSend":
		addParam("ch", types.NewPointer(c.getRuntimeType("channel")))
		addParam("value", types.Typ[types.UnsafePointer])
	case "chanRecv":
		addParam("ch", types.NewPointer(c.getRuntimeType("channel")))
		addParam("value", types.Typ[types.UnsafePointer])
		addResult(types.Typ[types.Bool])
	case "chanClose":
		addParam("ch", types.NewPointer(c.getRuntimeType("channel")))
	case "chanSelect":
		addParam("recvbuf", types.Typ[types.UnsafePointer])
		addParam("states", types.NewSlice(c.getRuntimeType("chanSelectState")))
		addParam("ops", types.NewSlice(c.getRuntimeType("channelBlockedList")))
		addResult(types.Typ[types.Uintptr])
		addResult(types.Typ[types.Bool])
	case "deadlock":
	case "interfaceEqual":
		addParam("x", types.NewInterfaceType(nil, nil))
		addParam("y", types.NewInterfaceType(nil, nil))
		addResult(types.Typ[types.Bool])
	case "interfaceImplements":
		addParam("typecode", types.Typ[types.Uintptr])
		addParam("interfaceMethodSet", types.NewPointer(types.NewPointer(types.Typ[types.Uint8])))
		addResult(types.Typ[types.Bool])
	case "interfaceMethod":
		addParam("typecode", types.Typ[types.Uintptr])
		addParam("interfaceMethodSet", types.NewPointer(types.NewPointer(types.Typ[types.Uint8])))
		addParam("signature", types.NewPointer(types.Typ[types.Uint8]))
		addResult(types.Typ[types.Uintptr])
	case "typeAssert":
		addParam("actualType", types.Typ[types.Uintptr])
		addParam("assertedType", types.NewPointer(c.getRuntimeType("typecodeID")))
		addResult(types.Typ[types.Bool])
	case "interfaceTypeAssert":
		addParam("ok", types.Typ[types.Bool])
	case "getFuncPtr":
		addParam("val", c.getRuntimeType("funcValue"))
		addParam("signature", types.NewPointer(c.getRuntimeType("typecodeID")))
		addResult(types.Typ[types.Uintptr])
	default:
		panic("unknown runtime call: runtime." + name)
	}
	return types.NewSignature(nil, types.NewTuple(params...), types.NewTuple(results...), false)
}
