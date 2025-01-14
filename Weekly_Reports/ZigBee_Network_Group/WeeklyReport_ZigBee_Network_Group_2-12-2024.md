<h1 align="center"> Weekly Report 02/12/2024 </h1>

## ❖ Team Report

### Previous Team Goals
- Get the ZigBees connected with other ZigBees.
- Connect and manipulate appliances to a Raspberry device node locally/remotely.
- Finalize the design and implementation of a JSON communication data structure.
- Continue looking into and succeed in configuring a ZigBee Coordinator for improved communication between ZigBee devices without relying on Wi-Fi.

### Key Points
- Explored Go-based solutions for communicating between ZigBee devices without Wi-Fi.

### Next Week Goals
- Explore ways to get our Raspberry Pis to manipulate physical appliances.
- Agree on a method of continuous listening of incoming messages by every ZigBee device on the network.
- Test for bugs in our ZigBee Network setup.

## ❖ Individual Report

### ❖ Love Patel

#### Previous Individual Goals/Objectives
- Finalize Appliances along with most of the functionality for Appliance usage.
- Work with the energy system in order to figure out energy optimization along with priority.
- Configure ZigBee devices and get them to communicate with each other.
- Get ZigBee connected to create a mesh network.
- Advance code development for Raspberry Pi testing.

#### Key Points
- Enhanced collaboration with the ZigBee team, progressing integration efforts. 
- Expanded knowledge and proficiency with ZigBee devices. 
- Start coding for the Appliances and work with Network Team and Security to finish AES Encryption.
- Advance code development for Raspberry Pi testing.

#### Next Week Goals
- Setup and complete Zigbee Communication 
- Finish Zigbee Device Configuration
- Conduct thorough testing of the ZigBee mesh network for stability and efficiency. 
- Begin Documentation for Zigbee Configuration and communication.
- Continue the code for the Appliances part of the project.

### ❖ Deep Patel

#### Previous Individual Goals/Objectives
- Finalize HCI setup; ensure comprehensive coverage.
- Continue collaboration with the ZigBee team on integration.
- Deepen understanding of ZigBee devices.
- Get ZigBee connected to create a mesh network.
- Advance code development for Raspberry Pi testing.

#### Key Points
- Enhanced collaboration with the ZigBee team, progressing integration efforts. 
- Expanded knowledge and proficiency with ZigBee devices. 
- Successfully established ZigBee connection to initiate a mesh network. 
- Continued advancement in Raspberry Pi testing code development.

#### Next Week Goals
- Set up a constant connection between ZigBees so they can send and receive constantly.
- Conduct thorough testing of the ZigBee mesh network for stability and efficiency. 
- Implement any necessary optimizations or adjustments based on testing results. 
- Begin the documentation process for the HCI setup and ZigBee integration procedures. 
- Explore additional functionalities or features to enhance the Raspberry Pi testing environment.

### ❖ Timothy Enders

#### Previous Individual Goals/Objectives
- Investigate alternative Go libraries for establishing communication between XBee devices, as Python is not preferred over Go.
- Gain further understanding of connecting ZigBee devices in a network.
- Employ communication channels to facilitate communication between ZigBee devices.
- Enhance the Go code for the advanced lighting system to integrate testing capabilities.

#### Key Points
- Integrated a draft color cycle function into the lighting system and created a test initiation file.
- Developed a draft Entity-Relationship Diagram (ERD) for the ZigBee mesh network as part of phase 2 documentation.
- Experimented with Pauleyj's gobee library on a Raspberry Pi, facing challenges in establishing communication with the serial port.
- Collaborated with Josh in configuring XBee devices using the go-serial library and XCTU program, successfully transmitting packets via XCTU but encountering difficulties with Raspberry Pi. Devices were handed over to team members Love and Deep for model and firmware configuration.

#### Next Week Goals
- Enable constant communication among ZigBee devices within the network.
- Modify and update the lighting system to accommodate specific ZigBee devices.
- Continue refining the phase 2 documentation.
- Enhance collaboration across teams to address communication aspects and functional interactions.
- Initiate physical testing of ZigBee devices both locally and remotely.
- Document and capture visual interactions of ZigBee devices.

### ❖ Joshua Murillo

