package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go-frpc/src/frp"
	"gopkg.in/ini.v1"
)

func ServiceScreen(_ fyne.Window) fyne.CanvasObject {

	//tabs := widget.NewTabContainer(
	//	widget.NewTabItem("common", CommonService("common")),
	//	widget.NewTabItem("http", WebService("web")),
	//	widget.NewTabItem("mstsc", MstscService("mstsc")),
	//)
	tabs := container.NewAppTabs()
	sections, err := frp.GetSections()
	if err != nil {
		return widget.NewLabel("please download frp...")
	}
	for _, section := range sections {
		if section == ini.DefaultSection {
			continue
		}
		tabs.Append(container.NewTabItem(section, loadService(section)))
	}

	tabs.OnChanged = func(t *container.TabItem) {
		println(t.Text)
		t.Content = loadService(t.Text)
	}
	return container.New(layout.NewBorderLayout(nil, nil, nil, nil),
		tabs,
	)

}

func loadService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	scroll := container.NewScroll(content)
	scroll.SetMinSize(fyne.NewSize(400, 500))
	return container.NewVBox(saveButton, container.NewAdaptiveGrid(1, scroll))
}
