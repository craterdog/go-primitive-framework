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

import ()

// CLASS INTERFACE

// Access Function

func TagClass() TagClassLike {
	return tagClass()
}

// Constructor Methods

func (c *tagClass_) Tag(
	string_ string,
) TagLike {
	return tag_(string_)
}

func (c *tagClass_) TagWithSize(
	size int,
) TagLike {
	return tag_("")
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v tag_) GetClass() TagClassLike {
	return tagClass()
}

func (v tag_) GetIntrinsic() string {
	return string(v)
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type tag_ string

// Class Structure

type tagClass_ struct {
	// Declare the class constants.
}

// Class Reference

func tagClass() *tagClass_ {
	return tagClassReference_
}

var tagClassReference_ = &tagClass_{
	// Initialize the class constants.
}
