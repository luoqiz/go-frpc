package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go-frpc/src/frp"
	"go-frpc/src/ui/nav"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func MainWindow() {
	a := app.NewWithID("top.luoqiz.go-frpc")
	a.SetIcon(theme.FyneLogo())

	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("frpc client")
	topWindow = w

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
	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	setNavItem := func(t nav.Item) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}

		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(makeNav(setNavItem, false))
	} else {
		split := container.NewHSplit(makeNav(setNavItem, true), tutorial)
		split.Offset = 0.2
		w.SetContent(split)
	}
	w.ShowAndRun()
}

func makeNav(setNavItem func(navItem nav.Item), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return nav.NavIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := nav.NavIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := nav.NavItems[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := nav.NavItems[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setNavItem(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

func CommonService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return container.NewVBox(saveButton, content)
}

func WebService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return container.NewVBox(saveButton, content)
}

func MstscService(section string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("save", func() {
		frp.SaveSection(section, content.Text)
	})
	content.SetText(frp.GetSection(section))
	return container.NewVBox(saveButton, content)
}
