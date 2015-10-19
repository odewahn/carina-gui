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

	//div grp
	divGrp1 := ui.NewGroup("", ui.Space())
	divGrp1.SetMargined(true)

	// Define the table that lists all running clusters
	var c libcarina.Cluster
	clusterListTable := ui.NewTable(reflect.TypeOf(c))

	// Create control buttons
	newBtn := ui.NewButton("New")
	growBtn := ui.NewButton("Grow")
	rebuildBtn := ui.NewButton("Rebuild")
	credentialsBtn := ui.NewButton("Credentials")
	deleteBtn := ui.NewButton("Delete")
	buttonStack := ui.NewVerticalStack(newBtn, growBtn, rebuildBtn, credentialsBtn, deleteBtn)

	mainGrid := ui.NewGrid()
	mainGrid.Add(loginGrid, nil, ui.East, true, ui.Fill, false, ui.Center, 12, 1)
	mainGrid.Add(divGrp1, loginGrid, ui.South, true, ui.Fill, false, ui.Center, 12, 1)
	mainGrid.Add(clusterListTable, divGrp1, ui.South, true, ui.Fill, false, ui.Center, 9, 1)
	mainGrid.Add(buttonStack, clusterListTable, ui.East, true, ui.Fill, false, ui.Center, 3, 1)
	mainGrid.SetPadded(true)

	connectBtn.OnClicked(func() {
		connect(apiEndpointTextField.Text(), usernameTextField.Text(), apiKeyTextField.Text())
		go monitorClusterList(clusterListTable)
	})

	newBtn.OnClicked(func() {
		newCluster()
	})

	//Main stack of the interfaces
	w = ui.NewWindow("Carina by Rackspace GUI Client", 620, 400, mainGrid)
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

func newCluster() {

	clusterNameLabel := ui.NewLabel("Cluster Name:")
	clusterNameTextField := ui.NewTextField()
	clusterNodeCountLabel := ui.NewLabel("Number of Nodes:")
	clusterNodeCountTextField := ui.NewTextField()
	clusterNodeCountTextField.SetText("1")
	autoscaleLabel := ui.NewLabel("Autoscale:")
	autoscaleCheckbox := ui.NewCheckbox("")
	newClusterBtn := ui.NewButton("Create Cluster")
	cancelBtn := ui.NewButton("Cancel")

	newClusterGrid := ui.NewGrid()

	//	loginGrid.Add(apiEndpointLabel, nil, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)

	newClusterGrid.Add(clusterNameLabel, nil, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	newClusterGrid.Add(clusterNameTextField, clusterNameLabel, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	newClusterGrid.Add(clusterNodeCountLabel, clusterNameLabel, ui.South, true, ui.LeftTop, false, ui.Center, 1, 1)
	newClusterGrid.Add(clusterNodeCountTextField, clusterNodeCountLabel, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	newClusterGrid.Add(autoscaleLabel, clusterNodeCountLabel, ui.South, true, ui.LeftTop, false, ui.Center, 1, 1)
	newClusterGrid.Add(autoscaleCheckbox, autoscaleLabel, ui.East, true, ui.LeftTop, false, ui.Center, 1, 1)
	newClusterGrid.Add(newClusterBtn, autoscaleLabel, ui.South, true, ui.Fill, false, ui.Center, 1, 1)
	newClusterGrid.Add(cancelBtn, newClusterBtn, ui.East, true, ui.Fill, false, ui.Center, 1, 1)
	newClusterGrid.SetPadded(true)

	newClusterGrp := ui.NewGroup("", newClusterGrid)
	newClusterGrp.SetMargined(true)

	newWin := ui.NewWindow("New Cluster", 400, 300, newClusterGrp)
	newWin.SetMargined(true)
	newWin.Show()

	cancelBtn.OnClicked(func() {
		newWin.Close()
	})

	newWin.OnClosing(func() bool {
		newWin.Close()
		return true
	})

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
