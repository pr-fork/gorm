package logger

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var ctx = context.Background()

type testWriter struct {
	result string
}

func newTestWriter() *testWriter {
	return &testWriter{}
}

func (t *testWriter) Printf(format string, args ...interface{}) {
	t.result += fmt.Sprintf(format, args...) + "\n"
}

func (t *testWriter) Result() string {
	// defer func() {
	// 	t.result = ""
	// }()

	return t.result
}

func TestLogger_Level(t *testing.T) {
	tests := []struct {
		name  string
		level int
		want  LogLevel
	}{
		{
			name:  "Silent",
			level: 1,
			want:  Silent,
		},
		{
			name:  "Error",
			level: 2,
			want:  Error,
		},
		{
			name:  "Warn",
			level: 3,
			want:  Warn,
		},
		{
			name:  "Info",
			level: 4,
			want:  Info,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogLevel(tt.level); got != tt.want {
				t.Errorf("LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestLogger_New(t *testing.T) {
	writer := newTestWriter()

	logger := New(writer, Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	if logger == nil {
		t.Errorf("Default logger is nil")
	}

	logger.Info(ctx, "Info")
	logger.Warn(ctx, "Warn")
	logger.Error(ctx, "Error")
	logger.Trace(ctx, time.Now(), nil, nil)
	fmt.Println(writer.Result())

	// logger.Warn(ctx, "test")
	// logger.Error(ctx, "test")
	// logger.Trace(ctx, time.Now(), nil, nil)
}
