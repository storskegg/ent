// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"github.com/storskegg/ent/entc/integration/privacy/ent/privacy"
	"github.com/storskegg/ent/entc/integration/privacy/rule"
	"github.com/storskegg/ent/schema/mixin"
)

// BaseMixin for all schemas.
type BaseMixin struct {
	mixin.Schema
}

// Shared policy for all schemas.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
		},
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
		},
	}
}
