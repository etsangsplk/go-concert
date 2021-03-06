// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package unison

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiErrGroup(t *testing.T) {
	t.Run("returns empty list if no go-routine was started", func(t *testing.T) {
		var grp MultiErrGroup
		assert.Equal(t, 0, len(grp.Wait()))
	})

	t.Run("returns empty list if no go-routine failed", func(t *testing.T) {
		var grp MultiErrGroup
		grp.Go(func() error { return nil })
		assert.Equal(t, 0, len(grp.Wait()))
	})

	t.Run("Returns multiple errors", func(t *testing.T) {
		var grp MultiErrGroup
		grp.Go(func() error { return errors.New("1") })
		grp.Go(func() error { return errors.New("2") })
		assert.Equal(t, 2, len(grp.Wait()))
	})
}
