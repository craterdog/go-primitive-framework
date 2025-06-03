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
Package "agent" provides a framework of aspects and class definitions for agents
that act on a rich set of primitive data types that can be iterated over.

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
package agent

import ()

// TYPE DECLARATIONS

/*
Slot is a constrained type representing the slot between to items in a sequence.
*/
type Slot uint

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
IteratorClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
iterator-like concrete class.
*/
type IteratorClassLike[V any] interface {
	// Constructor Methods
	Iterator(
		array []V,
	) IteratorLike[V]
}

// INSTANCE DECLARATIONS

/*
IteratorLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete iterator-like class.
*/
type IteratorLike[V any] interface {
	// Principal Methods
	GetClass() IteratorClassLike[V]
	IsEmpty() bool
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() V
	HasNext() bool
	GetNext() V

	// Attribute Methods
	GetSize() uint
	GetSlot() Slot
	SetSlot(
		slot Slot,
	)
}

// ASPECT DECLARATIONS
