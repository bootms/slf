// Copyright (c) 2016 Ventu.io, Oleg Sklyar, contributors
// The use of this source code is governed by a MIT style license found in the LICENSE file

package slf_test

import (
	"github.com/bootms/slf"
	"testing"
)

func TestChaining_withContext_success(t *testing.T) {
	log := slf.WithContext("context")
	log.Log(slf.LevelInfo, "").Trace(nil)
	log.WithError(nil).Debugf("%v", "").Trace(nil)
	log.WithField("a", "b").Info("").Trace(nil)
	log.WithFields(slf.Fields{"a": "b"}).Error("").Trace(nil)
	log.WithFields(slf.Fields{"a": "b"}).WithError(nil).Infof("%v", "").Trace(nil)
	log.WithError(nil).Errorf("%v", "").Trace(nil)
	log.WithField("a", "b").WithField("c", "d").Info("").Trace(nil)
	log.Debug("")
	log.WithField("a", "b").Debug("").Trace(nil)
	log.WithError(nil).Warn("").Trace(nil)
	log.WithError(nil).Warnf("%v", "").Trace(nil)
	log.WithCaller(slf.CallerShort).WithCaller(slf.CallerLong).WithCaller(slf.CallerNone).Error("test")
}

func TestChaning_panic_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	slf.WithContext("context").Panic("")
}

func TestChaning_log_levelpanic_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	slf.WithContext("context").Log(slf.LevelPanic, "")
}

func TestChaning_panicf_success(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expecte")
		}
	}()
	slf.WithContext("context").WithField("a", "b").Panicf("%v", "")
}

func TestExitProcessor_fatalf_success(t *testing.T) {
	triggered := false
	slf.ExitProcessor = func() {
		triggered = true
	}
	slf.WithContext("context").Fatalf("%v", "")
	if !triggered {
		t.Error("Exit not triggered")
	}
}

func TestExitProcessor_fatal_success(t *testing.T) {
	triggered := false
	slf.ExitProcessor = func() {
		triggered = true
	}
	slf.WithContext("context").Fatal("a")
	if !triggered {
		t.Error("Exit not triggered")
	}
}

func TestExitProcessor_log_levelfatal_success(t *testing.T) {
	triggered := false
	slf.ExitProcessor = func() {
		triggered = true
	}
	slf.WithContext("context").Log(slf.LevelFatal, "a")
	if !triggered {
		t.Error("Exit not triggered")
	}
}
