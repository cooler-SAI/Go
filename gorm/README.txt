1 - Get MinGW-W64-binaries:
https://github.com/niXman/mingw-builds-binaries?tab=readme-ov-file

Download the installer for the mingw-w64 version appropriate for your system (most likely x86_64 for 64-bit systems).
Run the installer and follow the prompts to install MinGW-w64. Make sure to select the gcc component during installation.
Add MinGW-w64 to PATH:

2 - Open the Control Panel and go to System and Security -> System -> Advanced system settings.
Click on "Environment Variables".
In the "System variables" section, find the Path variable, select it, and click "Edit".
Add the path to the bin directory of your MinGW-w64 installation (e.g., C:\mingw-w64\bin).
Click "OK" to close all dialog boxes.

3 - Get GORM packages: 

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite


TIP:  Binary was compiled with ‘CGO_ENABLED=0’, go-sqlite3 requires cgo to work. - ERROR!!!!!

Using Goland:

You can check the value with

go env CGO_ENABLED


and update it wtih 

go env -w CGO_ENABLED=<0 or 1>

example: go env -w CGO_ENABLED=1