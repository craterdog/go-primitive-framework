/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package series

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func BinaryClass() BinaryClassLike {
	return binaryClass()
}

// Constructor Methods

func (c *binaryClass_) Binary(
	bytes []byte,
) BinaryLike {
	return binary_(bytes)
}

func (c *binaryClass_) BinaryFromSequence(
	sequence col.Sequential[byte],
) BinaryLike {
	var list = col.ListFromSequence[byte](sequence)
	return binary_(list.AsArray())
}

func (c *binaryClass_) BinaryFromString(
	string_ string,
) BinaryLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the binary constructor method: %s",
			string_,
		)
		panic(message)
	}
	var base64 = matches[1]                   // Strip off the delimiters.
	base64 = sts.ReplaceAll(base64, " ", "")  // Remove all spaces.
	base64 = sts.ReplaceAll(base64, "\n", "") // Remove all newlines.
	var bytes = uti.Base64Decode(base64)
	return binary_(bytes)
}

// Constant Methods

// Function Methods

func (c *binaryClass_) Not(
	binary BinaryLike,
) BinaryLike {
	var bytes = binary.AsArray()
	var size = len(bytes)
	for i := 0; i < size; i++ {
		bytes[i] = ^bytes[i]
	}
	return c.Binary(bytes)
}

func (c *binaryClass_) And(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] & secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) San(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] &^ secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Ior(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] | secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Xor(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] ^ secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Concatenate(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make(
		[]byte,
		len(firstBytes)+len(secondBytes),
	)
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return c.Binary(allBytes)
}

// INSTANCE INTERFACE

// Principal Methods

func (v binary_) GetClass() BinaryClassLike {
	return binaryClass()
}

func (v binary_) GetIntrinsic() []byte {
	return []byte(v)
}

func (v binary_) AsString() string {
	var encoded = uti.Base64Encode(v)
	var length = len(encoded)
	var string_ = "'>\n"
	if length > 0 {
		var width = 60
		var indentation = "    "
		var index int
		for index = 0; index+width < length; index += width {
			string_ += indentation + encoded[index:index+width] + "\n"
		}
		string_ += indentation + encoded[index:] + "\n"
	}
	string_ += "<'"
	return string_
}

// Attribute Methods

// col.Sequential[byte] Methods

func (v binary_) IsEmpty() bool {
	return len(v) == 0
}

func (v binary_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v))
}

func (v binary_) AsArray() []byte {
	return uti.CopyArray(v)
}

func (v binary_) GetIterator() col.IteratorLike[byte] {
	var array = uti.CopyArray(v)
	var iteratorClass = col.IteratorClass[byte]()
	var iterator = iteratorClass.Iterator(array)
	return iterator
}

// col.Accessible[byte] Methods

func (v binary_) GetValue(
	index col.Index,
) byte {
	var list = col.ListFromArray[byte](v)
	return list.GetValue(index)
}

func (v binary_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[byte] {
	var list = col.ListFromArray[byte](v)
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v binary_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	alpha_        = "(?:[A-Za-z])"
	alphanumeric_ = "(?:(?:" + alpha_ + ")|(?:" + base10_ + "))"
	base10_       = "(?:[0-9])"
	base64_       = "(?:(?:" + alphanumeric_ + ")|[\\+/])"
	block_        = "(?: *(?:" + base64_ + "){1,60}" + eol_ + ")"
	eol_          = "\\r?\\n"
)

// Instance Structure

type binary_ []byte

// Class Structure

type binaryClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func binaryClass() *binaryClass_ {
	return binaryClassReference_
}

var binaryClassReference_ = &binaryClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^(?:'>(" + eol_ + "(?:" + block_ + ")+)? *<')",
	),
}
