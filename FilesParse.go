package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

//写入文件初始化格式
func WriteHead(file_name string, data []string) {
	f, err := os.Create(file_name) //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流
	w.Write(data)
	w.Flush()
}

//写入数据
func WriteData2(res result2) {
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流
	AvgTris := res.AvgTriAngles.(string)
	MaxTris := res.MaxTriAngles.(string)

	AVGtrian, err := strconv.ParseFloat(AvgTris, 32)
	if err != nil {
		panic(err)
	}
	MAXtrian, err := strconv.ParseFloat(MaxTris, 32)
	if err != nil {
		panic(err)
	}
	AVGtrian = AVGtrian / 1000
	MAXtrian = MAXtrian / 1000

	resData := []string{res.CaseName.(string), res.AvgFPS.(string), res.MinFPS.(string), res.FPSTP90.(string),
		res.SmoothPlay.(string), res.PeakMemory.(string), res.MaxGPUMemory.(string), res.AvgApp.(string), res.MaxApp.(string), res.AvgGPULoad.(string),
		res.MaxGPULoad.(string), res.AvgDrawcall.(string), res.PeakDrawcall.(string), strconv.FormatFloat(AVGtrian, 'E', -1, 32), strconv.FormatFloat(MAXtrian, 'E', -1, 32), res.AverageUpload.(string),
		res.MaxUpload.(string), res.AverageDownLoad.(string), res.MaxDownLoad.(string), res.BasicData}
	w.Write(resData)
	w.Flush()
}

//写入数据
func WriteData(res result) {
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流

	resData := []string{res.CaseName.(string), res.AvgFPS.(string), res.MaxFPS.(string), res.MinFPS.(string),
		res.FPSTP90.(string), res.Jank.(string), res.BigJank.(string), res.RatioFluctuate10.(string), res.RatioFluctuate30.(string), res.AvgApp.(string), res.MaxApp.(string), res.InitMemory.(string), res.AvgMemory.(string),
		res.PeakMemory.(string), res.AvgGPULoad.(string), res.MaxGPULoad.(string), res.AvgGPUMemory.(string), res.MaxGPUMemory.(string),
		res.AvgDrawcall.(string), res.PeakDrawcall.(string), res.AvgVertex.(string), res.PeakVertex.(string), res.AvgPrimitive.(string),
		res.PeakPrimitive.(string), res.AvgSend.(string), res.MaxSend.(string), res.AvgRecv.(string), res.MaxRecv.(string), res.AvgRead.(string),
		res.MaxRead.(string), res.AvgWrite.(string), res.MaxWrite.(string), res.BasicData}
	w.Write(resData)
	w.Flush()
}

//写入带label的数据
func WriteData3(res result3) {
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流

	resData := []string{res.CaseName.(string), res.LabelName.(string), res.AvgFPS.(string), res.MaxFPS.(string), res.MinFPS.(string),
		res.FPSTP90.(string), res.Jank.(string), res.BigJank.(string), res.RatioFluctuate10.(string), res.RatioFluctuate30.(string), res.AvgApp.(string), res.MaxApp.(string), res.InitMemory.(string), res.AvgMemory.(string),
		res.PeakMemory.(string), res.AvgGPULoad.(string), res.MaxGPULoad.(string), res.AvgGPUMemory.(string), res.MaxGPUMemory.(string),
		res.AvgDrawcall.(string), res.PeakDrawcall.(string), res.AvgVertex.(string), res.PeakVertex.(string), res.AvgPrimitive.(string),
		res.PeakPrimitive.(string), res.AvgSend.(string), res.MaxSend.(string), res.AvgRecv.(string), res.MaxRecv.(string), res.AvgRead.(string),
		res.MaxRead.(string), res.AvgWrite.(string), res.MaxWrite.(string), res.BasicData}
	w.Write(resData)
	w.Flush()
}
