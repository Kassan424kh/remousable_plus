@echo off

hover build windows --release >nul 2>&1
%cd%\rh.exe -open %cd%\go\build\outputs\windows-release\remousable_plus.exe -save %cd%\go\build\outputs\windows-release\remousable_plus.exe -action addskip -res %cd%\go\build\outputs\windows-release\assets\icon.ico -mask ICONGROUP,MAINICON,
cls
echo Successfully compiled executable binary for windows
