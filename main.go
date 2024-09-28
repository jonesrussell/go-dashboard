package main

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

func main() {
    app := tview.NewApplication()

    // Create a new Flex layout
    flex := tview.NewFlex().
        AddItem(tview.NewBox().SetBorder(true).SetTitle("Left"), 0, 1, false).
        AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle"), 0, 2, false).
        AddItem(tview.NewBox().SetBorder(true).SetTitle("Right"), 0, 1, false)

    // Create a form for configuration
    form := tview.NewForm().
        AddInputField("Title", "", 20, nil, nil).
        AddButton("Save", func() {
            // Handle save action
        }).
        AddButton("Quit", func() {
            app.Stop()
        })

    // Create a new Flex layout for the dashboard and form
    dashboard := tview.NewFlex().
        SetDirection(tview.FlexRow).
        AddItem(flex, 0, 1, true).
        AddItem(form, 0, 1, true)

    if err := app.SetRoot(dashboard, true).Run(); err != nil {
        panic(err)
    }
}

