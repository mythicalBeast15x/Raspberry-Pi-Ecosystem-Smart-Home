# Weekly Report 2/12/2024

## Team Report

### Previous Team Goals
- Setting up blockchain Framework
- Setting up & Testing various verification methods

### Key Points
- A basic blockchain framework is set up & has anticipatory functions for encryption and decryption
- Basic Verification functions & hashing functionality
- Functions for JSON Serialization and Deserialization

### Next Week Team Goals
- Discuss with the Network Analysis team to either begin implementation or modify existing functions for encryption/decryption to prepare for integration
- Discuss and review JSON
- Discuss with the Zigbee group to ensure we are anticipating their needs in the current setup


### Next Week's Goals
- Next week goals that you plan to focus on or learn about. These are 2-week sprints, so keep that in mind.

## Individual Report - Catherine Monteith
### Previous Goals
- Understand Consensus Algorithms
- Research and Gain better understanding of how Zigbee group/devices will interact with the blockchain
- Gain understanding of how blockchain is agreed upon by multiple nodes beyond

### Key Points
- Currently Proof of Work Consensus Algorithm has a basic implementation, but this algorithm might need adjustment
- Added function & Tests to check that every block on a chain is verified. 
  - From my research, I have found that some systems do this and check at certain intervals to see if an entire blockchain is valid. 
  - If the Consensus algorithm used is Proof of Work, nodes will often agree upon the longest valid chain 
  - This function is essentially a basis for the above in anticipation of such an issue but might need to be modified 
  as the project moves forward.

### Next Week's Goals
- Discuss with other groups to ensure the blockchain framework is correctly anticipating their needs. 
- Review JSON implementation 
- Look into HMAC to understand how this can help prevent tampered blocks & figure out how and where the blockchain framework might need to support this.  

## Individual Report - Xingyu Jiang
### Previous Goals
- Get a prototype of the consensus algorithm.
### Key Points
- More research on consensus algorithm is done.
- Wasn't able to get a prototype for consensus algorithm.
  - We Still need to figure our how to validate actions each device is going to perform.
  - We Still need to figure our how devices should be communicating.

### Next Week's Goals
- Discuss with other groups making sure everyone is on the same page.
- Research on HMAC.
- Come to an agreement on the consensus algorithm that is going to be implemented.
- Implement the P2P aspect of blockchain.

## Individual Report - David Weisman
### Previous Goals
- Set up better communication as a team
- Add more testing for the blockchain

### Key Points
- Created stub decryption function as placeholder until we get a better understanding of how the networking team plans on implementing that.
- Created testing folder and procedures on said function with room for expansion when it actually gets implemented.
- Reviewed verification function.

### Next Week's Goals
- Gain better understanding of the progress of the other groups, so we can work better together.
- Become more familiar with HMAC and how we can utilize this in our blockchain.
- Get on the same page as a team with the consensus algorithm.
