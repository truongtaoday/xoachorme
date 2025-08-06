@echo off
setlocal

echo === uninstall GOOGLE CHROME ===
echo.

REM End all Chrome processes
taskkill /f /im chrome.exe >nul 2>&1
taskkill /f /im GoogleUpdate.exe >nul 2>&1

REM uninstall GOOGLE CHROME (if exists)
wmic product where "name like 'Google Chrome%%'" call uninstall /nointeractive >nul 2>&1

REM Delete installation folder and user data
rmdir /s /q "%ProgramFiles%\Google\Chrome"
rmdir /s /q "%ProgramFiles(x86)%\Google\Chrome"
rmdir /s /q "%LOCALAPPDATA%\Google\Chrome"
rmdir /s /q "%APPDATA%\Google\Chrome"

REM Delete registry key (safe without Chrome)
reg delete "HKEY_CURRENT_USER\Software\Google\Chrome" /f >nul 2>&1
reg delete "HKEY_LOCAL_MACHINE\Software\Google\Chrome" /f >nul 2>&1
reg delete "HKEY_LOCAL_MACHINE\Software\WOW6432Node\Google\Chrome" /f >nul 2>&1

echo.
echo === DOWNLOAD THE LATEST GOOGLE CHROME ===

REM Create temporary folder
set TEMP_DIR=%TEMP%\chrome_installer
mkdir "%TEMP_DIR%" >nul 2>&1
cd /d "%TEMP_DIR%"

REM Download Chrome installation file from official site
powershell -Command "Invoke-WebRequest -Uri https://dl.google.com/chrome/install/latest/chrome_installer.exe -OutFile chrome_installer.exe"

echo.
echo === INSTALL GOOGLE CHROME ===
start /wait chrome_installer.exe /silent /install

echo.
echo Done. Google Chrome has been reinstalled.
pause
endlocal
