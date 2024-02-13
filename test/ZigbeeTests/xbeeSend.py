'''
This script showcases a potential way to send data to other xbee devices after reading the encrypted data from
a common text file called outgoingRequest.txt

Future functionalities include:
- Increase the robustness of the script to limit errors.
- Continue looking for better ways to communicate between Zigbee devices without relying on internet.
'''
import os
from digi.xbee.devices import XBeeDevice

# Check if outgoingRequest.txt file exists
if os.path.exists("outgoingRequest.txt"):
    # Open the file and read the encrypted data
    with open("outgoingRequest.txt", "r") as file:
        encryptedData = file.read()

    # Load the encrypted data into a Python variable.
    DATA_TO_SEND = encryptedData

    # Replace with the serial port where your local module is connected to!
    PORT = "/dev/ttyUSB0"
    # Make sure all devices are using the same baud rate!
    BAUD_RATE = 9600


    def main():
        print(" +------------------------------------------------+")
        print(" | XBee Python Library Send Broadcast Data Sample |")
        print(" +------------------------------------------------+\n")

        device = XBeeDevice(PORT, BAUD_RATE)

        try:
            device.open()
            print("Sending broadcast data: %s..." % DATA_TO_SEND)
            device.send_data_broadcast(DATA_TO_SEND)
            print("Success")

        finally:
            if device is not None and device.is_open():
                device.close()


    if __name__ == '__main__':
        main()
