package ulog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type wxMessage struct {
	MsgType string        `json:"msgtype"`
	Text    wxTextContent `json:"text"`
}
type wxTextContent struct {
	Content string `json:"content"`
}

var (
	wxClient     = &http.Client{}
	nilLogger    = func(prefix, timestamp, msg string, data ...any) {}
	normalLogger = func(prefix, timestamp, msg string, data ...any) {
		fmt.Printf(prefix+"| "+timestamp+" | "+msg+"\n", data...)
	}

	panicLogger = func(prefix, timestamp, msg string, data ...any) {
		panic(fmt.Sprintf(prefix+"| "+timestamp+" | "+msg+"\n", data...))
	}

	fatalLogger = func(prefix, timestamp, msg string, data ...any) {
		fmt.Printf(prefix+"| "+timestamp+" | "+msg+"\n", data...)
		os.Exit(1)
	}

	//todo bussiness_wx notify
	wxNilLogger = func(wxAddress, prefix, timestamp, msg string, data ...any) {}
	wxLogger    = func(wxAddress, prefix, timestamp, msg string, data ...any) {

		text := fmt.Sprintf(prefix+"| "+timestamp+" | "+msg+"\n", data...)
		// 创建新的HTTP请求
		// 请求体数据
		requestBody := wxMessage{
			MsgType: "text",
			Text:    wxTextContent{Content: text},
		}
		// 将请求体编码成JSON
		requestBodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error marshaling request body:", err)
			return
		}
		req, err := http.NewRequest("POST", wxAddress, bytes.NewBuffer(requestBodyBytes))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		_, err = wxClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		fmt.Printf(prefix+"| "+timestamp+" | "+msg+"\n", data...)
	}
	DefaultLogger = &logger{
		Mutex:      sync.Mutex{},
		timeFormat: "2006-01-02T15:04:05",
		writer:     os.Stdout,
		level:      LogLevelInfo,
		debug:      nilLogger,
		info:       normalLogger,
		warn:       normalLogger,
		error:      normalLogger,
		panic:      panicLogger,
		fatal:      fatalLogger,
		wx:         wxLogger,
		wxAddress:  "",
	}
)

func SetTimeFormat(format string) {
	DefaultLogger.SetTimeFormat(format)
}

func SetLogLevel(level LogLevel) {
	DefaultLogger.SetLogLevel(level)
}

func Debug(msg string, data ...any) {
	DefaultLogger.Debug(msg, data...)
}
func Info(msg string, data ...any) {
	DefaultLogger.Info(msg, data...)
}

func Warn(msg string, data ...any) {
	DefaultLogger.Warn(msg, data...)
}

func Error(msg string, data ...any) {
	DefaultLogger.Error(msg, data...)
}

func Panic(msg string, data ...any) {
	DefaultLogger.Panic(msg, data...)
}

func Fatal(msg string, data ...any) {
	DefaultLogger.Fatal(msg, data...)
}
func WX(msg string, data ...any) {
	DefaultLogger.WX(msg, data...)
}
