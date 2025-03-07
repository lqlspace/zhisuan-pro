/*
 * Copyright (c) 2024, VIRTAITECH CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dcgmexporter

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitWithTimeout(t *testing.T) {
	t.Run("Returns error by timeout", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		defer wg.Done()
		wg.Add(1)
		timeout := 500 * time.Millisecond
		err := WaitWithTimeout(wg, timeout)
		require.Error(t, err)
		assert.ErrorContains(t, err, "timeout waiting for WaitGroup")
	})

	t.Run("Returns no error", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		timeout := 500 * time.Millisecond
		wg.Done()
		err := WaitWithTimeout(wg, timeout)
		require.NoError(t, err)
	})
}

func TestDeepCopy(t *testing.T) {
	t.Run("Return error when pointer value is nil", func(t *testing.T) {
		got, err := deepCopy[*struct{}](nil)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("Return error when src is unsupported type", func(t *testing.T) {
		ch := make(chan int)
		got, err := deepCopy(ch)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}
