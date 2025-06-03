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

/*
Package "series" provides a framework of aspects and class definitions for a
rich set of primitive data types that can be iterated over.  All primitive types
are immutable and—for better performance—are implemented as extensions to
existing Go primitive types.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-primitive-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package series

import (
	age "github.com/craterdog/go-primitive-framework/v7/agent"
)

// TYPE DECLARATIONS

/*
Identifier is a constrained type representing a string of the form:
(LOWER | UPPER) (LOWER | UPPER | DIGIT | '-')
*/
type Identifier string

/*
Index is a constrained type representing the positive (or negative) ORDINAL
index of a value in a sequence.  The indices are ordinal rather than zero-based
which never really made sense except for pointer offsets.  What is the "zeroth
value" anyway?  It's the "first value", right?  So we start a fresh...

This approach allows for positive indices starting at the beginning of a
sequence—and negative indices starting at the end of the sequence, as follows:

	    1           2           3             N
	[value 1] . [value 2] . [value 3] ... [value N]
	   -N        -(N-1)      -(N-2)          -1

Notice that because the indices are ordinal based, the positive and negative
indices are symmetrical.  An index can NEVER be zero.
*/
type Index int

/*
Instruction is a constrained type representing a single bytecode instruction.
*/
type Instruction uint16

/*
Line is a constrained type representing a single line of a narrative.
*/
type Line string

/*
Ordinal is a constrained type representing a single ordinal number [1..MaxUint).
*/
type Ordinal uint

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
BinaryClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
binary-like concrete class.
*/
type BinaryClassLike interface {
	// Constructor Methods
	Binary(
		intrinsic []byte,
	) BinaryLike
	BinaryFromSequence(
		sequence Sequential[byte],
	) BinaryLike
	BinaryFromBase64(
		base64 string,
	) BinaryLike

	// Function Methods
	Not(
		binary BinaryLike,
	) BinaryLike
	And(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	San(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Ior(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Xor(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Concatenate(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
}

/*
BytecodeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
bytecode-like concrete class.
*/
type BytecodeClassLike interface {
	// Constructor Methods
	Bytecode(
		intrinsic []Instruction,
	) BytecodeLike
	BytecodeFromSequence(
		sequence Sequential[Instruction],
	) BytecodeLike
	BytecodeFromString(
		string_ string,
	) BytecodeLike
}

/*
NameClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
name-like concrete class.
*/
type NameClassLike interface {
	// Constructor Methods
	Name(
		intrinsic []Identifier,
	) NameLike
	NameFromSequence(
		sequence Sequential[Identifier],
	) NameLike
	NameFromString(
		string_ string,
	) NameLike

	// Function Methods
	Concatenate(
		first NameLike,
		second NameLike,
	) NameLike
}

/*
NarrativeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
narrative-like concrete class.
*/
type NarrativeClassLike interface {
	// Constructor Methods
	Narrative(
		intrinsic []Line,
	) NarrativeLike
	NarrativeFromSequence(
		sequence Sequential[Line],
	) NarrativeLike
	NarrativeFromString(
		string_ string,
	) NarrativeLike

	// Function Methods
	Concatenate(
		first NarrativeLike,
		second NarrativeLike,
	) NarrativeLike
}

/*
VersionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
version-like concrete class.
*/
type VersionClassLike interface {
	// Constructor Methods
	Version(
		intrinsic []Ordinal,
	) VersionLike
	VersionFromSequence(
		sequence Sequential[Ordinal],
	) VersionLike
	VersionFromString(
		string_ string,
	) VersionLike

	// Function Methods
	IsValidNextVersion(
		current VersionLike,
		next VersionLike,
	) bool
	GetNextVersion(
		current VersionLike,
		level Ordinal,
	) VersionLike
	Concatenate(
		first VersionLike,
		second VersionLike,
	) VersionLike
}

// INSTANCE DECLARATIONS

/*
BinaryLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete binary-like class.
*/
type BinaryLike interface {
	// Principal Methods
	GetClass() BinaryClassLike
	GetIntrinsic() []byte

	// Aspect Interfaces
	Sequential[byte]
}

/*
BytecodeLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete bytecode-like class.
*/
type BytecodeLike interface {
	// Principal Methods
	GetClass() BytecodeClassLike
	GetIntrinsic() []Instruction

	// Aspect Interfaces
	Sequential[Instruction]
}

/*
NameLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete name-like class.
*/
type NameLike interface {
	// Principal Methods
	GetClass() NameClassLike
	GetIntrinsic() []Identifier

	// Aspect Interfaces
	Sequential[Identifier]
}

/*
NarrativeLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete narrative-like class.
*/
type NarrativeLike interface {
	// Principal Methods
	GetClass() NarrativeClassLike
	GetIntrinsic() []Line

	// Aspect Interfaces
	Sequential[Line]
}

/*
VersionLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete version-like class.
*/
type VersionLike interface {
	// Principal Methods
	GetClass() VersionClassLike
	GetIntrinsic() []Ordinal

	// Aspect Interfaces
	Sequential[Ordinal]
}

// ASPECT DECLARATIONS

/*
Sequential[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a sequential concrete
class.
*/
type Sequential[V any] interface {
	IsEmpty() bool
	GetSize() uint
	AsArray() []V
	GetValue(
		index Index,
	) V
	GetValues(
		first Index,
		last Index,
	) Sequential[V]
	GetIterator() age.IteratorLike[V]
}
