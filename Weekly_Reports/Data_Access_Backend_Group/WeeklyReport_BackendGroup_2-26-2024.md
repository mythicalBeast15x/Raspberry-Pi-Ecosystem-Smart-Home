# Weekly Report 2/26/2024
Follow this template exactly for every report including the name of the document. Use this to help delegate tasks to each member. Everytime has one weekly report. There should be 1 team section and multiple individual sections for each person.
Format this well, just copy and paste this all and remove all the boiler plate stuff. Make it look professional and nice. This is seperate from the documentation you may be keeping in the Microsoft Teams. Although, you can talk about that in this. These are due every Monday unless stated otherwise. They will be placed within your team's respective folder.

## Team Report
This is the entire team and what you have worked on. Make these all bullet points.

### Previous Team Goals
- Schedule a meeting with the Network Analysis team.
- Schedule a meeting with the ZigBee team.
- Build prototype for user login token system.

### Key Points
- Met with the front end team.
- Met with the ZigBee team.
- Had a rough overview on how everything is at the moment.
- Prototype for user login token system are done.
- Progress on DAL
- Middleware connections to front end (lighting)
- We Still need to work on a user verification system
- We still need to build the database for verified user profile. 
  
### Next Week's Goals
- Database for verified user profile. 
- User login system.
- Begin helping other teams (if needed)

## Individual Report
This is for each person and what you have focused on. Make these all bullet points.

### Previous Team Goals
- Include things like what you worked on last week.
  
### Key Points
- What you worked on this week?
- What did work?
- What did not work if there were any?
- Anything important?
  
### Next Week's Goals
- Next week's goals that you plan to focus on or learn about. These are 2 week sprints, so keep that in mind.

## Individual Report - Xingyu Jiang
### Previous Goals
- Meeting with the Network Analysis team and the Zigbee team.
- Start build user login token system
- 
- Discuss with other groups making sure everyone is on the same page.
- Research on HMAC.
- Come to an agreement on the consensus algorithm that is going to be implemented.
- Implement the P2P aspect of blockchain.

### Key Points
- Met with the front end team and members on the ZigBee team.
- Gain a rough overview on where we are at regarding the project as a whole.
- Build a prototype for user login token system.
- Make sure key utilities are set up and ready to merge once everything is in place.
- Did more research on ZigBee devices to better understand what to expect.

### Next Week's Goals
- Start working on database building. 
- Set up data serialization and deserialization process for data transmission between HTML webpages and the back end system. 
- Set up server like structure for process user login information and verification of the said information.
- Ensure data integrity during data transformation with HMAC.

## Individual Report - David Weisman
### Previous Goals
- Continue collaboration with frontend as they add more functionalities
- Gain better understanding of FE needs
- Work on things more often as opposed to doing all the work on one day

### Key Points
- Met and worked with the front end (caught up with Zigbee and Networking as well)
- Worked the middleware connection, specifically for lighting
- Added to team deliverable
- Much better understanding of the project (each team and their roles and how it connects)

### Next Week's Goals
- User log in database set up
- Begin shifting to help other groups (Possibly?)

## Individual Report - Catherine Monteith
### Previous Goals
- Refine & Properly create base model of DAL (beyond Pseudo)
- Meet with Front End team to better shape DAL to fit their needs and figure out how to communicate to some of
  the existing functions
- Touch Base again with Zigbee to see if I can gain anymore insight
  

### Key Points
- DAL is structurally approved
  - Major Issues: 
    - OperationIDs - These are currently in a dictionary-like structure 
  and unfortunately have not yet been modified to not be 0-11,
    - Aspects of sending and receiving
      - Currently each core DAL function takes in a STRUCTURE and MODIFIES the ATTRIBUTES within it. We need to solidify How this structure is sent (ex: light struct) 
      - Pseudo//comments exist for the rough functionality of how DAL will take in messages, but, again, given the above there is currently a missing argument) 
    - Front-End team will need a way to access OIDs in order to work with the DAL
  - Discussed these points pretty heavily with Josh, I need more collaboration and group understanding to properly cater! 
  - Worked with Front-end to solidify existing structures and go over developer tools. 

### Next weeks goals: 
- SOLIDIFY DAL, work closer with Zigbee team so that better communication can happen with front-end. 
- Solidify aspects of communication 
- solidify plans for data storage and retrieval so that front-end will be able to dynamically load when they are ready to. 

