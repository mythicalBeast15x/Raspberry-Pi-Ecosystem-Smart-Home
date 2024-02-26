<h1 align="center"> Weekly Report 02/19/2024 </h1>

## ❖ Team Report

### Previous Team Goals
- Explore ways to get our Raspberry Pis to manipulate physical appliances.
- Agree on a method of continuous listening of incoming messages by every ZigBee device on the network.
- Test for bugs in our ZigBee Network setup.

### Key Points
- Update go versions on all Raspberry Pi due to conflict with the serial package.
- Working towards JSON file transfer between ZigBee network.

### Next Week Goals


## ❖ Individual Report

### ❖ Love Patel

#### Previous Individual Goals/Objectives
- Setup and complete Zigbee Communication 
- Finish Zigbee Device Configuration
- Conduct thorough testing of the ZigBee mesh network for stability and efficiency. 
- Begin Documentation for Zigbee Configuration and communication.
- Continue the code for the Appliances part of the project.

#### Key Points
- Communication Testing Done - Setup JSON data communication between Zigbees (part 2)
- Continue with Zigbee Network Implementation, JSON data structure for consistent communication.
- Documented and finished the configuration of the Zigbees completely along with getting them on the same firmware

#### Next Week Goals
- Complete JSON data transfer code.
- Enhance the encryption of the message communication with the security team
- Test the communication with multiple pi's
- Start creating Appliances function and send it over to the Data Layer Team
- Work with Backend and Data Layer to setup the communication workflow 

### ❖ Deep Patel

#### Previous Individual Goals/Objectives
- Set up a constant connection between ZigBees so they can send and receive constantly.
- Conduct thorough testing of the ZigBee mesh network for stability and efficiency.
- Implement any necessary optimizations or adjustments based on testing results.
- Begin the documentation process for the HCI setup and ZigBee integration procedures.
- Explore additional functionalities or features to enhance the Raspberry Pi testing environment.

#### Key Points
- Collaborated on frontend development for web app and HCI, focusing on a responsive dashboard and login/signup pages. 
- Partnered with Love Patel to advance Zigbee network implementation, working towards JSON data structure for consistent communication. 
- Documented Zigbee network formation using XCTU for clarity and future reference.

#### Next Week Goals
- Complete JSON data transfer code within Zigbee network. 
- Enhance message/command security with network team. 
- Plan phone connectivity via web app outside network. 
- Test Raspberry Pi to Zigbee connection. 
- Assist frontend team with bug testing and overcoming obstacles.

### ❖ Timothy Enders

#### Previous Individual Goals/Objectives
- Enable constant communication among ZigBee devices within the network.
- Modify and update the lighting system to accommodate specific ZigBee devices.
- Continue refining the phase 2 documentation.
- Enhance collaboration across teams to address communication aspects and functional interactions.
- Initiate physical testing of ZigBee devices both locally and remotely.
- Document and capture visual interactions of ZigBee devices.

#### Key Points
- Collaborated with the team to implement an echo function within Josh's messaging structure. This will be handled in the Data Access Layer.
- Introduced a TLS web certificate and associated keys, facilitating the transition from HTTP to HTTPS for enhanced security. It's very important to note that the web certificate will require trust from users. This should be further looked at.
- Initiated the documentation process detailing the integration and installation procedures of the TLS web certificate, ensuring comprehensive understanding and reference for other project members.

#### Next Week Goals
- Have discussions with the frontend/security members to review the implementation of the TLS web certificate and advance it.
- Assist in finalizing the communication structure of the ZigBee network.
- Initiate physical testing of ZigBee devices, along with documenting and capturing visual interactions of the ZigBee devices.
- Collaborate with other teams to define lighting functions for different areas.

### ❖ Joshua Murillo

#### Previous Individual Goals/Objectives
- Implement a priority queue to hold excess outgoing requests into the messaging package.
- Further, discuss the broadcasting structure of our ZigBee communication system to effectively target specific devices while sending out messages to all devices simultaneously.
- Further, discuss the overall ZigBee communication structure with the rest of the ZigBee team and with members of other teams to fully cement the communications system of our project's minimum viable product.
- Further, discuss the continuous listening of incoming requests by every ZigBee device with the rest of the team.
- Work with and implement code to manipulate physical devices with our Raspberry Pis.
- Make more contributions to the February assignment.

#### Key Points
- Most of the goals/objectives that I originally planned on working on got pushed to the side when I was introduced to a set of new ideas for the messaging structure by Love Divine.
- Dropped the Idea of a priority queue holding outgoing requests.
- Discussed the topic of a completely new version of the messaging structure with Love Divine to replace the current setup.
- Discussed the new messaging structure proposal with the backend team to gain their perspective and incorporate their suggestions.
- I completely reworked the messaging structure to reflect my discussions with Love and the backend team.
- Wrote up documentation explaining the structure of the messaging object along with some diagrams showcasing the structure, messaging lifecycle, and the message check-in and servicing stages for the sake of better understanding the structure and its place within our system. Added the new diagrams to the February assignment.
- Met with Backend team member Catherine over Discord to review the messaging structure and its documentation to make sure that the structure made sense and reflected our previous discussions accurately.
- The messaging structure is almost complete, some more discussion with the backend, frontend, and Zigbee group members is needed to make the connections needed to take the information of the frontend, form it into a message, and have that message moved by the backend, and finally placed into the appropriate functions from the Zigbee team with data parameters that are agreed upon by the whole group.

#### Next Week Goals
- Discuss the messaging structure with other groups and use these discussions to completely implement the messaging structure into our system.
- Retrieve the Zigbee device that was relinquished to other members of the Zigbee team and discuss the status of the Zigbee network with them.
- Work with and implement code to manipulate physical devices with our Raspberry Pis.
- Discuss with Love Patel about the relationship between the Appliances and the energy system and reworking the energy file to house functions relevant to energy management.

### ❖ Jahidul Robin

#### Previous Individual Goals/Objectives
- Work on AES authentication with Network Security group
- Setup authentication tokens

#### Key Points
- Documented and finished enhancing security for AES authentication.
- Worked on making psuedo code and documentation for authentication tokens. (Cannot implement until authentication is done by front-end and back-end.)

#### Next Week Goals
- Work with front-end and back-end group to finish setting up authentication and tokens.
- Work on the encryption of the message communication.
- Pick up any tasks that needs to be finished.
- Help any group that needs it.

### ❖ Mitkumar Patel

#### Previous Individual Goals/Objectives
- Refine HVAC Control Algorithms: Continue refining the HVAC system's control algorithms to ensure optimal performance and user satisfaction.
- Integration Testing: Start to conduct integration testing with the entire smart home system to identify and resolve issues.
- Documentation and Reporting: Begin compiling documentation on the HVAC system's development process and outcomes, preparing for project reporting.

#### Key Points

#### Next Week Goals
