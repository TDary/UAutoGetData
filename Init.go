package main

func Init() {
	ContentType = "application/json"
	Useragent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44"
	Cookie = "_ga=GA1.2.116075688.1657612386; Hm_lvt_4bfddcb32e5c5626aa3d10997c3dacd8=1657884787; app_key=e40280a0; Hm_lvt_eefc5ff12060e96822df38857e4cd9ed=1668304669,1669030213,1670317612; project_key=mecha; mysession=MTY3MDMzMDIwM3xOd3dBTkVwVU4weEdRek5FV2xCWFdrWkJTazlMTlVkS1ZsWk9RVUpDV2pOTlFsWkpTRXMwVjBoS1dVVTFSVE5ITmpRMFZWWlhVVUU9fCc72aIc6jEnUHNbO9ifML2kQ0NbHbXrDfDJ5NZLXs7D; email=chenderui1%40thewesthill.net; Hm_lpvt_eefc5ff12060e96822df38857e4cd9ed=1670330209"
	Refer = "http://perfeye.console.testplus.cn/case/list?appKey=mecha"
	//日常版本文件格式初始化
	dataD = []string{
		"案例名", "AvgFPS", "MaxFPS", "MinFPS", "FPS TP90", "Jank(/10min)", "BigJank(/10min)", "RatioFluctuate>10%", "RatioFluctuate>30%", "AvgApp(%)", "MaxApp(%)", "InitMemory(MB)", "AvgMemory(MB)", "PeakMemory(MB)",
		"Avg(GPULoad)[%]", "Max(GPULoad)[%]", "Avg(GPUMemoryUsed)[MB]", "Peak(GPUMemoryUsed)[MB]", "Avg(Drawcall)",
		"Peak(Drawcall)", "Avg(Vertex)", "Peak(Vertex)", "Avg(Primitive)", "Peak(Primitive)", "AvgSend(KB/s)", "MaxSend(KB/s)", "AvgRecv(KB/s)",
		"MaxRecv(KB/s)", "AvgReadBytes(KB/s)", "MaxReadBytes(KB/s)", "AvgWrittenBytes(KB/s)", "MaxWrittenBytes(KB/s)", "基础数据"}
	//周性能版本
	dataW = []string{
		"案例名", "平均FPS", "最低FPS", "90%FPS大于(>=40FPS)", "流畅度", "私有提交内存峰值<=8GB", "显存峰值", "平均CPU占用", "最大CPU占用", "平均GPU占用",
		"最大GPU占用", "平均Draw Call", "最大Draw Call", "平均三角面数(千)", "最大三角面数(千)",
		"平均上传(KB/s)", "上传峰值(KB/s)", "平均下载(KB/s)", "下载峰值(KB/s)", "性能报告地址"}
	//带Label标签版本
	dataL = []string{
		"案例名", "标签名", "AvgFPS", "MaxFPS", "MinFPS", "FPS TP90", "Jank(/10min)", "BigJank(/10min)", "RatioFluctuate>10%", "RatioFluctuate>30%", "AvgApp(%)", "MaxApp(%)", "InitMemory(MB)", "AvgMemory(MB)", "PeakMemory(MB)",
		"Avg(GPULoad)[%]", "Max(GPULoad)[%]", "Avg(GPUMemoryUsed)[MB]", "Peak(GPUMemoryUsed)[MB]", "Avg(Drawcall)",
		"Peak(Drawcall)", "Avg(Vertex)", "Peak(Vertex)", "Avg(Primitive)", "Peak(Primitive)", "AvgSend(KB/s)", "MaxSend(KB/s)", "AvgRecv(KB/s)",
		"MaxRecv(KB/s)", "AvgReadBytes(KB/s)", "MaxReadBytes(KB/s)", "AvgWrittenBytes(KB/s)", "MaxWrittenBytes(KB/s)", "基础数据"}
	perfeyeApi = "http://perfeye.console.testplus.cn/api/show/task/"
}
