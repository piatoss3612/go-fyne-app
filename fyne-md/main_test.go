package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
)

func Test_MakeUI(t *testing.T) {
	var testCfg config

	edit, preview := testCfg.makeUI()

	test.Type(edit, "Hello")

	if preview.String() != "Hello" {
		t.Error("Failed -- did not find expected value in preview")
	}
}

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
