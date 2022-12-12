//go:generate goversioninfo -icon=resource/app.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	Init() //初始化
	err := xc.WriteDll(dll)
	if err != nil {
		panic(err)
	}
	// 炫彩_初始化, 参数填true是启用D2D硬件加速, 效果更好. 但xp系统不支持d2d, 这时候你就得填false来关闭d2d了
	ap := app.New(true)

	w := window.NewByLayoutZipMem(zip, "main.xml", "", 0, 0)

	// 从内存加载图片自适应大小
	hIcon = xc.XImage_LoadMemoryAdaptive(icon, 0, 0, 0, 0)
	// 因为下面写信息框还要用, 所以这里禁止图片自动销毁, 这样就可以复用了, 否则用过之后它会自动释放掉的
	xc.XImage_EnableAutoDestroy(hIcon, false)
	// 设置窗口图标
	w.SetIcon(hIcon)

	//获取按钮以及输入框选择框
	btnD := widget.NewButtonByName("BtnD")
	btnW := widget.NewButtonByName("BtnW")
	btnL := widget.NewButtonByName("BtnL")
	input := widget.NewEditByName("Info")
	openResult := widget.NewButtonByName("OpenBtn")
	//注册获取周性能数据被单机事件
	btnW.Event_BnClick(func(pbHandled *bool) int {
		now = time.Now().UnixNano()
		currentTime = strconv.FormatInt(now, 10)
		file_name = "./" + currentTime + "_ResultWeek" + ".csv"
		WriteHead(file_name, dataW) //写入文件头格式
		url_input = input.GetText_Temp()
		if len(url_input) < 10 {
			ap.Alert("提示", "输入url有误,请重新输入")
			return 0
		}
		isMoreUrl = IsMoreUrls(url_input)
		if isMoreUrl {
			realUrls, originUrls := ProcessUrls(url_input)
			for index, item := range realUrls {
				isSuccess = getData(item, originUrls[index], "week")
				if isSuccess == "Success" {
					continue
				} else if isSuccess == "false" {
					ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
				} else {
					ap.Alert("提示", "有输入错误的url,请重新输入")
				}
			}
			// 创建信息框, 本质是一个模态窗口
			hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
			// 设置窗口边框大小
			xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
			// 设置窗口图标
			xc.XWnd_SetIcon(hWindow, hIcon)
			// 显示模态窗口
			xc.XModalWnd_DoModal(hWindow)
			input.SelectAll() //获取成功后删除当前url
			input.DeleteSelect()
		} else {
			realUrl = ProcessUrl(url_input)
			isSuccess = getData(realUrl, url_input, "week")
			if isSuccess == "Success" {
				// 创建信息框, 本质是一个模态窗口
				hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
				// 设置窗口边框大小
				xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
				// 设置窗口图标
				xc.XWnd_SetIcon(hWindow, hIcon)
				// 显示模态窗口
				xc.XModalWnd_DoModal(hWindow)
				input.SelectAll() //获取成功后删除当前url
				input.DeleteSelect()
			} else if isSuccess == "false" {
				ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
			} else {
				ap.Alert("提示", "输入url有误,请重新输入")
			}
		}
		return 0
	})
	//注册获取日常数据按钮被单击事件
	btnD.Event_BnClick(func(pbHandled *bool) int {
		now = time.Now().UnixNano()
		currentTime = strconv.FormatInt(now, 10)
		file_name = "./" + currentTime + "_ResultDaily" + ".csv"
		WriteHead(file_name, dataD) //写入文件
		url_input = input.GetText_Temp()
		if len(url_input) < 10 {
			ap.Alert("提示", "输入url有误,请重新输入")
			return 0
		}
		isMoreUrl = IsMoreUrls(url_input)
		if isMoreUrl {
			realUrls, originUrls := ProcessUrls(url_input)
			for index, item := range realUrls {
				isSuccess = getData(item, originUrls[index], "daily")
				if isSuccess == "Success" {
					continue
				} else if isSuccess == "false" {
					ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
				} else {
					ap.Alert("提示", "有输入错误的url,请重新输入")
				}
			}
			// 创建信息框, 本质是一个模态窗口
			hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
			// 设置窗口边框大小
			xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
			// 设置窗口图标
			xc.XWnd_SetIcon(hWindow, hIcon)
			// 显示模态窗口
			xc.XModalWnd_DoModal(hWindow)
			input.SelectAll() //获取成功后删除当前url
			input.DeleteSelect()
		} else {
			realUrl = ProcessUrl(url_input)
			isSuccess = getData(realUrl, url_input, "daily")
			if isSuccess == "Success" {
				// 创建信息框, 本质是一个模态窗口
				hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
				// 设置窗口边框大小
				xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
				// 设置窗口图标
				xc.XWnd_SetIcon(hWindow, hIcon)
				// 显示模态窗口
				xc.XModalWnd_DoModal(hWindow)
				input.SelectAll() //获取成功后删除当前url
				input.DeleteSelect()
			} else if isSuccess == "false" {
				ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
			} else {
				ap.Alert("提示", "输入url有误,请重新输入")
			}
		}
		return 0
	})
	//注册获取带标签名的按钮被单击事件
	btnL.Event_BnClick(func(pbHandled *bool) int {
		now = time.Now().UnixNano()
		currentTime = strconv.FormatInt(now, 10)
		file_name = "./" + currentTime + "_ResultLabel" + ".csv"
		WriteHead(file_name, dataL) //写入文件
		url_input = input.GetText_Temp()
		if len(url_input) < 10 {
			ap.Alert("提示", "输入url有误,请重新输入")
			return 0
		}
		isMoreUrl = IsMoreUrls(url_input)
		if isMoreUrl {
			realUrls, originUrls := ProcessUrls(url_input)
			for index, item := range realUrls {
				isSuccess = getData(item, originUrls[index], "label")
				if isSuccess == "Success" {
					continue
				} else if isSuccess == "false" {
					ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
				} else {
					ap.Alert("提示", "有输入错误的url,请重新输入")
				}
			}
			// 创建信息框, 本质是一个模态窗口
			hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
			// 设置窗口边框大小
			xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
			// 设置窗口图标
			xc.XWnd_SetIcon(hWindow, hIcon)
			// 显示模态窗口
			xc.XModalWnd_DoModal(hWindow)
			input.SelectAll() //获取成功后删除当前url
			input.DeleteSelect()
		} else {
			realUrl = ProcessUrl(url_input)
			isSuccess = getData(realUrl, url_input, "label")
			if isSuccess == "Success" {
				// 创建信息框, 本质是一个模态窗口
				hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
				// 设置窗口边框大小
				xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
				// 设置窗口图标
				xc.XWnd_SetIcon(hWindow, hIcon)
				// 显示模态窗口
				xc.XModalWnd_DoModal(hWindow)
				input.SelectAll() //获取成功后删除当前url
				input.DeleteSelect()
			} else if isSuccess == "false" {
				ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
			} else {
				ap.Alert("提示", "输入url有误,请重新输入")
			}
		}
		return 0
	})
	//注册打开文件按钮被单机事件
	openResult.Event_BnClick(func(pbHandled *bool) int {
		currentDir, _ := os.Getwd() //获取当前程序的目录
		cmd := exec.Command("explorer", currentDir)
		er := cmd.Start()
		if er != nil { // 运行命令
			ap.Alert("警告", "无法打开当前文件所在目录！"+er.Error())
		}
		return 1
	})
	//调整布局
	w.AdjustLayout()
	// 显示窗口
	w.Show(true)
	// 运行消息循环, 程序会被阻塞在这里不退出, 当炫彩窗口数量为0时退出
	ap.Run()
	// 退出界面库释放资源
	ap.Exit()
}
