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
	age "github.com/craterdog/go-primitive-framework/v7/agent"
)

// CLASS INTERFACE

// Access Function

func BinaryClass() BinaryClassLike {
	return binaryClass()
}

// Constructor Methods

func (c *binaryClass_) Binary(
	intrinsic []byte,
) BinaryLike {
	return binary_(intrinsic)
}

func (c *binaryClass_) BinaryFromSequence(
	sequence Sequential[byte],
) BinaryLike {
	var instance BinaryLike
	// TBD - Add the constructor implementation.
	return instance
}

func (c *binaryClass_) BinaryFromBase64(
	base64 string,
) BinaryLike {
	var instance BinaryLike
	// TBD - Add the constructor implementation.
	return instance
}

// Constant Methods

// Function Methods

func (c *binaryClass_) Not(
	binary BinaryLike,
) BinaryLike {
	var result_ BinaryLike
	// TBD - Add the function implementation.
	return result_
}

func (c *binaryClass_) And(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result_ BinaryLike
	// TBD - Add the function implementation.
	return result_
}

func (c *binaryClass_) San(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result_ BinaryLike
	// TBD - Add the function implementation.
	return result_
}

func (c *binaryClass_) Ior(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result_ BinaryLike
	// TBD - Add the function implementation.
	return result_
}

func (c *binaryClass_) Xor(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result_ BinaryLike
	// TBD - Add the function implementation.
	return result_
}

func (c *binaryClass_) Concatenate(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result_ BinaryLike
	// TBD - Add the function implementation.
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v binary_) GetClass() BinaryClassLike {
	return binaryClass()
}

func (v binary_) GetIntrinsic() []byte {
	return []byte(v)
}

// Attribute Methods

// Sequential[byte] Methods

func (v binary_) IsEmpty() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v binary_) GetSize() uint {
	var result_ uint
	// TBD - Add the method implementation.
	return result_
}

func (v binary_) AsArray() []byte {
	var result_ []byte
	// TBD - Add the method implementation.
	return result_
}

func (v binary_) GetValue(
	index Index,
) byte {
	var result_ byte
	// TBD - Add the method implementation.
	return result_
}

func (v binary_) GetValues(
	first Index,
	last Index,
) Sequential[byte] {
	var result_ Sequential[byte]
	// TBD - Add the method implementation.
	return result_
}

func (v binary_) GetIterator() age.IteratorLike[byte] {
	var result_ age.IteratorLike[byte]
	// TBD - Add the method implementation.
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type binary_ []byte

// Class Structure

type binaryClass_ struct {
	// Declare the class constants.
}

// Class Reference

func binaryClass() *binaryClass_ {
	return binaryClassReference_
}

var binaryClassReference_ = &binaryClass_{
	// Initialize the class constants.
}
