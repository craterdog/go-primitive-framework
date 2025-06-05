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

func QuoteClass() QuoteClassLike {
	return quoteClass()
}

// Constructor Methods

func (c *quoteClass_) Quote(
	string_ string,
) QuoteLike {
	reg.MustCompile(string_)
	return quote_(string_)
}

func (c *quoteClass_) QuoteFromSequence(
	sequence col.Sequential[rune],
) QuoteLike {
	var instance QuoteLike
	// TBD - Add the constructor implementation.
	return instance
}

func (c *quoteClass_) QuoteFromString(
	string_ string,
) QuoteLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the quote constructor method: %s",
			string_,
		)
		panic(message)
	}
	switch matches[0] {
	case "none":
		return c.none_
	case "any":
		return c.any_
	default:
		reg.MustCompile(matches[1]) // Make sure it is a valid regular expression.
		return quote_(matches[1]) // Strip off the double quotes and '?'.
	}
}

// Constant Methods

func (c *quoteClass_) None() QuoteLike {
	return c.none_
}

func (c *quoteClass_) Any() QuoteLike {
	return c.any_
}

// Function Methods

func (c *quoteClass_) Concatenate(
	first QuoteLike,
	second QuoteLike,
) QuoteLike {
	return c.Quote(first.GetIntrinsic() + second.GetIntrinsic())
}

// INSTANCE INTERFACE

// Principal Methods

func (v quote_) GetClass() QuoteClassLike {
	return quoteClass()
}

func (v quote_) GetIntrinsic() string {
	return string(v)
}

func (v quote_) AsString() string {
	var string_ string
	switch v {
	case `^none$`:
		string_ = `none`
	case `.*`:
		string_ = `any`
	default:
		string_ = `"` + string(v) + `"?`
	}
	return string_
}

func (v quote_) AsRegexp() *reg.Regexp {
	return reg.MustCompile(string(v))
}

func (v quote_) MatchesText(
	text string,
) bool {
	var matcher = reg.MustCompile(string(v))
	return matcher.MatchString(text)
}

func (v quote_) GetMatches(
	text string,
) []string {
	var matcher = reg.MustCompile(string(v))
	return matcher.FindStringSubmatch(text)
}

// Attribute Methods

// col.Sequential[rune] Methods

func (v quote_) IsEmpty() bool {
	return len(v) == 0
}

func (v quote_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v))
}

func (v quote_) AsArray() []rune {
	return []rune(v)
}

func (v quote_) GetIterator() col.IteratorLike[rune] {
	var iteratorClass = col.IteratorClass[rune]()
	var iterator = iteratorClass.Iterator([]rune(v))
	return iterator
}

// col.Accessible[rune] Methods

func (v quote_) GetValue(
	index col.Index,
) rune {
	var list = col.ListFromArray[rune]([]rune(v))
	return list.GetValue(index)
}

func (v quote_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[rune] {
	var list = col.ListFromArray[rune]([]rune(v))
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v quote_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string quotes for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	character_ = "(?:(?:" + escape_ + ")|\\\\\"|[^\"" + control_ + "])"
	escape_    = "(?:\\\\((?:" + unicode_ + ")|[abfnrtv\\\\]))"
	unicode_   = "(?:(u(?:" + base16_ + "){4})|(U(?:" + base16_ + "){8}))"
	control_   = "\\p{Cc}"
)

// Instance Structure

type quote_ string

// Class Structure

type quoteClass_ struct {
	// Declare the class constants.
	none_    QuoteLike
	any_     QuoteLike
	matcher_ *reg.Regexp
}

// Class Reference

func quoteClass() *quoteClass_ {
	return quoteClassReference_
}

var quoteClassReference_ = &quoteClass_{
	// Initialize the class constants.
	none_:    quote_(`^none$`),
	any_:     quote_(`.*`),
	matcher_: reg.MustCompile("^(?:\"(?:" + character_ + ")*\")"),
}
