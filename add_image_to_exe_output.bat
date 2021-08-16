# to convert the png logo to ico file just go to https://icoconvert.com/ website and upload the logo and select all options in the Custom sizes radio, at last download the ico file and move/copy it to go/assets/ folder

@echo off 

set root-dir = %cd%
set release-outbut-dir = 
set assets-outbut-dir = %release-outbut-dir%\assets

%cd%\rh.exe -open %cd%\go\build\outputs\windows-release\remousable_plus.exe -save %cd%\go\build\outputs\windows-release\remousable_plus.exe -action addskip -res %cd%\go\build\outputs\windows-release\assets\icon.ico -mask ICONGROUP,MAINICON,