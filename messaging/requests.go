package messaging

/*
This sample code contains the foundation of a data structure for creating requests to be sent out to
relevant device nodes.

Further discussion with team members is required before advancing this script with further functionalities.
Future functionalities include:
- Implementing a priority queue to hold outgoing requests.
*/
import (
	"fmt"
)

// deviceIDs holds the Zigbee device 64-bit IDs for targeting requests
var deviceIDs = map[string]string{
	"lighting": "0013A200419117FF",
	"PI2":      "0013A20041911BAB",
	"PI3":      "PLACEHOLDER",
	"PI4":      "PLACEHOLDER",
	"PI5":      "PLACEHOLDER",
	"PI6":      "PLACEHOLDER",
}

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
	Target   string    `json:"target"`
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
	// Check if the domain matches a key in the deviceIDs map
	deviceID, ok := deviceIDs[domain]
	if !ok {
		// If the domain is not found in the map, set the Target to an empty string
		deviceID = ""
	}
	return Request{
		Source:  source,
		Domain:  domain,
		Target:  deviceID, // Set the Target field based on the domain
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
	fmt.Println("Target:", req.Target)
	fmt.Println("Subject:", req.Subject)
	fmt.Println("State:", req.State)
	fmt.Println("Commands:")

	for _, cmd := range req.Commands {
		fmt.Println(" Command:", cmd.Name)
		for _, param := range cmd.Parameters {
			fmt.Printf(" - Parameter: %s = %v\n", param.Name, param.Value)
		}
	}
}

/*
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


	//	At this point, the data can be added to a priority queue to
	//	where it can then be encrypted, run through an HMAC process
	//	and sent out to be broadcast to other Zigbee devices.


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
*/
