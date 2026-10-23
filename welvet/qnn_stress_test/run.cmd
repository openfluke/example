@echo off
REM Optional — plain "go run ." already bootstraps PATH + INPROC.
cd /d "%~dp0"
go run .
exit /b %ERRORLEVEL%