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

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
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
	tag         string
	packageName string
}

// NewLogger creates a new Logger with the specified tag
func NewLogger(tag string) *Logger {
	return &Logger{tag: tag}
}

func (l *Logger) SetPackageName(packageName string) *Logger {
	l.packageName = packageName
	return l
}

// log writes a message with the specified priority level
func (l *Logger) log(priority int, msg string) {
	tag := C.CString(l.tag)
	message := C.CString(msg)
	defer C.free(unsafe.Pointer(tag))
	defer C.free(unsafe.Pointer(message))

	C.android_log(C.int(priority), tag, message)

	db := l.getLogDb()
	if db != nil {
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO log (tag, level, msg) VALUES (?, ?, ?)")
		if err != nil {
			fmt.Println(fmt.Errorf("Error preparing statement: %v", err))
		}
		defer stmt.Close()

		if _, err := stmt.Exec(l.tag, priority, msg); err != nil {
			fmt.Println(fmt.Errorf("Error inserting row: %v", err))
		}
	}
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

// Verbosef logs a formatted message with VERBOSE priority
func (l *Logger) Verbosef(format string, args ...interface{}) {
	l.Verbose(fmt.Sprintf(format, args...))
}

// Debugf logs a formatted message with DEBUG priority
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

// Infof logs a formatted message with INFO priority
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

// Warnf logs a formatted message with WARN priority
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

// Errorf logs a formatted message with ERROR priority
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

// Fatalf logs a formatted message with FATAL priority
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) getLogDb() *sql.DB {
	if l.packageName == "" {
		return nil
	}

	db, err := sql.Open("sqlite3", "/data/data/"+l.packageName+"/databases/log.db")
	if err != nil {
		fmt.Println(fmt.Errorf("Error opening database: %v", err))
		return nil
	}

	return db
}
