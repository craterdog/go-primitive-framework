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
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func VersionClass() VersionClassLike {
	return versionClass()
}

// Constructor Methods

func (c *versionClass_) Version(
	ordinals []uti.Ordinal,
) VersionLike {
	return version_(ordinals)
}

func (c *versionClass_) VersionFromSequence(
	sequence col.Sequential[uti.Ordinal],
) VersionLike {
	var list = col.ListFromSequence[uti.Ordinal](sequence)
	return version_(list.AsArray())
}

func (c *versionClass_) VersionFromString(
	string_ string,
) VersionLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the version constructor method: %s",
			string_,
		)
		panic(message)
	}
	var match = matches[1] // Strip off the leading "v".
	var levels = sts.Split(match, ".")
	var ordinals = make([]uti.Ordinal, len(levels))
	for index, level := range levels {
		var ordinal, _ = stc.ParseUint(level, 10, 64)
		ordinals[index] = uti.Ordinal(ordinal)
	}
	return version_(ordinals)
}

// Constant Methods

// Function Methods

func (c *versionClass_) IsValidNextVersion(
	current VersionLike,
	next VersionLike,
) bool {
	// Make sure the version sizes are compatible.
	var currentOrdinals = current.AsArray()
	var currentSize = len(currentOrdinals)
	var nextOrdinals = next.AsArray()
	var nextSize = len(nextOrdinals)
	if nextSize > currentSize+1 {
		return false
	}

	// Iterate through the versions comparing level values.
	var currentIterator = col.Iterator[uti.Ordinal](current.AsArray())
	var nextIterator = col.Iterator[uti.Ordinal](next.AsArray())
	for currentIterator.HasNext() && nextIterator.HasNext() {
		var currentLevel = currentIterator.GetNext()
		var nextLevel = nextIterator.GetNext()
		if currentLevel == nextLevel {
			// So far the level values are the same.
			continue
		}
		// The last level for the next version must be one more.
		return !nextIterator.HasNext() && nextLevel == currentLevel+1
	}
	// The last level for the next version must be one.
	return nextIterator.HasNext() && nextIterator.GetNext() == 1
}

func (c *versionClass_) GetNextVersion(
	current VersionLike,
	level uti.Ordinal,
) VersionLike {
	// Adjust the size of the ordinals as needed.
	var ordinals = current.AsArray()
	var size = uti.Ordinal(len(ordinals))
	switch {
	case level == 0:
		level = size // Normalize the level to the current size.
	case level < size:
		// The next version will require fewer levels.
		ordinals = ordinals[:level]
	case level > size:
		// The next version will require another level.
		size++
		level = size // Normalize the level to the new size.
		ordinals = append(ordinals, 0)
	}

	// Increment the specified version level.
	var index = level - 1 // Convert to zero based indexing.
	ordinals[index]++

	var version = c.Version(ordinals)
	return version
}

func (c *versionClass_) Concatenate(
	first VersionLike,
	second VersionLike,
) VersionLike {
	var firstOrdinals = first.AsArray()
	var secondOrdinals = second.AsArray()
	var allOrdinals = make(
		[]uti.Ordinal,
		len(firstOrdinals)+len(secondOrdinals),
	)
	copy(allOrdinals, firstOrdinals)
	copy(allOrdinals[len(firstOrdinals):], secondOrdinals)
	return c.Version(allOrdinals)
}

// INSTANCE INTERFACE

// Principal Methods

func (v version_) GetClass() VersionClassLike {
	return versionClass()
}

func (v version_) GetIntrinsic() []uti.Ordinal {
	return []uti.Ordinal(v)
}

func (v version_) AsString() string {
	var index = 0
	var string_ = "v" + stc.Itoa(int(v[index]))
	for index++; index < len(v); index++ {
		string_ += "." + stc.Itoa(int(v[index]))
	}
	return string_
}

// Attribute Methods

// col.Sequential[uti.Ordinal] Methods

func (v version_) IsEmpty() bool {
	return len(v) == 0
}

func (v version_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v))
}

func (v version_) AsArray() []uti.Ordinal {
	return uti.CopyArray(v)
}

func (v version_) GetIterator() col.IteratorLike[uti.Ordinal] {
	var array = uti.CopyArray(v)
	var iteratorClass = col.IteratorClass[uti.Ordinal]()
	var iterator = iteratorClass.Iterator(array)
	return iterator
}

// col.Accessible[uti.Ordinal] Methods

func (v version_) GetValue(
	index col.Index,
) uti.Ordinal {
	var list = col.ListFromArray[uti.Ordinal](v)
	return list.GetValue(index)
}

func (v version_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[uti.Ordinal] {
	var list = col.ListFromArray[uti.Ordinal](v)
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v version_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each version to lessen the chance of a version collision with other private Go
// class constants in this package.
const (
	ordinal_ = "(?:[1-9](?:" + base10_ + ")*)"
)

// Instance Structure

type version_ []uti.Ordinal

// Class Structure

type versionClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func versionClass() *versionClass_ {
	return versionClassReference_
}

var versionClassReference_ = &versionClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"(?:v((?:" + ordinal_ + ")(?:\\.(?:" + ordinal_ + "))*))",
	),
}
