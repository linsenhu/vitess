// Copyright 2013, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync2

import (
	"sync/atomic"
	"time"
)

type AtomicInt32 int32

func (i *AtomicInt32) Add(n int32) int32 {
	return atomic.AddInt32((*int32)(i), n)
}

func (i *AtomicInt32) Set(n int32) {
	atomic.StoreInt32((*int32)(i), n)
}

func (i *AtomicInt32) Get() int32 {
	return atomic.LoadInt32((*int32)(i))
}

type AtomicInt64 int64

func (i *AtomicInt64) Add(n int64) int64 {
	return atomic.AddInt64((*int64)(i), n)
}

func (i *AtomicInt64) Set(n int64) {
	atomic.StoreInt64((*int64)(i), n)
}

func (i *AtomicInt64) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

type AtomicDuration int64

func (d *AtomicDuration) Add(duration time.Duration) time.Duration {
	return time.Duration(atomic.AddInt64((*int64)(d), int64(duration)))
}

func (d *AtomicDuration) Set(duration time.Duration) {
	atomic.StoreInt64((*int64)(d), int64(duration))
}

func (d *AtomicDuration) Get() time.Duration {
	return time.Duration(atomic.LoadInt64((*int64)(d)))
}