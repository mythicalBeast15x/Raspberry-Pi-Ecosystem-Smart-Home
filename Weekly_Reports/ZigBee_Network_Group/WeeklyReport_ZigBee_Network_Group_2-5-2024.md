<h1 align="center"> Weekly Report 02/05/2024 </h1>

## ❖ Team Report

### Previous Team Goals
- Prioritize the development and testing of code for Raspberry Pi.
- Establish communication pathways for Zigbee devices to interact with Raspberry Pi.
- Manage the operation of tangible devices from each unit.

### Key Points
- Our team was able to communicate more effectively with each other to establish a more consistent communication structure.
- Managed to establish initial connections between device nodes using the Digi Xbee Python library over a Wifi connection.
- Discussed possible data structures to hold communication data to be sent between device nodes.

### Next Week Goals
- Get the Zigbees connected with other Zigbees.
- Connect and manipulate appliances to a Raspberry device node locally/remotely.
- Finalize the design and implementation of a JSON communication data structure.
- Continue looking into and succeed in configuring a Zigbee Coordinator for improved communication between Zigbee devices without relying on Wifi.

## ❖ Individual Report

### ❖ Love Patel

#### Previous Individual Goals/Objectives
- Get started on working on the setup for the Appliances System and map out the codebase.
- Communicate with other Zigbee members and collaborate to understand better how to connect all the Raspberry Pi to build an overall system.
- Learn and try to work with the Zigbee devices and attempt to mess around to understand how they work.
- Try to have some code to test on the Raspberry for the Appliances system.

#### Key Points


#### Next 2 Weeks Goals


### ❖ Deep Patel

#### Previous Individual Goals/Objectives
- Begin the setup process for the HCI (Human-Computer Interaction) part and outline the codebase structure.
- Collaborate with other members of the Zigbee group focusing on Xbee communication with Raspberry Pi units to comprehend the integration for creating an overarching system.
- Dedicate time to learning and experimenting with Xbee devices, aiming to gain insight into their functionalities.
- Work towards generating code for testing on the Raspberry Pi within the context of the HCI.

#### Key Points
- Advanced efforts in setting up Zigbee connectivity, focusing on establishing connections and creating a mesh network. 
- Utilized Go lang libraries for Zigbee network setup, actively working towards successful device connection. 
- Integrated HCI demo code into the project, showcasing initial functionalities and interactions. 
- Collaborated with the frontend team to incorporate HCI features into the user interface, ensuring seamless integration and user experience.

#### Next Week Goals
- Finalize HCI setup; ensure comprehensive coverage.
- Continue collaboration with the Zigbee team on integration.
- Deepen understanding of Zigbee devices.
- Get Zigbees connected to create a mesh network.
- Advance code development for Raspberry Pi testing.

### ❖ Timothy Enders

#### Previous Individual Goals/Objectives
- Contemplate and initiate the development of a smart home lighting system on Raspberry Pi.
- Explore the fundamentals of connecting with Zigbee devices and writing code, collaborating with fellow team members.

#### Key Points
- Assisted in establishing a communication link between two Xbee devices using the Digi Xbee Python library.
- Enhanced the initial Go code for the lighting system to incorporate potential extra features.

#### Next Week Goals
- Investigate alternative Go libraries for establishing communication between Xbee devices, as Python is not preferred over Go.
- Gain further understanding of connecting Zigbee devices in a network.
- Employ communication channels to facilitate communication between Zigbee devices.
- Enhance the Go code for the advanced lighting system to integrate testing capabilities.

### ❖ Joshua Murillo

#### Previous Individual Goals/Objectives
- Form a solid communication network with the rest of my team.
- Integrate the Go-Zigbee package into my current appliance state change functions to both receive and service requests from multiple device nodes.
- Implement Blockchain logging into each step of the appliance creation and modification process.

#### Key Points
- Was able to communicate with several Zigbee group members to progress our assigned tasks.
- Worked with Zigbee group members to establish a communication link between two Xbee devices with the Digi Xbee Python library.
- Discussed and documented a proposal for a common data structure for our entire Thursday group to use for user and system-level communication.
- Worked on foundational Python and Go scripts for implementing the proposed data structure into JSON objects.
- Worked on finding Golang solutions to communicating between Zigbee devices but ran into issues with functions from go-zigbee repositories not functioning with feedback stating that the *Address* function was not a valid command. The Repo tested was the following: https://github.com/ebusto/xbee
- I was not able to work with Blockchain logging yet as I am still working on the implementation of the request communication structure.

#### Next Week Goals
- Discuss with the Thursday team about the request data structure proposal and incorporate any relevant suggestions into a final version of the structure.
- Fully implement the agreed-upon structure into our code base so that anyone can call a function to create a request object and send it out to other connected devices.
- See how a physical appliance could be connected to a device node and get a connected appliance to turn on locally. After this, if the request data structure implementation progresses further, then manipulating the appliance via a request from a remote device node would be the next logical step if I have enough time to get my previous goals done.
- Currently, the working Python solution for our communication system works specifically with Xbee devices. More work is needed to configure a Zigbee Coordinator to allow a Zigbee device to initiate a Zigbee network system for other devices to link up to without relying on Wifi. Since removing the need for Wifi is a big part of our minimum viable product, this is one of the more important goals to work on implementing as soon as possible.
- Continue looking for Go-based solutions to Zigbee communications.


### ❖ Jahidul Robin

#### Previous Individual Goals/Objectives
- Collaborate more closely with the Network Traffic Analysis System group to develop a monitoring system that alerts us to any unusual activity within the network, enhancing our security measures.
- Start implementing the blockchain logging mechanism for the security system to ensure that the history of security-related events is immutable and securely logged. This will involve understanding SHA256 encryption and how it can be applied to our use case.
- Plan for a small-scale test of the security system integration with the overall mesh.

#### Key Points


#### Next Week Goals
