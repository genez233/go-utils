package glog

import (
	"encoding/json"
	"fmt"
	"github.com/genez233/go-utils/dc"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var (
	//logUrl  = "http://logs.zhiyunai.com.cn/api/default/%s/_json"
	_config *Config // 配置项
)

type GLog struct {
}

// Config 配置项
type Config struct {
	ServerName       string // 应用名称(openobserve 仓库名)
	Version          string // 应用版本
	ConsoleLog       bool   // 是否打印到控制台
	IsUpload         bool   // 是否上传
	RunMode          string // 运行环境
	LogUrl           string // 上传地址
	OpenobserveToken string // Openobserve的Token
}

// Message Open-observe 实体
type Message struct {
	Project        string `json:"project"`
	ProjectVersion string `json:"project_version"`
	Content        string `json:"content"`
	Level          string `json:"level"`
	IP             string `json:"ip"`
	DateTime       string `json:"datetime"`
	GOOS           string `json:"goos"`
	GOARCH         string `json:"goarch"`
	RunMode        string `json:"run_mode"`
}

// New 创建 zlog 实例
func New(config *Config) *GLog {
	if config == nil {
		log.Fatal("glog must has config.")
	}

	_config = config
	if _config.ServerName == "" {
		_config.ServerName = "default"
	}

	_config.LogUrl = fmt.Sprintf(_config.LogUrl, _config.ServerName)
	return &GLog{}
}

func (z *GLog) Info(a ...interface{}) {
	msg := z.getAnyString(a)
	if _config.ConsoleLog {
		printf("info", msg)
	}
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:             ip,
		Content:        msg,
		Project:        _config.ServerName,
		ProjectVersion: _config.Version,
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
		RunMode:        _config.RunMode,
		Level:          "info",
		DateTime:       time.Now().Format("2006-01-02 15:04:05"),
	}
	arr = append(arr, param)

	if _config.IsUpload {
		post(_config.LogUrl, arr)
	}
}

func (z *GLog) Debug(a ...interface{}) {
	msg := z.getAnyString(a)
	if _config.ConsoleLog {
		printf("debug", msg)
	}
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:             ip,
		Content:        msg,
		Project:        _config.ServerName,
		ProjectVersion: _config.Version,
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
		RunMode:        _config.RunMode,
		Level:          "debug",
		DateTime:       time.Now().Format("2006-01-02 15:04:05"),
	}
	arr = append(arr, param)

	if _config.IsUpload {
		post(_config.LogUrl, arr)
	}
}

func (z *GLog) Warn(a ...interface{}) {
	msg := z.getAnyString(a)
	if _config.ConsoleLog {
		printf("warn", msg)
	}
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:             ip,
		Content:        msg,
		Project:        _config.ServerName,
		ProjectVersion: _config.Version,
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
		RunMode:        _config.RunMode,
		Level:          "warn",
		DateTime:       time.Now().Format("2006-01-02 15:04:05"),
	}
	arr = append(arr, param)

	if _config.IsUpload {
		post(_config.LogUrl, arr)
	}
}

func (z *GLog) Error(a ...interface{}) {
	msg := z.getAnyString(a)
	if _config.ConsoleLog {
		printf("error", msg)
	}
	if !_config.IsUpload {
		return
	}
	ip := dc.GetIP()
	arr := make([]Message, 0)
	param := Message{
		IP:             ip,
		Content:        msg,
		Project:        _config.ServerName,
		ProjectVersion: _config.Version,
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
		RunMode:        _config.RunMode,
		Level:          "error",
		DateTime:       time.Now().Format("2006-01-02 15:04:05"),
	}
	arr = append(arr, param)

	if _config.IsUpload {
		post(_config.LogUrl, arr)
	}
}

func (z *GLog) getAnyString(a ...interface{}) string {
	str := ""

	for _, arg := range a {
		for i2, arg2 := range arg.([]interface{}) {
			if i2 > 0 {
				str += " "
			}

			str += fmt.Sprintf("%v", arg2)
		}
	}

	return str
}

// post HttpPost
func post(url string, data interface{}) {
	method := "POST"
	jsonBytes, _ := json.Marshal(data)
	//fmt.Println(string(jsonBytes))
	payload := strings.NewReader(string(jsonBytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", _config.OpenobserveToken)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "logs.zhiyunai.com.cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("organization", "default")
	req.Header.Add("stream-name", "default")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))

	return
}

func printf(a ...interface{}) {
	if _config.ConsoleLog {
		fmt.Println(a)
	}
}
