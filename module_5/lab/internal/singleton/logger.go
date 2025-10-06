package singleton

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type config struct {
	Level    string `json:"level"`
	FilePath string `json:"file_path"`
}

type Logger struct {
	mu       sync.Mutex
	file     *os.File
	level    LogLevel
	filePath string
}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{level: INFO, filePath: "app.log"}
		_ = instance.openFile()
	})
	return instance
}

func (l *Logger) openFile() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file != nil {
		return nil
	}
	f, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.file = f
	return nil
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil
	return err
}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	l.level = level
	l.mu.Unlock()
}

func (l *Logger) getLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

func (l *Logger) SetLogFilePath(path string) error {
	l.mu.Lock()
	if l.file != nil {
		_ = l.file.Close()
		l.file = nil
	}
	l.filePath = path
	l.mu.Unlock()
	return l.openFile()
}

func (l *Logger) Log(message string, level LogLevel) error {
	if level < l.getLevel() {
		return nil
	}
	ts := time.Now().Format(time.RFC3339)
	line := fmt.Sprintf("%s [%s] %s\n", ts, level.String(), message)
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file == nil {
		if err := l.openFile(); err != nil {
			return err
		}
	}
	_, err := l.file.WriteString(line)
	return err
}

func (l *Logger) LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var c config
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	switch strings.ToUpper(c.Level) {
	case "INFO":
		l.SetLogLevel(INFO)
	case "WARNING":
		l.SetLogLevel(WARNING)
	case "ERROR":
		l.SetLogLevel(ERROR)
	}
	if c.FilePath != "" {
		if err := l.SetLogFilePath(c.FilePath); err != nil {
			return err
		}
	}
	return nil
}

type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

func parseLine(line string) LogEntry {
	parts := strings.SplitN(line, " ", 3)
	var le LogEntry
	if len(parts) >= 3 {
		ts := parts[0]
		lvl := strings.Trim(parts[1], "[]")
		msg := strings.TrimSpace(parts[2])
		t, err := time.Parse(time.RFC3339, ts)
		if err == nil {
			le.Timestamp = t
		} else {
			le.Timestamp = time.Now()
		}
		switch lvl {
		case "INFO":
			le.Level = INFO
		case "WARNING":
			le.Level = WARNING
		case "ERROR":
			le.Level = ERROR
		default:
			le.Level = INFO
		}
		le.Message = msg
	} else {
		le.Timestamp = time.Now()
		le.Level = INFO
		le.Message = line
	}
	return le
}

type LogReader struct {
	path string
}

func NewLogReader(path string) LogReader {
	return LogReader{path: path}
}

func (r LogReader) ReadAll() ([]LogEntry, error) {
	f, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var res []LogEntry
	for sc.Scan() {
		line := sc.Text()
		le := parseLine(line)
		res = append(res, le)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r LogReader) ReadByLevel(minLevel LogLevel) ([]LogEntry, error) {
	all, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	var res []LogEntry
	for _, e := range all {
		if e.Level >= minLevel {
			res = append(res, e)
		}
	}
	return res, nil
}
