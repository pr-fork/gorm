package logger

import "testing"

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
