//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func serverStatus(servs map[string]int) {
	totalCount := 0
	onlineCount := 0
	offlineCount := 0
	maintenanceCount := 0
	retiredCount := 0

	for _, value := range servs {
		if value == 0 {
			onlineCount++
		}
		if value == 1 {
			offlineCount++
		}
		if value == 2 {
			maintenanceCount++
		}
		if value == 3 {
			retiredCount++
		}
		totalCount++
	}
	fmt.Println("Total number of servers: ", totalCount)
	fmt.Println("Online servers:", onlineCount)
	fmt.Println("Offline  servers:", offlineCount)
	fmt.Println("Servers on Maintenance:", maintenanceCount)
	fmt.Println("Retired servers:", retiredCount)
	fmt.Println("__________________________________________")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	statuses := make(map[string]int)
	for i := 0; i < len(servers); i++ {
		statuses[servers[i]] = Online
	}

	serverStatus(statuses)

	statuses["darkstar"] = Retired
	statuses["aiur"] = Offline
	serverStatus(statuses)

	for i := 0; i < len(servers); i++ {
		statuses[servers[i]] = Maintenance
	}
	serverStatus(statuses)
}
