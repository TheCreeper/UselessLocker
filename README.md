```
| |  | |           | |                 | |                 | |
| |  | | ___   ___ | |  ___  ___  ___  | |      ___    ___ | | __ ___  _ __
| |  | |/ __| / _ \| | / _ \/ __|/ __| | |     / _ \  / __|| |/ // _ \| '__|
| |__| |\__ \|  __/| ||  __/\__ \\__ \ | |____| (_) || (__ |   <|  __/| |
 \____/ |___/ \___||_| \___||___/|___/ |______|\___/  \___||_|\_\\___||_|
```
Randomware-like sample that can be easily modified and used. For educational
purposes.

## Features

- Highly portable. Should run on Windows/Linux/BSD.
- Opens a webpage to notify the user.
- Sends encryption key to a server.
- Store any necessary files in a zip archive appended to the executable.
- Uses RSA + AES128/256 CBC PKCS7.
- Works offline.

## Usage

This software lacks several components that require a significant amount of
work to implement. This is left upto the reader to do.

- Remove the safety pin.
- Compile the software.
- Create a webserver capable of processing requests from useless locker.
- Generate a private/public key pair using openssl or similar.
- Create zip archive of all the files in the required hierarchy.
- Append zip archive onto executable.

## Appending the zip archive

Follow the instructions in the store [README](store/README.md).

## Legal

This is a fun peice of software that I enjoyed writing. I learned a few things
while writing this and I hope that others can too. Please do not use this for
ransomware or otherwise malicious purposes. This is only meant to be used for
educational purposes.
