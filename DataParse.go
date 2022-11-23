package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//获取数据
func getData(url string, report_url string, isdaily bool) string {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err.Error()
	}
	request.Header.Add("Content-Type", ContentType)
	request.Header.Add("User-Agent", Useragent)
	request.Header.Add("Cookie", Cookie)
	request.Header.Add("Referer", Refer)
	response, err := client.Do(request)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	if isdaily {
		result := ProcessData(string(body), report_url)
		if result != "" {
			return "false"
		}
	} else {
		result := ProcessData2(string(body), report_url)
		if result != "" {
			return "false"
		}
	}
	return "Success"
}

//周性能版本
func ProcessData2(resultData string, report_Url string) string {
	err := json.Unmarshal([]byte(resultData), &DetailMap)
	if err != nil {
		return err.Error()
	}
	casename := DetailMap["data"].(map[string]interface{})["BaseInfo"].(map[string]interface{})["CaseName"]
	DetailFPS := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelFPS"]
	DetailCpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelCPU"]
	DetailGpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelGPU"]
	DetailMemory := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelMemory"]
	DetailRender := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelRenderer"]
	DetailNetWork := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelNetwork"]
	avgfps := DetailFPS.(map[string]interface{})["AvgFPS"]
	minfps := DetailFPS.(map[string]interface{})["MinFPS"]
	tp90fps := DetailFPS.(map[string]interface{})["TP90"]
	smooth := DetailFPS.(map[string]interface{})["Smoothness(%)"]
	avgapp := DetailCpu.(map[string]interface{})["AvgApp(%)"]
	maxapp := DetailCpu.(map[string]interface{})["MaxApp(%)"]
	PeakMemory := DetailMemory.(map[string]interface{})["PeakMemory(MB)"]
	avgGpuLoad := DetailGpu.(map[string]interface{})["Avg(GPULoad)[%]"]
	maxGpuLoad := DetailGpu.(map[string]interface{})["Max(GPULoad)[%]"]
	maxGpuMemry := DetailGpu.(map[string]interface{})["Peak(GPUMemoryUsed)[MB]"]
	avgDrawcall := DetailRender.(map[string]interface{})["Avg(Drawcall)"]
	maxDrawcall := DetailRender.(map[string]interface{})["Peak(Drawcall)"]
	avgPrimitive := DetailRender.(map[string]interface{})["Avg(Primitive)"]
	maxPrimitive := DetailRender.(map[string]interface{})["Peak(Primitive)"]
	avgUpload := DetailNetWork.(map[string]interface{})["AvgSend(KB/s)"]
	maxUpload := DetailNetWork.(map[string]interface{})["MaxSend(KB/s)"]
	avgDownload := DetailNetWork.(map[string]interface{})["AvgRecv(KB/s)"]
	maxDownload := DetailNetWork.(map[string]interface{})["MaxRecv(KB/s)"]
	//获取到的数据，以json形式输出为csv
	resDataW = result2{casename, avgfps, minfps, tp90fps, smooth, PeakMemory, maxGpuMemry, avgapp, maxapp,
		avgGpuLoad, maxGpuLoad, avgDrawcall, maxDrawcall, avgPrimitive, maxPrimitive, avgUpload, maxUpload, avgDownload,
		maxDownload, report_Url}
	WriteData2(resDataW)
	return ""
}

//处理数据日常版本
func ProcessData(resultData string, report_Url string) string {
	err := json.Unmarshal([]byte(resultData), &DetailMap)
	if err != nil {
		return err.Error()
	}
	casename := DetailMap["data"].(map[string]interface{})["BaseInfo"].(map[string]interface{})["CaseName"]
	DetailFPS := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelFPS"]
	DetailCpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelCPU"]
	DetailGpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelGPU"]
	DetailMemory := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelMemory"]
	DetailRender := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelRenderer"]
	DetailNetWork := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelNetwork"]
	DetailIOBytyes := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelIOBytes"]
	avgfps := DetailFPS.(map[string]interface{})["AvgFPS"]
	maxfps := DetailFPS.(map[string]interface{})["MaxFPS"]
	minfps := DetailFPS.(map[string]interface{})["MinFPS"]
	tp90fps := DetailFPS.(map[string]interface{})["TP90"]
	jank := DetailFPS.(map[string]interface{})["Jank(/10min)"]
	bigjank := DetailFPS.(map[string]interface{})["BigJank(/10min)"]
	ratio10 := DetailFPS.(map[string]interface{})["RatioFluctuate\u003e10(%)"]
	ratio30 := DetailFPS.(map[string]interface{})["RatioFluctuate\u003e30(%)"]
	avgapp := DetailCpu.(map[string]interface{})["AvgApp(%)"]
	maxapp := DetailCpu.(map[string]interface{})["MaxApp(%)"]
	InitMemory := DetailMemory.(map[string]interface{})["InitMemory(MB)"]
	AvgMemory := DetailMemory.(map[string]interface{})["AvgMemory(MB)"]
	PeakMemory := DetailMemory.(map[string]interface{})["PeakMemory(MB)"]
	avgGpuLoad := DetailGpu.(map[string]interface{})["Avg(GPULoad)[%]"]
	maxGpuLoad := DetailGpu.(map[string]interface{})["Max(GPULoad)[%]"]
	avgGpuMemry := DetailGpu.(map[string]interface{})["Avg(GPUMemoryUsed)[MB]"]
	maxGpuMemry := DetailGpu.(map[string]interface{})["Peak(GPUMemoryUsed)[MB]"]
	avgDrawcall := DetailRender.(map[string]interface{})["Avg(Drawcall)"]
	maxDrawcall := DetailRender.(map[string]interface{})["Peak(Drawcall)"]
	avgVertex := DetailRender.(map[string]interface{})["Avg(Vertex)"]
	maxVertex := DetailRender.(map[string]interface{})["Peak(Vertex)"]
	avgPrimitive := DetailRender.(map[string]interface{})["Avg(Primitive)"]
	maxPrimitive := DetailRender.(map[string]interface{})["Peak(Primitive)"]
	avgSend := DetailNetWork.(map[string]interface{})["AvgSend(KB/s)"]
	maxSend := DetailNetWork.(map[string]interface{})["MaxSend(KB/s)"]
	avgRecv := DetailNetWork.(map[string]interface{})["AvgRecv(KB/s)"]
	maxRecv := DetailNetWork.(map[string]interface{})["MaxRecv(KB/s)"]
	avgRead := DetailIOBytyes.(map[string]interface{})["AvgReadBytes(KB/s)"]
	maxRead := DetailIOBytyes.(map[string]interface{})["MaxReadBytes(KB/s)"]
	avgWrite := DetailIOBytyes.(map[string]interface{})["AvgWrittenBytes(KB/s)"]
	maxWrite := DetailIOBytyes.(map[string]interface{})["MaxWrittenBytes(KB/s)"]
	//获取到的数据，以json形式输出为csv
	resDataD = result{casename, avgfps, maxfps, minfps, tp90fps, jank, bigjank, ratio10, ratio30, avgapp, maxapp, InitMemory, AvgMemory, PeakMemory,
		avgGpuLoad, maxGpuLoad, avgGpuMemry, maxGpuMemry, avgDrawcall, maxDrawcall, avgVertex, maxVertex, avgPrimitive, maxPrimitive,
		avgSend, maxSend, avgRecv, maxRecv, avgRead, maxRead, avgWrite, maxWrite, report_Url}
	WriteData(resDataD)
	return ""
}
