/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package series_test

import (
	ser "github.com/craterdog/go-primitive-framework/v7/series"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

var BinaryClass = ser.BinaryClass()

func TestEmptyBinary(t *tes.T) {
	var binary = `'> <'`
	var v = BinaryClass.BinaryFromString(binary)
	ass.Equal(t, binary, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestBinary(t *tes.T) {
	var b1 = `'>
    abcd1234
<'`
	var b2 = `'>
    abcd
<'`
	var v = BinaryClass.BinaryFromString(b1)
	ass.Equal(t, b1, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, int(v.GetSize()))
	ass.Equal(t, byte(0x69), v.GetValue(1))
	ass.Equal(t, byte(0xf8), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), BinaryClass.Binary(v.AsArray()).AsArray())
	ass.Equal(t, b2, BinaryClass.BinaryFromSequence(v.GetValues(1, 3)).AsString())
}

func TestBinaryLibrary(t *tes.T) {
	var b1 = `'>
    abcd
<'`
	var b2 = `'>
    12345678
<'`
	var b3 = `'>
    abcd12345678
<'`
	var v1 = BinaryClass.BinaryFromString(b1)
	var v2 = BinaryClass.BinaryFromString(b2)
	ass.Equal(t, b3, BinaryClass.Concatenate(v1, v2).AsString())

	v1 = BinaryClass.Binary([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = BinaryClass.Binary([]byte{0x03, 0x00, 0x01, 0x02})
	var not = BinaryClass.Binary([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = BinaryClass.Binary([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = BinaryClass.Binary([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = BinaryClass.Binary([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = BinaryClass.Binary([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = BinaryClass.Binary([]byte{0x03, 0x00, 0x01, 0x00, 0x00})
	ass.Equal(t, not, BinaryClass.Not(v1))
	ass.Equal(t, and, BinaryClass.And(v1, v2))
	ass.Equal(t, sans, BinaryClass.San(v1, v2))
	ass.Equal(t, or, BinaryClass.Ior(v1, v2))
	ass.Equal(t, xor, BinaryClass.Xor(v1, v2))
	ass.Equal(t, sans2, BinaryClass.San(v2, v1))
}

var BytecodeClass = ser.BytecodeClass()

func TestEmptyBytecode(t *tes.T) {
	var bytecode = `''`
	var v = BytecodeClass.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestBytecode(t *tes.T) {
	var bytecode = `'abcd'`
	var v = BytecodeClass.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 1, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), BytecodeClass.Bytecode(v.AsArray()).AsArray())

	bytecode = `'abcd 1234'`
	v = BytecodeClass.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), BytecodeClass.Bytecode(v.AsArray()).AsArray())
}

var NameClass = ser.NameClass()

func TestName(t *tes.T) {
	var v1 = NameClass.NameFromString("/bali/types/abstractions/String")
	ass.Equal(t, "/bali/types/abstractions/String", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, int(v1.GetSize()))
	ass.Equal(t, ser.Identifier("bali"), v1.GetValue(1))
	ass.Equal(t, ser.Identifier("String"), v1.GetValue(-1))
	var v2 = NameClass.Name(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = NameClass.NameFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, "/bali/types", v3.AsString())
}

func TestNamesLibrary(t *tes.T) {
	var v1 = NameClass.NameFromString("/bali/types/abstractions")
	var v2 = NameClass.NameFromString("/String")
	ass.Equal(t, "/bali/types/abstractions/String", NameClass.Concatenate(v1, v2).AsString())
}
