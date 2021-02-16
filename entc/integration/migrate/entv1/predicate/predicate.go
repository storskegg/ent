// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package predicate

import (
	"github.com/storskegg/ent/dialect/sql"
)

// Car is the predicate function for car builders.
type Car func(*sql.Selector)

// Conversion is the predicate function for conversion builders.
type Conversion func(*sql.Selector)

// CustomType is the predicate function for customtype builders.
type CustomType func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)
