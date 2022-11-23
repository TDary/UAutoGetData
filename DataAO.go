package main

import (
	_ "embed"
	"net/http"
	"time"
)

//go:embed resource/Reptitle.zip
var zip []byte

//go:embed resource/xcgui.dll
var dll []byte

//go:embed resource/Title.ico
var icon []byte

//窗口图标句柄设置
var hIcon = 0

var apiUrl string
var uuid string
var listUrl []string
var url_input string
var everyUrl []string
var apiurls []string
var originUrls []string
var oriUrl string
var isSuccess string
var realUrl string
var isMoreUrl bool
var file_name string
var now int64
var currentTime string
var ContentType string
var Useragent string
var Cookie string
var Refer string
var DetailMap = make(map[string]interface{})
var client = http.Client{
	Timeout: 10 * time.Second,
}
var dataD []string
var dataW []string
var perfeyeApi string
var resDataD result
var resDataW result2

type result2 struct {
	CaseName        interface{}
	AvgFPS          interface{}
	MinFPS          interface{}
	FPSTP90         interface{}
	SmoothPlay      interface{}
	PeakMemory      interface{}
	MaxGPUMemory    interface{}
	AvgApp          interface{}
	MaxApp          interface{}
	AvgGPULoad      interface{}
	MaxGPULoad      interface{}
	AvgDrawcall     interface{}
	PeakDrawcall    interface{}
	AvgTriAngles    interface{}
	MaxTriAngles    interface{}
	AverageUpload   interface{}
	MaxUpload       interface{}
	AverageDownLoad interface{}
	MaxDownLoad     interface{}
	BasicData       string
}

type result struct {
	CaseName         interface{}
	AvgFPS           interface{}
	MaxFPS           interface{}
	MinFPS           interface{}
	FPSTP90          interface{}
	Jank             interface{}
	BigJank          interface{}
	RatioFluctuate10 interface{}
	RatioFluctuate30 interface{}
	AvgApp           interface{}
	MaxApp           interface{}
	InitMemory       interface{}
	AvgMemory        interface{}
	PeakMemory       interface{}
	AvgGPULoad       interface{}
	MaxGPULoad       interface{}
	AvgGPUMemory     interface{}
	MaxGPUMemory     interface{}
	AvgDrawcall      interface{}
	PeakDrawcall     interface{}
	AvgVertex        interface{}
	PeakVertex       interface{}
	AvgPrimitive     interface{}
	PeakPrimitive    interface{}
	AvgSend          interface{}
	MaxSend          interface{}
	AvgRecv          interface{}
	MaxRecv          interface{}
	AvgRead          interface{}
	MaxRead          interface{}
	AvgWrite         interface{}
	MaxWrite         interface{}
	BasicData        string
}
