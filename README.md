# remousable_plus

A new Flutter project, that use golang to execute 

## Getting Started

This project is a starting point for a Flutter application.

A few resources to get you started if this is your first Flutter project:

- [Lab: Write your first Flutter app](https://flutter.dev/docs/get-started/codelab)
- [Cookbook: Useful Flutter samples](https://flutter.dev/docs/cookbook)

For help getting started with Flutter, view our
[online documentation](https://flutter.dev/docs), which offers tutorials,
samples, guidance on mobile development, and a full API reference.


## Add icon to the output exe file
1. convert the .png logo to .ico file just go to https://icoconvert.com/ website and upload your .png logo then select all boxes in the Custom sizes radio, at last download the .ico file, rename it to icon.ico, and move or copy it to go/assets/ folder.
2. install Resource Hacker from http://www.angusj.com/resourcehacker/
3. copy the ResourceHacker.exe file from C:\Program Files (x86)\Resource Hacker\ to your project directory, and rename it to rh.exe
4. in the .gitignore file in your project add rh.ini
5. create new file named build_windows_release.bat and add this code inside it:
@echo off 

hover build windows --release
 & 'C:\Program Files (x86)\Resource Hacker\ResourceHacker.exe' -open %cd%\go\build\outputs\windows-release\<<<your file output name here>>>.exe -save %cd%\go\build\outputs\windows-release\<<<your file output name here>>>.exe -action addskip -res %cd%\go\build\outputs\windows-release\assets\icon.ico -mask ICONGROUP,MAINICON,
%cd%\rh.exe -open %cd%\go\build\outputs\windows-release\<<<your file output name here>>>.exe -save %cd%\go\build\outputs\windows-release\<<<your file output name here>>>.exe -action addskip -res %cd%\go\build\outputs\windows-release\assets\icon.ico -mask ICONGROUP,MAINICON,
6. replace <<<your file output name here>>> with your output .exe file name.
7. run .\build_windows_release.bat in cmd in your project directory.
8. done ✓ :)

now every time you want to build your project just run .\build_windows_release.bat, and you would have the .exe output file with your favorite icon