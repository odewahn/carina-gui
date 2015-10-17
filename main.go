package main

import (
	"log"
	"os"

	"github.com/andlabs/ui"
)

var (
	w ui.Window
)

func gui() {

	//Define credentials area
	usernameLabel := ui.NewLabel("Username:")
	usernameTextField := ui.NewTextField()
	if len(os.Getenv("RACKSPACE_USERNAME")) > 0 {
		usernameTextField.SetText(os.Getenv("RACKSPACE_USERNAME"))
	}
	apiKeyLabel := ui.NewLabel("API Key:")
	apiKeyTextField := ui.NewPasswordField()
	if len(os.Getenv("RACKSPACE_APIKEY")) > 0 {
		apiKeyTextField.SetText(os.Getenv("RACKSPACE_APIKEY"))
	}
	connectBtn := ui.NewButton("Connect")

	// layout the login controls on a grid
	loginGrid := ui.NewGrid()
	loginGrid.Add(usernameLabel, nil, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.Add(usernameTextField, usernameLabel, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	loginGrid.Add(apiKeyLabel, usernameLabel, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.Add(apiKeyTextField, apiKeyLabel, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	loginGrid.Add(connectBtn, nil, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.SetPadded(true)

	//Define main interface
	loginGrp := ui.NewGroup("Rackspace Credentials", loginGrid)
	loginGrp.SetMargined(true)
	clusterGrp := ui.NewGroup("Cluster Information", ui.Space())

	mainGrid := ui.NewGrid()
	mainGrid.Add(loginGrid, nil, ui.East, true, ui.Fill, false, ui.Center, 1, 1)
	mainGrid.Add(clusterGrp, loginGrid, ui.South, true, ui.Fill, false, ui.Center, 1, 4)
	//mainGrid.Add(ui.Space(), clusterGrp, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	//mainGrid.Add(ui.Space(), clusterGrp, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	//mainGrid.Add(ui.Space(), clusterGrp, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	mainGrid.SetPadded(true)

	//mainStack := ui.NewVerticalStack(loginGrp, clusterGrp)
	//mainStack.SetStretchy(0)
	//mainStack.SetStretchy(1)

	//Main stack of the interfaces
	w = ui.NewWindow("Carina by Rackspace GUI Client", 600, 450, mainGrid)
	w.SetMargined(true)

	w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	w.Show()

}

/*
func updateTable(table ui.Table) {
	for {
		table.Lock()
		d := table.Data().(*[]Container)
		*d = running
		table.Unlock()
		time.Sleep(1 * time.Second)
	}
}
*/

func main() {

	// This runs the code that displays our GUI.
	// All code that interfaces with package ui (except event handlers) must be run from within a ui.Do() call.
	go ui.Do(gui)

	err := ui.Go()
	if err != nil {
		log.Print(err)
	}

}
