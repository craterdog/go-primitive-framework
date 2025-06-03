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

func SymbolClass() SymbolClassLike {
	return symbolClass()
}

// Constructor Methods

func (c *symbolClass_) Symbol(
	string_ string,
) SymbolLike {
	return symbol_(string_)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v symbol_) GetClass() SymbolClassLike {
	return symbolClass()
}

func (v symbol_) GetIntrinsic() string {
	return string(v)
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type symbol_ string

// Class Structure

type symbolClass_ struct {
	// Declare the class constants.
}

// Class Reference

func symbolClass() *symbolClass_ {
	return symbolClassReference_
}

var symbolClassReference_ = &symbolClass_{
	// Initialize the class constants.
}
