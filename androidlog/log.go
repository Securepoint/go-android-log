// Package androidlog provides Go bindings for Android's native logging system
package androidlog

/*
#cgo LDFLAGS: -llog

#include <android/log.h>
#include <stdlib.h>

void android_log(int prio, const char* tag, const char* msg) {
    __android_log_write(prio, tag, msg);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// priority levels matching android/log.h
const (
	VERBOSE = C.ANDROID_LOG_VERBOSE
	DEBUG   = C.ANDROID_LOG_DEBUG
	INFO    = C.ANDROID_LOG_INFO
	WARN    = C.ANDROID_LOG_WARN
	ERROR   = C.ANDROID_LOG_ERROR
	FATAL   = C.ANDROID_LOG_FATAL
)

// Logger represents an Android logger instance
type Logger struct {
	tag string
}

// NewLogger creates a new Logger with the specified tag
func NewLogger(tag string) *Logger {
	return &Logger{tag: tag}
}

// log writes a message with the specified priority level
func (l *Logger) log(priority int, msg string) {
	tag := C.CString(l.tag)
	message := C.CString(msg)
	defer C.free(unsafe.Pointer(tag))
	defer C.free(unsafe.Pointer(message))

	C.android_log(C.int(priority), tag, message)
}

// Verbose logs a message with VERBOSE priority
func (l *Logger) Verbose(msg string) {
	l.log(VERBOSE, msg)
}

// Debug logs a message with DEBUG priority
func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}

// Info logs a message with INFO priority
func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

// Warn logs a message with WARN priority
func (l *Logger) Warn(msg string) {
	l.log(WARN, msg)
}

// Error logs a message with ERROR priority
func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

// Fatal logs a message with FATAL priority
func (l *Logger) Fatal(msg string) {
	l.log(FATAL, msg)
}

func (l *Logger) Verbosef(format string, args ...interface{}) {
	l.Verbose(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
