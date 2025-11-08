@echo off
REM Tizen TV App Build Script for Windows
REM This script packages the Tizen app into a .wgt file

echo === Tizen TV IPTV Player - Build Script (Windows) ===
echo.

REM Check if tizen-app directory exists
if not exist "tizen-app" (
    echo Error: tizen-app directory not found!
    echo Please run this script from the repository root.
    pause
    exit /b 1
)

REM Check if Tizen CLI is installed
set TIZEN_CLI=
if exist "C:\tizen-studio\tools\ide\bin\tizen.bat" (
    set TIZEN_CLI=C:\tizen-studio\tools\ide\bin\tizen.bat
) else if exist "%USERPROFILE%\tizen-studio\tools\ide\bin\tizen.bat" (
    set TIZEN_CLI=%USERPROFILE%\tizen-studio\tools\ide\bin\tizen.bat
) else (
    echo Warning: Tizen CLI not found in standard locations.
    echo Please install Tizen Studio from:
    echo https://developer.samsung.com/tizen/tizen-studio/download
    echo.
    set /p TIZEN_CLI="Enter the full path to tizen.bat: "
    if not exist "!TIZEN_CLI!" (
        echo Error: Invalid Tizen CLI path
        pause
        exit /b 1
    )
)

echo Found Tizen CLI: %TIZEN_CLI%
echo.

REM Create icon placeholder if it doesn't exist
if not exist "tizen-app\icon.png" (
    echo Warning: icon.png not found. Creating a placeholder...
    echo IPTV Player Icon > tizen-app\icon.png
    echo Created placeholder icon.png
    echo.
)

REM Step 1: Check certificate profiles
echo Step 1: Checking certificate profiles...
call "%TIZEN_CLI%" security-profiles list > profiles.tmp 2>&1

findstr /C:"There is no profile" profiles.tmp >nul
if %errorlevel% equ 0 (
    del profiles.tmp
    echo Error: No certificate profiles found!
    echo.
    echo You need to create a certificate profile first.
    echo Please follow these steps:
    echo.
    echo 1. Open Tizen Studio
    echo 2. Go to Tools - Certificate Manager
    echo 3. Click '+' to create a new profile
    echo 4. Select 'Samsung' - 'TV'
    echo 5. Follow the wizard to create your certificate
    echo.
    echo After creating the certificate, run this script again.
    pause
    exit /b 1
)

echo Available certificate profiles:
type profiles.tmp
echo.

REM Ask for certificate profile
set /p CERT_PROFILE="Enter certificate profile name (or press Enter for default): "

if "%CERT_PROFILE%"=="" (
    for /f "tokens=1" %%a in (profiles.tmp) do (
        set CERT_PROFILE=%%a
        goto :got_profile
    )
)
:got_profile

del profiles.tmp
echo Using profile: %CERT_PROFILE%
echo.

REM Step 2: Clean previous builds
echo Step 2: Cleaning previous builds...
if exist tizen-app.wgt del tizen-app.wgt
if exist tizen-app\.buildResult del tizen-app\.buildResult
if exist tizen-app\.build rmdir /s /q tizen-app\.build
echo.

REM Step 3: Build the package
echo Step 3: Building .wgt package...
cd tizen-app

call "%TIZEN_CLI%" package -t wgt -s "%CERT_PROFILE%" -- .

if %errorlevel% equ 0 (
    cd ..
    
    REM Find the generated .wgt file
    for %%f in (tizen-app\*.wgt) do (
        move "%%f" tizen-app.wgt
        goto :build_success
    )
    
    echo Error: .wgt file not found after build
    pause
    exit /b 1
    
    :build_success
    echo.
    echo === Build Successful! ===
    echo WGT file created: tizen-app.wgt
    for %%A in (tizen-app.wgt) do echo File size: %%~zA bytes
    echo.
    echo Next steps:
    echo 1. Enable Developer Mode on your TV (press 1-2-3-4-5 quickly in Apps)
    echo 2. Connect your TV to the same network as your computer
    echo 3. Use Tizen Studio Device Manager to install the app
    echo    OR
    echo 4. Use SDB command:
    echo    C:\tizen-studio\tools\sdb connect [TV_IP]:26101
    echo    C:\tizen-studio\tools\sdb install tizen-app.wgt
    echo.
    echo For detailed instructions, see: DEPLOYMENT_GUIDE_TR.md
    echo.
) else (
    cd ..
    echo Error: Build failed!
    echo Please check the error messages above.
    pause
    exit /b 1
)

pause
