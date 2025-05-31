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
		string_ string,
	) BinaryLike
	BinaryFromArray(
		array []byte,
	) BinaryLike
	BinaryFromSequence(
		sequence Sequential[byte],
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
		string_ string,
	) BytecodeLike
	BytecodeFromArray(
		array []Instruction,
	) BytecodeLike
	BytecodeFromSequence(
		sequence Sequential[Instruction],
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
		string_ string,
	) NameLike
	NameFromArray(
		array []Identifier,
	) NameLike
	NameFromSequence(
		sequence Sequential[Identifier],
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
		string_ string,
	) NarrativeLike
	NarrativeFromArray(
		array []Line,
	) NarrativeLike
	NarrativeFromSequence(
		sequence Sequential[Line],
	) NarrativeLike

	// Function Methods
	Concatenate(
		first NarrativeLike,
		second NarrativeLike,
	) NarrativeLike
}

/*
QuoteClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
quote-like concrete class.
*/
type QuoteClassLike interface {
	// Constructor Methods
	Quote(
		string_ string,
	) QuoteLike
	QuoteFromArray(
		array []rune,
	) QuoteLike
	QuoteFromSequence(
		sequence Sequential[rune],
	) QuoteLike

	// Function Methods
	Concatenate(
		first QuoteLike,
		second QuoteLike,
	) QuoteLike
}

/*
SymbolClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
symbol-like concrete class.
*/
type SymbolClassLike interface {
	// Constructor Methods
	Symbol(
		string_ string,
	) SymbolLike
	SymbolFromArray(
		array []rune,
	) SymbolLike
	SymbolFromSequence(
		sequence Sequential[rune],
	) SymbolLike
}

/*
TagClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
tag-like concrete class.
*/
type TagClassLike interface {
	// Constructor Methods
	Tag(
		string_ string,
	) TagLike
	TagWithSize(
		size age.Size,
	) TagLike
	TagFromArray(
		array []byte,
	) TagLike
	TagFromSequence(
		sequence Sequential[byte],
	) TagLike
}

/*
VersionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
version-like concrete class.
*/
type VersionClassLike interface {
	// Constructor Methods
	VersionFromString(
		string_ string,
	) VersionLike
	VersionFromArray(
		array []Ordinal,
	) VersionLike
	VersionFromSequence(
		sequence Sequential[Ordinal],
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

	// Aspect Interfaces
	Intrinsic[[]byte]
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

	// Aspect Interfaces
	Intrinsic[[]Instruction]
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

	// Aspect Interfaces
	Intrinsic[[]Identifier]
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

	// Aspect Interfaces
	Intrinsic[string]
	Sequential[Line]
}

/*
QuoteLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete quote-like class.
*/
type QuoteLike interface {
	// Principal Methods
	GetClass() QuoteClassLike

	// Aspect Interfaces
	Intrinsic[string]
	Sequential[rune]
}

/*
SymbolLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete symbol-like class.
*/
type SymbolLike interface {
	// Principal Methods
	GetClass() SymbolClassLike

	// Aspect Interfaces
	Intrinsic[string]
	Sequential[rune]
}

/*
TagLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete tag-like class.
*/
type TagLike interface {
	// Principal Methods
	GetClass() TagClassLike

	// Aspect Interfaces
	Intrinsic[string]
	Sequential[byte]
}

/*
VersionLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete version-like class.
*/
type VersionLike interface {
	// Principal Methods
	GetClass() VersionClassLike

	// Aspect Interfaces
	Intrinsic[[]Ordinal]
	Sequential[Ordinal]
}

// ASPECT DECLARATIONS

/*
Intrinsic[V any] is an aspect interface that declares a set of method signatures
that must be supported by each instance of an intrinsic concrete class.
*/
type Intrinsic[V any] interface {
	GetIntrinsic() V
	AsString() string
}

/*
Sequential[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a sequential concrete
class.
*/
type Sequential[V any] interface {
	IsEmpty() bool
	GetSize() age.Size
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
