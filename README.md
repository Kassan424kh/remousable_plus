# remousable_plus

GUI To run [remouseable](https://github.com/Kassan424kh/remouseable) on windows/linux or macos, and connect to the remarkable tablet with only one click, after adding ip-address and the tablet password.

## configuration
1. install go from https://golang.org/dl/
2. install hover https://github.com/go-flutter-desktop/hover
3. install c compiler from https://github.com/golang/go/wiki/InstallFromSource#install-c-tools

## Run the Project
```CMD
> hover run
```

## build release
On Windows:
```CMD
> .\build_windows_release.bat
```
Other Plattforms:
```SH
$ hover build <Platform_Name_here> --release
```

<br/>
<br/>

## Add icon to the output exe file
1. convert the .png logo to .ico file just go to https://icoconvert.com/ website and upload your .png logo then select all boxes in the Custom sizes radio, at last download the .ico file, rename it to icon.ico, and move or copy it to go/assets/ folder.
2. install Resource Hacker from http://www.angusj.com/resourcehacker/
3. copy the ResourceHacker.exe file from C:\Program Files (x86)\Resource Hacker\ to your project directory, and rename it to rh.exe
4. in the .gitignore file in your project add rh.ini
5. create new file named build_windows_release.bat and add this code inside it:
```BAT

@echo off

hover build windows --release >nul 2>&1
%cd%\rh.exe -open %cd%\go\build\outputs\windows-release\remousable_plus.exe -save %cd%\go\build\outputs\windows-release\remousable_plus.exe -action addskip -res %cd%\go\build\outputs\windows-release\assets\icon.ico -mask ICONGROUP,MAINICON,
cls
echo Successfully compiled executable binary for windows

```
6. replace <<<your file output name here>>> with your output .exe file name.
7. run .\build_windows_release.bat in cmd in your project directory.
8. done ✓ :)

now every time you want to build your project just run .\build_windows_release.bat, and you would have the .exe output file with your favorite icon