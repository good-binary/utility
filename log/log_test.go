package customlog

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name    string
		opts    *LogOptions
		wantErr bool
	}{
		{
			name: "valid options",
			opts: &LogOptions{
				ToStdout: true,
				Level:    Info,
			},
			wantErr: false,
		},
		{
			name:    "nil options",
			opts:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLogger(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Logger_log(t *testing.T) {
	tests := []struct {
		name    string
		l       *Logger
		level   LogLevel
		msg     string
		data    interface{}
		wantErr bool
	}{
		{
			name: "debug level",
			l: &Logger{
				options: &LogOptions{
					Level: Debug,
				},
			},
			level:   Debug,
			msg:     "test message",
			data:    nil,
			wantErr: false,
		},
		{
			name: "info level",
			l: &Logger{
				options: &LogOptions{
					Level: Info,
				},
			},
			level:   Info,
			msg:     "test message",
			data:    nil,
			wantErr: false,
		},
		{
			name: "warning level",
			l: &Logger{
				options: &LogOptions{
					Level: Warning,
				},
			},
			level:   Warning,
			msg:     "test message",
			data:    nil,
			wantErr: false,
		},
		{
			name: "error level",
			l: &Logger{
				options: &LogOptions{
					Level: Error,
				},
			},
			level:   Error,
			msg:     "test message",
			data:    nil,
			wantErr: false,
		},
		{
			name: "disabled level",
			l: &Logger{
				options: &LogOptions{
					Level: Debug,
				},
			},
			level:   Debug - 1,
			msg:     "test message",
			data:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.log(tt.level, tt.msg, tt.data)
			if tt.wantErr && tt.l.options.ToFile {
				t.Errorf("log() error = %v, wantErr %v", nil, tt.wantErr)
			}
		})
	}
}

// func Test_Logger_Debug(t *testing.T) {
//     tests := []struct {
//         name    string
//         l       *Logger
//         msg     string
//         data    interface{}
//         wantErr bool
//     }{
//         {
//             name: "debug level",
//             l: &Logger{
//                 options: &LogOptions{
//                     Level: Debug,
//                 },
//             },
//             msg:     "test message",
//             data:    nil,
//             wantErr: false,
//         },
//         {
//             name: "info level",
//             l: &Logger{
//                 options: &LogOptions{
//                     Level: Info,
//                 },
//             },
//             msg:     "test message",
//             data:    nil,
//             wantErr: false,
//         },
//         {
//             name: "warning level",
//             l: &Logger{
//                 options: &LogOptions{
//                     Level: Warning,
//                 },
//             },
//             msg:     "test message",
//             data:    nil,
//             wantErr: false,
//         },
//         {
//             name: "error level",
//             l: &Logger{
//                 options: &LogOptions{
//                     Level: Error,
//                 },
//             },
//             msg:     "test message",
//             data:    nil,
//             wantErr: false,
//         },
//         {
//             name: "disabled level",
//             l: &Logger{
//                 options: &LogOptions{
//                     Level: Debug,
//                 },
//             },
//             msg:     "test message",
//             data:    nil,
//             wantErr: true,
//         },
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.
