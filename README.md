# Penn State Abington Capstone 2024 - Raspberry Pi Ecosystem Smart Home Project
<p align="center">
  <img src="https://img.shields.io/badge/Contributors-16-blue" /></a>
  <img src="https://img.shields.io/badge/Development_Stage-Early_Development-orange" /></a>
  <img src="https://img.shields.io/badge/License-Not_Picked-green" /></a>
</p>

## About the Project
This projected is being developed by the graduating Computer Science Majors at Penn State Abington. The project aims to demonstrate the knowledge acquired throughout the course of each student's educational journey. Team collaboration is
heavily emphasized through all development stages of the application.
**More will be added**

## Mission Statement
“Creating a secure network of interconnected devices in order to automate lives.”
Our mission is to add simplicity into the home automation space while also delivering security and reliability. Our product will feature a user-friendly design so that those who may not be experts in this domain will be able to effortlessly introduce smart technology into their homes. Our product will be able to adapt to the new technologies that are installed and seamlessly add them to the network to create a customizable ecosystem that enhances the quality of life.

### Problem Definition
More and more homes are taking advantage of new smart technologies and integrating them into their homes to automate parts of their lives. However, integrating smart technologies has a couple drawbacks and challenges when setting it up. Problems like the evolution of technology cause people to have to control their devices with multiple applications, which takes away from the convenience of having the smart technologies. Other problems such as security of the network become a concern when there are multiple devices connected. This creates distrust in people considering integrating smart technology. While solutions exist to solve aspects of these problems, they are often incomplete, complicated, or very expensive. 
Our goal is to create an adaptable, user-friendly system that integrates the different technologies in a smart home in a reliable and secure way that removes the barrier to entry. Our system will allow a user to control, manage, and monitor each device from a single place, removing the need to manage multiple applications to control every device. This system will allow for multiple ways to manage the smart devices: a mobile website that one can access from anywhere, and a physical input one can use from inside the house. 

### Objectives and Features
Our objectives are to create a mesh network made up of Raspberry Pi’s that will communicate with each other using Zigbee. Each Raspberry Pi will have a different device connected to it and will control it according to the user. To ensure that the history of events will not be tampered, every Raspberry Pi will log the events using blockchain with Sha256. To ensure that the data is not read in between devices, each Raspberry Pi will encrypt data using AES Encryption. 
The main feature of our system is the reliability of the network. Any number of devices will be able to go offline at any time, and the system and remaining devices will still function normally. And when the devices come back online, it will receive an up-to-date history of the logs, and will function normally.
Lastly, our objective is to be open with the documentation, so that others can add new capabilities to the system with relative ease. In order to do this, all the functions and classes will be heavily documented.

### Additional Features
The product we have is the base system, but we hope to add certain features to the system to enhance its usability. One such improvement is the ability to add several profiles and preferences for one system, so that the user can switch between certain settings for the system without having to remember each setting for every device. Another improvement would be to add scheduling to the system. This would allow certain devices to be activated by a specific time, or when a specific event occurs. We would also like to optimize the energy consumption of the system, by powering the devices when the energy rate is lower. Lastly, we would like to allow new devices to join the system without having to reset the system. This would allow the user to add as many devices as possible without having to restart their home system.


## Thursday Team Starting Groups

### Important Note
These groups are changeable and may change depending on where work is needed. Some groups might have less work in the beginning (other than research) so people might be shuffled around. The mesh home ZigBee network is itself one whole group, where each person will focus on a portion of each system. All groups will have to communicate effectively with each other, and some groups might work with other groups often. Outside meetings will be needed to meet the goals of the project. There are some things that every group will need to focus on, like block chain logging. Also, please understand these are brief descriptions of everything discussed and things may change throughout the project. No person will be dead set in one area. This means within your group, you need to focus your efforts on wherever it is needed at that time. Everyone must thoroughly research all areas of the project. Understanding the big picture is essential. 

#### Mesh Home Zigbee Network Devices 
This group is treated as one team where each person focuses on a certain device or the energy of those devices. Collaborate accordingly as many devices may interface similarly. Of course, some devices may require more work than others, so this means you may have to help each other out. 
- Lighting System: Timothy Enders 
- Security System: Jahidul Robin 
- HVAC System: Mitkumar Patel 
- Appliances System: Love Patel 
- Energy System: Joshua Murillo 
- HCI System: Deep Patel 

#### Network Traffic Analysis System 
This group will work closely with the ZigBee Network device group to ensure the network is working correctly and receiving alerts when things change. The user will be able to see what is happening within the network. This group is close to a subgroup of the ZigBee Network device group. 
- Network Traffic Analysis: Harkaran Kaur 
- Network Traffic Analysis: Darshit Patel 

#### Data Access Layer and Backend PC Group 
This group allows everything to connect together. They will be responsible for the database, as well as any other API type functionality that will be needed. The data access layer will expose any needed information to the users of the application. 
- Data Access Layer and Backend PC: David Weisman 
- Data Access Layer and Backend PC: Xingyu Jiang 
- Data Access Layer and Backend PC: Catherine Monteith 
- Data Access Layer and Backend PC: Caitlin Crowe 

#### User Web Frontend and Middleware PC Group 
This group will be responsible for the website portion of the project. This is what the user will see when accessing the website. They will also be responsible for correctly interacting with any API type functionality that is created to retrieve information. 
- User Web Frontend and Middleware PC Group: Joshua Delva 
- User Web Frontend and Middleware PC Group: Elizabeth Johney 

#### Project Manager 
Coordinate project and provide help wherever it may be needed. Main communication to professor and resources that may be needed. Answer questions to the best of their ability.  
- Project Manager: Tyler Thompson 

#### Leader Roles 
These roles include working with the project manager to ensure everything is flowing together. These leader roles are on top of the other roles they are assigned and just provide another means of communication for specialized portions of the project. Will meet with project manager to ensure everyone is on same page. 

##### Team Lead 
This role includes researching all technical portions of the project and ensuring that all groups are functioning according to plan. Works closely with project manager to do code reviews and define scope. Has meetings with project manager often. May have to help in all areas of project. 
- Team Lead: Love-Divine Onwulata 

##### Hardware and ZigBee Specialist 
This role includes trying to answer any questions and do some research in the ZigBee portion of the project. These questions may come from any of the groups involved with working with ZigBee directly. May also look into Raspberry Pi questions if needed. 
- Hardware and ZigBee Specialist: Timothy Enders
## Suggestions
Please direct any suggestions to the project manager or team lead. If you wish to support this project in any way, please contact the project manager.
