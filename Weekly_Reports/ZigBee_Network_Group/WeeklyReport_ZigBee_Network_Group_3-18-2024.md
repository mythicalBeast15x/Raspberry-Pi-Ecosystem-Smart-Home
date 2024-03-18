
<h1 align="center"> Weekly Report 3/18/2024 </h1>

## ❖ Team Report

### Previous Team Goals
- Have the AES integrated when sending data to the network and possibly blockchain.

### Key Points
- Integrated AES for data transmission with the network.

### Next Week Goals
- Continue the steps towards blockchain integration.

## ❖ Individual Report

### ❖ Love Patel

#### Previous Individual Goals/Objectives

#### Key Points


#### Next Week Goals


### ❖ Deep Patel

#### Previous Individual Goals/Objectives
- Have the AES working when data is sent to the network with the help of network team.
- Start integrating blockchain functions and test them.

#### Key Points
- Modified and tested client and controller that utilizes Josh's messaging structure, and it has AES implemented.
- Modified entire frontend again and connected each file with each other.
- Wrote go functions for frontend in login.go file and successfully built and ran the code in Pi's environment.
- As of now I have implement hardcoded users to test and generate HTML.

#### Next Week Goals
- Get the MongoDB connect with login and signup.
- Figure out why some data are being corrupted when sent through the RF.
- Write go functions for dashboard which is being generated when user logs or sign in.

### ❖ Timothy Enders

#### Previous Individual Goals/Objectives
- Determine how to adapt lighting functions to fit within the backend/frontend structure.
- Support the communication and interaction of ZigBee devices with the backend.
- Implementate blockchain functions and conduct testing.
- Collaborate with other teams to finalize the core logic of the project.
- Upon completion of the foundational structure, commence physical testing of ZigBee devices while concurrently documenting and capturing visual interactions of these devices.
- Aid in the development of startup functions.
- Complete the setup of the web certificate.

#### Key Points
- Revamped Cait's config files to read in a temporary AES key and a temporary hash key. The list of appliances will still need to be done.
- For now, there is a function within startup.go that generates the private and public key.
- The response function currently sets up a message with the private key and temporary hashed hash key on the network. It is encrypted with the temporary AES key.
- I will need more insight on how to send the message out and if it needs blockchain to function.
- In the response function, it goes to announceStartUp function with an int value. The value given depends on the response. The announceStartUp has switch cases to determine what to do next.

#### Next Week Goals
- Get ready for MVP presentation.
- Further configure startup.go and config file based on team suggestion and once main.go prototype is set up.
- Help with DAL integration.

### ❖ Joshua Murillo

#### Previous Individual Goals/Objectives

#### Key Points

#### Next Week Goals

### ❖ Jahidul Robin

#### Previous Individual Goals/Objectives
- Enhance the code for doorlock, camera, and motion detection
- Keep going around asking groups if they need assistance.
- Work on anything that needs to get done.

#### Key Points


#### Next Week Goals

### ❖ Mitkumar Patel

#### Previous Individual Goals/Objectives

#### Key Points

#### Next Week Goals
