import time
import board
import neopixel
import usb_cdc
import touchio
import usb_hid
from adafruit_hid.keyboard import Keyboard
from adafruit_hid.keyboard_layout_us import KeyboardLayoutUS

# Configuration
num_pixels = 4  # Number of NeoPixels on your board
exec_program = 'skystatus'  # Name of the program to execute

# Initialize USB serial communication
uart = usb_cdc.data

touch1 = touchio.TouchIn(board.TOUCH1)
touch2 = touchio.TouchIn(board.TOUCH2)
keyboard = Keyboard(usb_hid.devices)
keyboard_layout = KeyboardLayoutUS(keyboard)
from adafruit_hid.keycode import Keycode

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
blueList = [{'blink': False, 'show': False}, {'blink': False, 'show': False}, {'blink': False, 'show': False}, {'blink': False, 'show': False}]
colors = [{'red': RED}, {'green': GREEN}, {'blue': BLUE}, {'black': BLACK}, {'yellow': YELLOW}, {'cyan': CYAN}, {'purple': PURPLE}]

time.monotonic()
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
        time.monotonic()
    elif blueList[index]['blink'] and not blueList[index]['show']:
        pixels[index] = BLUE
        pixels.show()
        blueList[index]['show'] = True
        time.monotonic()

while True:
    # Blink blue
    if touch1.value or touch2.value:
        while touch1.value or touch2.value:
            time.sleep(0.1)
        keyboard_layout.write(exec_program)
        time.monotonic()
        keyboard.send(Keycode.ENTER)
    if uart.in_waiting > 0:
        command = uart.read(uart.in_waiting)
        command = ''.join([chr(b) for b in command])
        command = command.split('_')
        print(command)
        # check if index not out of range
        index = int(command[0])
        if type(index) == int and int(index) < num_pixels and int(index) >= 0:
            index = int(index)
            color = command[1]
            # check if the color is in the list
            colorItem = [item for item in colors if color in item]
            if len(colorItem) > 0:
                showColor(index, colorItem[0][color])
            else:
                print('Color is not valid')
                continue
        else:
            print('Index is not valid')
            continue

    for i in range(num_pixels):
        blinkBlue(i)
   
    # Your other code logic here
    time.sleep(0.3)  # Adjust sleep duration as needed

