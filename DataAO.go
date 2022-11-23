package main

import (
	_ "embed"
)

//go:embed resource/Reptitle.zip
var zip []byte

//go:embed resource/xcgui.dll
var dll []byte

//go:embed resource/Title.ico
var icon []byte

//窗口图标句柄设置
var hIcon = 0

var url_input string
var isSuccess string
var realUrl string
var isMoreUrl bool
var file_name string
var now int64
var currentTime string