#### Previous Individual Goals/Objectives
- Discuss with the Thursday team about the request data structure proposal and incorporate any relevant suggestions into a final version of the structure.
- Fully implement the agreed-upon structure into our code base so that anyone can call a function to create a request object and send it out to other connected devices.
- See how a physical appliance could be connected to a device node and get a connected appliance to turn on locally. After this, if the request data structure implementation progresses further, then manipulating the appliance via a request from a remote device node would be the next logical step if I have enough time to get my previous goals done.
- Currently, the working Python solution for our communication system works specifically with XBee devices. More work is needed to configure a ZigBee Coordinator to allow a ZigBee device to initiate a ZigBee network system for other devices to link up to without relying on Wi-Fi. Since removing the need for Wi-Fi is a big part of our minimum viable product, this is one of the more important goals to work on implementing as soon as possible.
- Continue looking for Go-based solutions to ZigBee communications.

#### Key Points
- Worked with Tim to test our ZigBee devices using the XCTU software. We managed to work out some targeted communication between our devices using the XCTU Digi software.
- Tested out Go packages and repositories in an effort to replace the Python-based systems we currently have in place.
- Was unable to successfully identify the ZigBee 16-bit local address necessary to connect the other devices to the ZigBee network.
- Had issues with the current ZigBee configuration not providing the PAN ID, leading to me and other members handing our ZigBee devices to Deep and Love in order for them to homogenize our team's ZigBee devices to the same configuration.
- Discussed the current state of the messaging structure with other group members and updated the structure accordingly both in its documentation and code.
- Documented the messaging structure in our February assignment into a Data Structure Diagram (DSD) with relevant context.
- Began working on documenting the workflow of our messaging system into a Workflow Diagram for our February assignment.
- Reviewed the circuit diagrams created by Love Divine in preparation for manipulating physical appliances with our Raspberry Pis.

#### Next Week Goals
- Implement a priority queue to hold excess outgoing requests into the messaging package.
- Further, discuss the broadcasting structure of our ZigBee communication system to effectively target specific devices while sending out messages to all devices simultaneously.
- Further, discuss the overall ZigBee communication structure with the rest of the ZigBee team and with members of other teams to fully cement the communications system of our project's minimum viable product.
- Further, discuss the continuous listening of incoming requests by every ZigBee device with the rest of the team.
- Work with and implement code to manipulate physical devices with our Raspberry Pis.
- Make more contributions to the February assignment.

### ❖ Jahidul Robin

#### Previous Individual Goals/Objectives
- Address the challenges that hindered progress last week, starting with a thorough review of the technical and collaborative obstacles encountered. This will involve seeking additional resources and possibly redefining the approach to collaboration with the Network Traffic Analysis System group.
- Reassess the blockchain logging mechanism's implementation plan. This may include simplifying the initial model to make the development process more manageable and ensure timely completion. The focus will be on understanding and applying SHA256 encryption in a more streamlined manner.
- Prioritize tasks more effectively to ensure that foundational elements of the security system are developed and ready for integration testing. This might involve setting more focused, short-term goals to achieve incremental progress.
- Establish a more realistic timeline for the completion of the security monitoring system and the blockchain logging mechanism, taking into account the lessons learned from the setbacks encountered.

#### Key Points
-

#### Next Week Goals
-

### ❖ Mitkumar Patel

#### Previous Individual Goals/Objectives
- Begin advancing the HVAC code by implementing more sophisticated control algorithms and features, such as temperature scheduling, energy-saving modes, and responsive adjustments based on real-time sensor data.
- Continue improving ZigBee network stability and reliability, focusing on optimizing signal strength and interference mitigation.
- Enhance the HVAC system's user interface for better usability and accessibility, incorporating feedback from initial testing.
- Collaborate with other team members to ensure seamless integration between the HVAC system and other components of the smart home project.
- Advance Raspberry Pi config and connect with the mesh network.

#### Key Points
- HVAC Code Development: Made progress in developing advanced control features for the HVAC system. Implemented algorithms for temperature scheduling and attempted energy-saving modes setup as well, enhancing the system's efficiency and user control.
- Zigbee Network Optimization: Achieved improvements in the Zigbee network's stability and reliability. Conducted extensive testing to identify and reduce interference, resulting in stronger signal strength across the network.
- Collaboration and Integration: Worked closely with other team members to ensure that the Zigbee team work is moving with the flow.
- Raspberry Pi: Provided Love the Raspberry Pi to have him work on networking them together. 

#### Next Week Goals
- Refine HVAC Control Algorithms: Continue refining the HVAC system's control algorithms to ensure optimal performance and user satisfaction.
- Integration Testing: Start to conduct integration testing with the entire smart home system to identify and resolve issues.
- Documentation and Reporting: Begin compiling documentation on the HVAC system's development process and outcomes, preparing for project reporting.
