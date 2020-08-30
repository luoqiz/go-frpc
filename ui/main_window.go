package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go-frpc/frp"
	"go-frpc/utils"
	"path/filepath"
)

const preferenceCurrentTab = "currentTab"

func MainWindow() {
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(theme.FyneLogo())

	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("frpc client")
	w.Resize(fyne.NewSize(800, 600))
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("frp",
			fyne.NewMenuItem("download", func() {
				frp.Download()
			}),
			fyne.NewMenuItem("start", func() {
				frp.Start()
			}),
			fyne.NewMenuItem("stop", func() {
				frp.Stop()
			}),
		),
		fyne.NewMenu("help",
			fyne.NewMenuItem("go-frpc", func() { fmt.Println("go-frpc") }),
			fyne.NewMenuItem("frp", func() { fmt.Println("frp") }),
			fyne.NewMenuItem("fyne", func() {
				fmt.Println("fyne")
				println("frp download start...")
				workDir, _ := filepath.Abs("")
				utils.UnZip(workDir+"/frp.zip", workDir+"/frpcc", true)
			}),
		)),
	)
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("status", theme.HomeIcon(), statusScreen(a)),
		widget.NewTabItemWithIcon("service", theme.ContentCopyIcon(), SettingScreen()),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), AdvancedScreen(w)))
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
func AdvancedScreen(w fyne.Window) fyne.CanvasObject {
	// 文本框
	content := widget.NewMultiLineEntry()
	// 保存按钮
	saveButton := widget.NewButton("save", func() {
		frp.SetContent(content.Text)
	})
	// 重新加载按钮
	reloadButton := widget.NewButton("reload", func() {
		content.SetText(frp.FullContent())
		// 发送消息
		utils.SendNotifiction("成功加载最新文件内容")
	})
	content.SetText(frp.FullContent())
	return widget.NewVBox(saveButton, reloadButton, content)
}

func CommonService() fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection("common", content.Text)
	})
	content.SetText(frp.GetSection("common"))
	return widget.NewVBox(saveButton, content)
}

func WebService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return widget.NewVBox(saveButton, content)
}

func MstscService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return widget.NewVBox(saveButton, content)
}

func SettingScreen() fyne.CanvasObject {

	tabs := widget.NewTabContainer(
		widget.NewTabItem("common", CommonService()),
		widget.NewTabItem("http", WebService("web")),
		widget.NewTabItem("mstsc", MstscService("mstsc")),
	)
	tabs.OnChanged = func(t *widget.TabItem) {
		println(t.Content)
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		tabs,
	)

}

func statusScreen(a fyne.App) fyne.CanvasObject {

	pid := frp.CheckStatus()
	showText := ""
	if pid != 0 {
		showText = fmt.Sprintln("running ... pid: %v", pid)
	} else {
		showText = fmt.Sprintln("frp not runing...")
	}
	statusShow := widget.NewLabelWithStyle(showText, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	return widget.NewVBox(

		widget.NewGroup("status",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("start", func() {
					frp.Start()
					pid := frp.CheckStatus()
					statusShow.SetText(fmt.Sprintln("running pid: ", pid))
				}),
				widget.NewButton("stop", func() {
					if frp.CheckStatus() != 0 {
						frp.Stop()
						statusShow.SetText("frp success closed...")
					} else {
						statusShow.SetText("frp not running...")
					}

				}),
			),
		),
		statusShow,
	)
}
