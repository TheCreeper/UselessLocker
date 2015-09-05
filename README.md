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
- Uses RSA + AES128/256 CBC PKCS7.
- Sends encryption key to a server.
- Works offline.
- Store any necessary files in a zip archive appended to the executable.
- Opens a webpage to notify the user.
- Uses only the go standard library.

## Requirements

- Linux/BSD/cygwin enviroment
- Golang
- zip

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

Fist create a zip archive
```
zip -r my.zip assets/
```

Append the zip archive to the end of the executable
```
cat my.zip >> myprogram
```

Fix the zip offset in the file
```
zip -q -A myprogram
```

## Legal

This is a fun peice of software that I enjoyed writing. I learned a few things
while writing this and I hope that others can too. Please do not use this for
ransomware or otherwise malicious purposes. This is only meant to be used for
educational purposes.
