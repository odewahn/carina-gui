package main

import (
	"log"
	"os"
	"reflect"
	"time"

	"github.com/andlabs/ui"
	"github.com/rackerlabs/libcarina"
)

var (
	w            ui.Window
	carinaClient *libcarina.ClusterClient
)

func gui() {

	//Define endpoint
	apiEndpointLabel := ui.NewLabel("API Endpoint:")
	apiEndpointTextField := ui.NewTextField()
	if len(os.Getenv("RACKSPACE_API_ENDPOINT")) > 0 {
		apiEndpointTextField.SetText(os.Getenv("RACKSPACE_API_ENDPOINT"))
	} else {
		apiEndpointTextField.SetText(libcarina.BetaEndpoint)
	}
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
	loginGrid.Add(apiEndpointLabel, nil, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.Add(apiEndpointTextField, apiEndpointLabel, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	loginGrid.Add(usernameLabel, apiEndpointLabel, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.Add(usernameTextField, usernameLabel, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	loginGrid.Add(apiKeyLabel, usernameLabel, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.Add(apiKeyTextField, apiKeyLabel, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	loginGrid.Add(connectBtn, nil, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	loginGrid.SetPadded(true)

	// Define the table that lists all running clusters
	var c libcarina.Cluster
	clusterListTable := ui.NewTable(reflect.TypeOf(c))

	mainGrid := ui.NewGrid()
	mainGrid.Add(loginGrid, nil, ui.East, true, ui.Fill, false, ui.Center, 1, 1)
	mainGrid.Add(clusterListTable, loginGrid, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	mainGrid.SetPadded(true)

	connectBtn.OnClicked(func() {
		connect(apiEndpointTextField.Text(), usernameTextField.Text(), apiKeyTextField.Text())
		go monitorClusterList(clusterListTable)

	})

	//Main stack of the interfaces
	w = ui.NewWindow("Carina by Rackspace GUI Client", 600, 450, mainGrid)
	w.SetMargined(true)

	w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	w.Show()

}

// Set up global connection to the cluster
func connect(endpoint, username, apiKey string) {
	// Connect to Carina
	var err error
	carinaClient, err = libcarina.NewClusterClient(endpoint, username, apiKey)
	if err != nil {
		log.Fatal("Cannot create cluster client: ", err)
	}
}

// monitor the carina client
func monitorClusterList(t ui.Table) {
	for {
		clusters, _ := carinaClient.List()
		t.Lock()
		d := t.Data().(*[]libcarina.Cluster)
		*d = clusters
		t.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// This runs the code that displays our GUI.
	// All code that interfaces with package ui (except event handlers) must be run from within a ui.Do() call.
	go ui.Do(gui)

	err := ui.Go()
	if err != nil {
		log.Print(err)
	}

}
