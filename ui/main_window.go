package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/cmd/fyne_demo/data"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go-frpc/utils"
)

const preferenceCurrentTab = "currentTab"

func MainWindow() {
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(theme.FyneLogo())

	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("Fyne Demo")
	//w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("file",
	//	fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
	//	// a quit item will be appended to our first menu
	//), fyne.NewMenu("Edit",
	//	fyne.NewMenuItem("Cut", func() { fmt.Println("Menu Cut") }),
	//	fyne.NewMenuItem("Copy", func() { fmt.Println("Menu Copy") }),
	//	fyne.NewMenuItem("Paste", func() { fmt.Println("Menu Paste") }),
	//)))
	//w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcomeScreen(a)),
		widget.NewTabItemWithIcon("Widgets", theme.ContentCopyIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Graphics", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Windows", theme.ViewFullScreenIcon(), screens.DialogScreen(w)),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), screens.AdvancedScreen(w)))
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}

func welcomeScreen(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	logo.SetMinSize(fyne.NewSize(228, 167))

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),

		widget.NewHBox(layout.NewSpacer(),
			widget.NewHyperlink("fyne.io", utils.ParseURL("https://fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", utils.ParseURL("https://fyne.io/develop/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("sponsor", utils.ParseURL("https://github.com/sponsors/fyne-io")),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}
