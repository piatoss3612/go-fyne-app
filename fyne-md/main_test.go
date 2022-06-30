package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
)

// test if parsing text to markdown properly works
func Test_MakeUI(t *testing.T) {
	var testCfg config

	edit, preview := testCfg.makeUI()

	test.Type(edit, "Hello") // type "Hello" on edit widget

	// if text is not parsed to markdown
	if preview.String() != "Hello" {
		t.Error("Failed -- did not find expected value in preview")
	}
}

// test entire application
func Test_RunApp(t *testing.T) {
	var testCfg config

	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test Markdown")

	edit, preview := testCfg.makeUI()

	testCfg.createMenuItems(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))

	testApp.Run()

	test.Type(edit, "Some Text")

	if preview.String() != "Some Text" {
		t.Error("failed")
	}
}
