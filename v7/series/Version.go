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
	col "github.com/craterdog/go-collection-framework/v2"
	age "github.com/craterdog/go-primitive-framework/v7/agent"
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

func (c *versionClass_) VersionFromString(
	string_ string,
) VersionLike {
	var instance VersionLike
	// TBD - Add the constructor implementation.
	return instance
}

func (c *versionClass_) VersionFromArray(
	array []Ordinal,
) VersionLike {
	var instance VersionLike
	// TBD - Add the constructor implementation.
	return instance
}

func (c *versionClass_) VersionFromSequence(
	sequence Sequential[Ordinal],
) VersionLike {
	var instance VersionLike
	// TBD - Add the constructor implementation.
	return instance
}

// Constant Methods

// Function Methods

// This function determines whether or not the specified next version string
// is a valid next version string for the specified current version string. In
// order for the next version to be valid the last level in the next version
// string must be one more than the corresponding level in the current version
// string; or it must be '1' and the next version string must have one more
// level of version numbers than the current version string, for example:
//
//	           Current             Next          What Likely Changed
//	level 1:    v5.7              v6         (interface/symantic changes)
//	level 2:    v5.7              v5.8       (optimization/bug fixes)
//	level 3:    v5.7              v5.7.1     (changes being tested)
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
	var currentIterator = col.Iterator[Ordinal](current)
	var nextIterator = col.Iterator[Ordinal](next)
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

// This function returns a copy of the specified version string with the ordinal
// at the specified level incremented by one. For example:
//
//	           Current             Next          What Likely Changed
//	level 1:    v5.7              v6         (interface/symantic changes)
//	level 2:    v5.7              v5.8       (optimization/bug fixes)
//	level 3:    v5.7              v5.7.1     (changes being tested)
//
// A level of `-1` will result in the last ordinal in the copy of the current
// version string being incremented. A level that is greater than the size of
// current version will result in a new level with the value of `1` being
// appended to the copy of the current version string.
func (c *versionClass_) GetNextVersion(
	current VersionLike,
	level Ordinal,
) VersionLike {
	// Adjust the size of the ordinals as needed.
	var ordinals = current.AsArray()
	var size = Ordinal(len(ordinals))
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

	var version = VersionFromArray(ordinals)
	return version
}

func (c *versionClass_) Concatenate(
	first VersionLike,
	second VersionLike,
) VersionLike {
	var version = first.AsString() + "." + second.AsString()
	return Version(version)
}

// INSTANCE INTERFACE

// Principal Methods

func (v version_) GetClass() VersionClassLike {
	return versionClass()
}

// Attribute Methods

// Intrinsic[[]Ordinal] Methods

func (v version_) GetIntrinsic() []Ordinal {
	var result_ []Ordinal
	// TBD - Add the method implementation.
	return result_
}

func (v version_) AsString() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

// Sequential[Ordinal] Methods

func (v version_) IsEmpty() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v version_) GetSize() age.Size {
	var result_ age.Size
	// TBD - Add the method implementation.
	return result_
}

func (v version_) AsArray() []Ordinal {
	var result_ []Ordinal
	// TBD - Add the method implementation.
	return result_
}

func (v version_) GetValue(
	index Index,
) Ordinal {
	var result_ Ordinal
	// TBD - Add the method implementation.
	return result_
}

func (v version_) GetValues(
	first Index,
	last Index,
) Sequential[Ordinal] {
	var result_ Sequential[Ordinal]
	// TBD - Add the method implementation.
	return result_
}

func (v version_) GetIterator() age.IteratorLike[Ordinal] {
	var result_ age.IteratorLike[Ordinal]
	// TBD - Add the method implementation.
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type version_ struct {
	// Declare the instance attributes.
}

// Class Structure

type versionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func versionClass() *versionClass_ {
	return versionClassReference_
}

var versionClassReference_ = &versionClass_{
	// Initialize the class constants.
}
