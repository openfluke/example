@echo off
setlocal
REM Prepend welvet.qnn DLL dir so libwelvet_qnn.dll loads (cmd.exe, not PowerShell).
set "QNN_DLL=%~dp0..\..\..\welvet\native\qnn\windows_arm64"
if not exist "%QNN_DLL%\libwelvet_qnn.dll" (
  echo missing: %QNN_DLL%\libwelvet_qnn.dll
  exit /b 1
)
set "PATH=%QNN_DLL%;%PATH%"
cd /d "%~dp0"
go run .
exit /b %ERRORLEVEL%
