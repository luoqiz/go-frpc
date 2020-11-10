package service

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"go-frpc/frp"
	"gopkg.in/ini.v1"
)

func ServiceScreen(_ fyne.Window) fyne.CanvasObject {

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

	tabs.OnChanged = func(t *container.TabItem) {
		println(t.Text)
		t.Content = loadService(t.Text)
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		tabs,
	)

}

func loadService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return container.NewVBox(saveButton, content)
}
