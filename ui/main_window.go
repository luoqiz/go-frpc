package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go-frpc/frp"
	"gopkg.in/ini.v1"
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
			//fyne.NewMenuItem("download", func() {
			//}),
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
			}),
		)),
	)
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("status", theme.HomeIcon(), statusScreen(a)),
		widget.NewTabItemWithIcon("service", theme.ContentCopyIcon(), ServiceScreen()),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), AdvancedScreen(w)),
		widget.NewTabItemWithIcon("frpc", theme.ConfirmIcon(), InfoScreen(w)))
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	tabs.OnChanged = func(tab *widget.TabItem) {
		if tab.Text == "service" {
			tab.Content = ServiceScreen()
		}
	}

	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
func InfoScreen(w fyne.Window) fyne.CanvasObject {
	// 文本框
	labelStr := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	// 保存按钮
	downloadButton := widget.NewButton("download", func() {
		frp.Download(func(length, downLen int64) {
			labelStr.SetText(fmt.Sprintln("process: ", int(downLen)*100/int(length), "%"))
			fmt.Println(length, downLen, float32(downLen)/float32(length))
		})
		labelStr.SetText("download complete!")
	})
	return widget.NewVBox(downloadButton, labelStr)
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
		fc, err := frp.FullContent()
		if err != nil {
			content.SetText("not fount the file : frpc.ini")
		} else {
			content.SetText(fc)
		}

	})
	//content.SetText(frp.FullContent())
	return widget.NewVBox(saveButton, reloadButton, content)
}

func CommonService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
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

func loadService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return widget.NewVBox(saveButton, content)
}

func ServiceScreen() fyne.CanvasObject {

	//tabs := widget.NewTabContainer(
	//	widget.NewTabItem("common", CommonService("common")),
	//	widget.NewTabItem("http", WebService("web")),
	//	widget.NewTabItem("mstsc", MstscService("mstsc")),
	//)
	tabs := widget.NewTabContainer()
	sections, err := frp.GetSections()
	if err != nil {
		return widget.NewLabel("please download frp...")
	}
	for _, section := range sections {
		if section == ini.DefaultSection {
			continue
		}
		tabs.Append(widget.NewTabItem(section, loadService(section)))
	}

	tabs.OnChanged = func(t *widget.TabItem) {
		println(t.Text)
		t.Content = loadService(t.Text)
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		tabs,
	)

}

func statusScreen(a fyne.App) fyne.CanvasObject {

	pid := frp.CheckStatus()
	showText := ""
	if pid != 0 {
		showText = fmt.Sprintln("running ... pid:", pid)
	} else {
		showText = fmt.Sprintln("frp not runing...")
	}
	statusShow := widget.NewLabelWithStyle(showText, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	return widget.NewVBox(

		widget.NewGroup("status",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("start", func() {
					pid := frp.Start()
					if pid != 0 {
						statusShow.SetText(fmt.Sprintln("running pid: ", pid))
					} else {
						statusShow.SetText(fmt.Sprintln("please download frp... ", pid))
					}
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
