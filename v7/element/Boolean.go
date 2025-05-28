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

package element

import (
	fmt "fmt"
	reg "regexp"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func BooleanClass() BooleanClassLike {
	return booleanClass()
}

// Constructor Methods

func (c *booleanClass_) Boolean(
	boolean bool,
) BooleanLike {
	var instance = boolean_(boolean)
	return instance
}

func (c *booleanClass_) BooleanFromString(
	string_ string,
) BooleanLike {
	// Our booleans are more restrictive than the Go strconv package.
	if !booleanMatcher_.MatchString(string_) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the boolean constructor method: %s",
			string_,
		)
		panic(message)
	}
	var boolean, _ = stc.ParseBool(string_)
	return boolean_(boolean)
}

// Constant Methods

func (c *booleanClass_) False() BooleanLike {
	return c.false_
}

func (c *booleanClass_) True() BooleanLike {
	return c.true_
}

// Function Methods

func (c *booleanClass_) Not(
	boolean BooleanLike,
) BooleanLike {
	var result_ = boolean_(!boolean.AsBoolean())
	return result_
}

func (c *booleanClass_) And(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = boolean_(first.AsBoolean() && second.AsBoolean())
	return result_
}

func (c *booleanClass_) San(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = boolean_(first.AsBoolean() && !second.AsBoolean())
	return result_
}

func (c *booleanClass_) Ior(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = boolean_(first.AsBoolean() || second.AsBoolean())
	return result_
}

func (c *booleanClass_) Xor(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	var result_ = c.Ior(c.San(first, second), c.San(second, first))
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v boolean_) GetClass() BooleanClassLike {
	return booleanClass()
}

func (v boolean_) GetIntrinsic() bool {
	return bool(v)
}

// Attribute Methods

// Discrete Methods

func (v boolean_) AsBoolean() bool {
	var result_ = bool(v)
	return result_
}

func (v boolean_) AsInteger() int64 {
	var result_ int64
	if v {
		result_ = 1
	}
	return result_
}

// Lexical Methods

func (v boolean_) AsString() string {
	var result_ = stc.FormatBool(bool(v))
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
var booleanMatcher_ = reg.MustCompile("^(?:false|true)")

// Instance Structure

type boolean_ bool

// Class Structure

type booleanClass_ struct {
	// Declare the class constants.
	false_ BooleanLike
	true_  BooleanLike
}

// Class Reference

func booleanClass() *booleanClass_ {
	return booleanClassReference_
}

var booleanClassReference_ = &booleanClass_{
	// Initialize the class constants.
	false_: boolean_(false),
	true_:  boolean_(true),
}
