import time
import board
import neopixel
import usb_cdc

# Configuration
num_pixels = 4  # Number of NeoPixels on your board

# Initialize USB serial communication
uart = usb_cdc.data

# Initialize NeoPixel strip
pixels = neopixel.NeoPixel(board.NEOPIXEL, num_pixels, brightness=0.1, auto_write=False)

RED = (255, 0, 0)
YELLOW = (255, 150, 0)
GREEN = (0, 255, 0)
CYAN = (0, 255, 255)
BLUE = (0, 0, 255)
PURPLE = (180, 0, 255)
BLACK = (0, 0, 0)
WHITE = (20, 20, 20)

# only blue is blinking
ledBlue0 = {'blink': False, 'show': False}
ledBlue1 = {'blink': False, 'show': False}
ledBlue2 = {'blink': False, 'show': False}
ledBlue3 = {'blink': False, 'show': False}

blueList = [ledBlue0, ledBlue1, ledBlue2, ledBlue3]

colors = [{'red': RED}, {'green': GREEN}, {'blue': BLUE}, {'black': BLACK}, {'yellow': YELLOW}, {'cyan': CYAN}, {'purple': PURPLE}]


pixels.fill(WHITE)
pixels.show()

def showColor(index, color):
    if color == BLUE:
        blueList[index]['blink'] = True
    else:
        pixels[index] = color
        if blueList[index]['blink']:
            blueList[index]['blink'] = False
        pixels.show()

def blinkBlue(index):
    if blueList[index]['blink'] and blueList[index]['show']:
        pixels[index] = BLACK
        pixels.show()
        blueList[index]['show'] = False
    elif blueList[index]['blink'] and not blueList[index]['show']:
        pixels[index] = BLUE
        pixels.show()
        blueList[index]['show'] = True

while True:
    if uart.in_waiting > 0:
        command = uart.read(uart.in_waiting)
        command = ''.join([chr(b) for b in command])
        command = command.split('_')
        print(command)
        # check if index not out of range
        index = int(command[0])
        color = command[1]
        # check if the color is in the list
        colorItem = [item for item in colors if color in item]
        if len(colorItem) > 0:
            showColor(index, colorItem[0][color])
        
    blinkBlue(0)
    blinkBlue(1)
    blinkBlue(2)
    blinkBlue(3)
    # Your other code logic here
    time.sleep(0.3)  # Adjust sleep duration as needed
