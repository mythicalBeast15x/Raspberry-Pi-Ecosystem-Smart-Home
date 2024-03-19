# Weekly Report 3/18/2024

## ❖ Team Report

### Previous Team Goals

- Generate device status map from frontend webpage

### Key Points

- Generated device status map from frontend webpage
- Linked map to Go using AJAX
- Function to create and operate user collection in mongoDB
- Function to createa and operate blockchain collection in mongoDB
- Serializing and deserializing user list 
- Serializing and deserializing blockchian

### Next Week Goals

- Integrate device status map with messaging structure to ensure data is consistent across all devices

## Individual Report

### ❖ Catherine Monteith

### Key Points

- Generated device status map from frontend webpage
- Linked map to Go using AJAX
- Worked on DAL

### Next Week Goals

### ❖ Caitlin Crowe

### Key Points

- Generated device startup config file
- Stored config file into in a structure (this is where the aforementioned device map will be held)
- Generated a map of device statuses with their respective device IDs

### Next Week Goals

- Properly link FrontEnd package to the main function
- Send device map to other Pis via messaging structure

### ❖ Joshua Murillo

#### Previous Individual Goals/Objectives

- Work with Zigbee and Network teams to fix issues related to the go-serial library and start consistently sending and receiving data in our Message structured format.
- Work with the rest of the Frontend-Data-Access team to progress the DAL to service incoming requests.

#### Key Points

- Met with Cat to discuss a clearer picture between the messaging structure, DAL, and the main.go file.
- Set up two scripts including a sending and receiving script to transmit messages between PIs one-way.
  - Ran tests with three Raspberry PIs to document the frequency of successfully received messages vs missed messages.
    - The tests were conducted 45 times (15 times per PI acting as a sender with the other two acting as a receiver).
    - In the end, the results were 31.11% success rate, 55.55% partial success rate, and a 13.33% complete failure rate.
- Set up a prototype of the main.go file based on the previous sender and receiver scripts which featured three go-routines including a receiver, sender, and servicer routine which allows for this single script to handle all aspects of the message transmission system.
  - Ran into issues with corrupted messages frequently being pulled down from the receiver routine.
    - Corrupted data appeared to be some amalgamation of two messages meshed together as if the receiver pulled down the message right as another message was intersecting it.
    - Occasionally, messages would not be received corrupted or otherwise by at least one of the other PIs.
  - Ran several test trials to see the failure rate given some length of buffer time between receiving attempts.
    - This time, all Raspberry PIs could act as sender receiver, and servicer all at once.
    - Buffer times started at 1000 ms (1 second), and increased by 500 ms for each set of trials until 4500 ms. Each set of trials was run 5 times with each PI's success, failure, and corrupted data rates recorded.
    - In the end, there seemed to be no real improvement by varying the buffer times as corrupted data occurrences continued at a somewhat random rate regardless. I concluded that the more messages are being sent at once, the more likely that messages will be corrupted, however, with potentially 5 PIs being used for our smart home system, the chances of many messages being up in the air at a similar period is very likely, as such, so is the likelihood of many corrupted messages.
    - Note: corrupted messages do not break the system, but are simply ignored by the system if detected potentially leading to missed messages.
- Met with Love Patel and Deep Patel a few times to try and increase the consistency of successfully transmitting messages.
  - We ultimately decided that it may just be a limitation of the Zigbee hardware concerning having many messages up in the air.
  - We still plan on tinkering with the system further.
- Discussed the current setup of the prototype main.go file with Love Divine and updated the file to use channels instead of so frequently using unnecessary for-loops for everything.
  - Confirmed that the entire message transmission system still operates as expected.

#### Next Week Goals

- Help the DAL converge with the main.go system.
- Help integrate the main.go startup function into the prototype.
- Update hardcoded constants related to the messaging package to use the configuration file values instead.
- Prepare for the MVP demonstration.

### ❖ Xingyu Jiang

#### Previous Individual Goals

* Set up connection with the mongoDB

#### Key Points

* Function for the user collection  creation.

* Function for the blockchain collection creation.

* Basic operation for both the user collection and blockchain collection are set up.

* Serialization and deserialization for the data. Prepare them to be sent over message.

#### Next Week's Goals

* Work with the front end group to complete the user sign-up process.
  
  * User credential gathering.
  
  * Verification.
  
  * Registration.
  
  * Storing and announcing. 
