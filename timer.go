// Copyright 2010 Sonia Keys
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package timer offers a simple way to time a Go function.
//
// Typically, you add
//
//     defer timer.From(timer.Now())
//
// as the first line of your function, and this will print an elapsed time
// message when the function returns.
package timer

import (
	"fmt"
	"runtime"
	"time"
)

// Now samples and returns the current time in nanoseconds.
func Now() int64 {
	return time.Nanoseconds()
}

// From prints an elapsed time message.
//
// It prints the the name of the function and the time elapsed between
// two times:  The time passed as the argument, and the current time.
func From(t0 int64) {
	dt := float64(Now() - t0)
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s ",runtime.FuncForPC(pc).Name())
	}
	if dt > 1e9 {
		fmt.Printf("%.3f *seconds* elapsed.\n", dt*1e-9)
	} else {
		fmt.Printf("%.3f ms elapsed.\n", dt*1e-6)
	}
}
