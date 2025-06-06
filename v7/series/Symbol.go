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
)

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

func (c *symbolClass_) SymbolFromSequence(
	sequence col.Sequential[rune],
) SymbolLike {
	var list = col.ListFromSequence[rune](sequence)
	return symbol_(list.AsArray())
}

func (c *symbolClass_) SymbolFromString(
	string_ string,
) SymbolLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the symbol constructor method: %s",
			string_,
		)
		panic(message)
	}
	return symbol_(matches[1]) // Strip off the leading "$".
}

// Constant Methods

// Function Methods

func (c *symbolClass_) Concatenate(
	first SymbolLike,
	second SymbolLike,
) SymbolLike {
	return c.Symbol(first.GetIntrinsic() + second.GetIntrinsic())
}

// INSTANCE INTERFACE

// Principal Methods

func (v symbol_) GetClass() SymbolClassLike {
	return symbolClass()
}

func (v symbol_) GetIntrinsic() string {
	return string(v)
}

func (v symbol_) AsString() string {
	return "$" + v.GetIntrinsic()
}

// Attribute Methods

// col.Sequential[rune] Methods

func (v symbol_) IsEmpty() bool {
	return len(v) == 0
}

func (v symbol_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v.AsArray()))
}

func (v symbol_) AsArray() []rune {
	return []rune(v)
}

func (v symbol_) GetIterator() col.IteratorLike[rune] {
	var iteratorClass = col.IteratorClass[rune]()
	var iterator = iteratorClass.Iterator(v.AsArray())
	return iterator
}

// col.Accessible[rune] Methods

func (v symbol_) GetValue(
	index col.Index,
) rune {
	var list = col.ListFromArray[rune](v.AsArray())
	return list.GetValue(index)
}

func (v symbol_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[rune] {
	var list = col.ListFromArray[rune](v.AsArray())
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v symbol_) String() string {
	return v.AsString()
}

// Private Methods

// Instance Structure

type symbol_ string

// Class Structure

type symbolClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func symbolClass() *symbolClass_ {
	return symbolClassReference_
}

var symbolClassReference_ = &symbolClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^\\$(" + identifier_ + ")"),
}
