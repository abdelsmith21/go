#!/bin/bash
# Tizen TV App Build Script
# This script packages the Tizen app into a .wgt file

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Tizen TV IPTV Player - Build Script ===${NC}\n"

# Check if tizen-app directory exists
if [ ! -d "tizen-app" ]; then
    echo -e "${RED}Error: tizen-app directory not found!${NC}"
    echo "Please run this script from the repository root."
    exit 1
fi

# Check if Tizen CLI is installed
TIZEN_CLI=""
if [ -f "$HOME/tizen-studio/tools/ide/bin/tizen" ]; then
    TIZEN_CLI="$HOME/tizen-studio/tools/ide/bin/tizen"
elif [ -f "/Applications/tizen-studio/tools/ide/bin/tizen" ]; then
    TIZEN_CLI="/Applications/tizen-studio/tools/ide/bin/tizen"
elif [ -f "C:/tizen-studio/tools/ide/bin/tizen.bat" ]; then
    TIZEN_CLI="C:/tizen-studio/tools/ide/bin/tizen.bat"
elif command -v tizen &> /dev/null; then
    TIZEN_CLI="tizen"
else
    echo -e "${YELLOW}Warning: Tizen CLI not found in standard locations.${NC}"
    echo "Please install Tizen Studio from:"
    echo "https://developer.samsung.com/tizen/tizen-studio/download"
    echo ""
    echo "If Tizen Studio is installed, please provide the path to the tizen command:"
    read -p "Tizen CLI path: " TIZEN_CLI
    
    if [ ! -f "$TIZEN_CLI" ] && ! command -v "$TIZEN_CLI" &> /dev/null; then
        echo -e "${RED}Error: Invalid Tizen CLI path${NC}"
        exit 1
    fi
fi

echo -e "${GREEN}Found Tizen CLI:${NC} $TIZEN_CLI\n"

# Create icon if it doesn't exist (placeholder)
if [ ! -f "tizen-app/icon.png" ]; then
    echo -e "${YELLOW}Warning: icon.png not found. Creating a placeholder...${NC}"
    
    # Check if ImageMagick is available
    if command -v convert &> /dev/null; then
        convert -size 512x512 xc:#1a1a2e -fill white -pointsize 72 -gravity center \
                -annotate +0+0 "IPTV" tizen-app/icon.png
        echo -e "${GREEN}Created placeholder icon.png${NC}\n"
    else
        echo -e "${YELLOW}ImageMagick not found. Please add icon.png manually.${NC}"
        echo "For now, creating a simple text file as placeholder..."
        echo "IPTV Player Icon" > tizen-app/icon.png
    fi
fi

# List available certificate profiles
echo -e "${GREEN}Step 1: Checking certificate profiles...${NC}"
PROFILES=$("$TIZEN_CLI" security-profiles list 2>&1 || echo "")

if [ -z "$PROFILES" ] || [[ "$PROFILES" == *"There is no profile"* ]]; then
    echo -e "${RED}Error: No certificate profiles found!${NC}"
    echo ""
    echo "You need to create a certificate profile first."
    echo "Please follow these steps:"
    echo ""
    echo "1. Open Tizen Studio"
    echo "2. Go to Tools → Certificate Manager"
    echo "3. Click '+' to create a new profile"
    echo "4. Select 'Samsung' → 'TV'"
    echo "5. Follow the wizard to create your certificate"
    echo ""
    echo "After creating the certificate, run this script again."
    exit 1
fi

echo -e "${GREEN}Available certificate profiles:${NC}"
echo "$PROFILES"
echo ""

# Ask for certificate profile
read -p "Enter certificate profile name (or press Enter for default): " CERT_PROFILE

if [ -z "$CERT_PROFILE" ]; then
    CERT_PROFILE=$(echo "$PROFILES" | grep -oP '^\s*\K\w+' | head -1)
    echo -e "${YELLOW}Using profile:${NC} $CERT_PROFILE"
fi

# Clean previous builds
echo -e "\n${GREEN}Step 2: Cleaning previous builds...${NC}"
rm -f tizen-app.wgt
rm -f tizen-app/.buildResult
rm -rf tizen-app/.build

# Build the package
echo -e "\n${GREEN}Step 3: Building .wgt package...${NC}"
cd tizen-app

if "$TIZEN_CLI" package -t wgt -s "$CERT_PROFILE" -- .; then
    cd ..
    
    # Find the generated .wgt file
    WGT_FILE=$(find tizen-app -name "*.wgt" -type f | head -1)
    
    if [ -n "$WGT_FILE" ]; then
        # Move to root directory
        mv "$WGT_FILE" ./tizen-app.wgt
        
        echo -e "\n${GREEN}=== Build Successful! ===${NC}"
        echo -e "${GREEN}WGT file created:${NC} tizen-app.wgt"
        echo -e "${GREEN}File size:${NC} $(du -h tizen-app.wgt | cut -f1)"
        echo ""
        echo -e "${YELLOW}Next steps:${NC}"
        echo "1. Enable Developer Mode on your TV (press 1-2-3-4-5 quickly in Apps)"
        echo "2. Connect your TV to the same network as your computer"
        echo "3. Use Tizen Studio Device Manager to install the app"
        echo "   OR"
        echo "4. Use SDB command:"
        echo "   ~/tizen-studio/tools/sdb connect <TV_IP>:26101"
        echo "   ~/tizen-studio/tools/sdb install tizen-app.wgt"
        echo ""
        echo "For detailed instructions, see: DEPLOYMENT_GUIDE_TR.md"
    else
        echo -e "${RED}Error: .wgt file not found after build${NC}"
        exit 1
    fi
else
    cd ..
    echo -e "${RED}Error: Build failed!${NC}"
    echo "Please check the error messages above."
    exit 1
fi
