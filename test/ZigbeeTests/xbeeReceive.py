'''
This script showcases a potential way to receive data from other xbee devices before writing the data into
a common text file called incomingRequest.txt

Future functionalities include:
- Increase the robustness of the script to limit errors.
- Continue looking for better ways to communicate between Zigbee devices without relying on internet.
'''
import os
from digi.xbee.devices import XBeeDevice

# Replace with the serial port where your local module is connected to!
PORT = "/dev/ttyUSB0"
# Make sure all devices are using the same baud rate!
BAUD_RATE = 9600

def main():
    print(" +-----------------------------------------+")
    print(" | XBee Python Library Receive Data Sample |")
    print(" +-----------------------------------------+\n")
    OUTPUT_FILE = "incomingRequest.txt"

    # Check if incomingRequest.txt file exists
    if os.path.exists(OUTPUT_FILE):
        device = XBeeDevice(PORT, BAUD_RATE)

        try:
            device.open()

            def data_receive_callback(xbee_message):
                receivedData = xbee_message.data.decode()
                print("From %s >> %s" % (xbee_message.remote_device.get_64bit_addr(),
                                         receivedData))
                # Write the received message to an output file
                with open(OUTPUT_FILE, "a") as f:
                    f.write(receivedData)

            device.add_data_received_callback(data_receive_callback)
            print("Waiting for data...\n")
            input()

        finally:
            if device is not None and device.is_open():
                device.close()


if __name__ == '__main__':
    main()