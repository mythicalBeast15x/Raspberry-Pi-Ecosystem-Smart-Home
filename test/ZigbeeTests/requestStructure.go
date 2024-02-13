/*
This sample code contains the foundation of a potential data structure for creating requests to be sent out to
relevant device nodes.

Further discussion with team members is required before advancing this script with further functionalities.
Future functionalities include:
- A function to write requests into a text file with a known common name: outgoingRequest.txt
- A function for checking a text file with a known name to see if it is populated with a request: incomingRequest.txt
- Implementing a way to manage multiple outgoing/incoming requests being called simultaneously.
- Add a function to invoke the python file to send a message from Golang: "python xbeeSend.py"
*/
package main

import (
	"encoding/json"
	"fmt"
)

// CommandParameter represents a parameter for a command.
type CommandParameter struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// Command represents a sub-command with optional parameters.
type Command struct {
	Name       string             `json:"name"`
	Parameters []CommandParameter `json:"parameters,omitempty"`
}

// Request represents the structure of a request object.
type Request struct {
	Source   string    `json:"source"`
	Domain   string    `json:"domain"`
	Subject  string    `json:"subject"`
	State    string    `json:"state"`
	Commands []Command `json:"commands,omitempty"`
}

// NewRequest creates a new Request object with the given parameters.
/*
Purpose: Creates a new request object.
Data:
- source: The source of the request (e.g., user or system).
- domain: The domain of the request (e.g., lighting, HVAC, security).
- subject: The specific device or appliance related to the request.
- state: The state of the device or system (e.g., on, off).
Returns: Request: The newly created Request object.
*/
func NewRequest(source, domain, subject, state string) Request {
	return Request{
		Source:  source,
		Domain:  domain,
		Subject: subject,
		State:   state,
	}
}

// AddCommand adds a new command to the request object.
/*
Purpose: Attaches an additional command to a request.
Data:
- name: The name of the command.
- params: Optional parameters associated with the command.
Params:
- name: The name of the parameter.
- value: The value of the parameter.
Returns: None
*/
func (r *Request) AddCommand(name string, params ...CommandParameter) {
	command := Command{Name: name, Parameters: params}
	r.Commands = append(r.Commands, command)
}

// DisplayRequest prints out the Json data.
/*
Purpose: Shows the data in the output window for testing purposes.
Data :
- req: The Request object to be displayed.
Prints:
- Source: The source of the request.
- Domain: The domain of the request.
- Subject: The subject of the request.
- State: The state of the request.
- Commands: The list of commands associated with the request, along with their parameters.
Returns: None
*/
func DisplayRequest(req Request) {
	fmt.Println("Source:", req.Source)
	fmt.Println("Domain:", req.Domain)
	fmt.Println("Subject:", req.Subject)
	fmt.Println("State:", req.State)
	fmt.Println("Commands:")

	for _, cmd := range req.Commands {
		fmt.Println("  Command:", cmd.Name)
		for _, param := range cmd.Parameters {
			fmt.Printf("    Parameter: %s = %v\n", param.Name, param.Value)
		}
	}
}

func main() {
	// Create an initial request with no command arguments.
	req := NewRequest("admin", "lighting", "bathroom", "on")

	// Attach relevant commands. As many extra commands can be added as needed
	req.AddCommand("setTime", CommandParameter{"time", "0600"}, CommandParameter{"reoccurring", true})
	req.AddCommand("setState", CommandParameter{"color", "white"}, CommandParameter{"intensity", 50})

	// Serialize the request into JSON format
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the raw JSON data.
	fmt.Println()
	fmt.Println(string(jsonData))

	/*
		At this point, the data can be encrypted and saved to a text file.
	*/

	// deserialize the JSON data into a Go Request Object.
	var loadedReq Request
	err = json.Unmarshal(jsonData, &loadedReq)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the contents of the request object.
	fmt.Println()
	DisplayRequest(loadedReq)
}
