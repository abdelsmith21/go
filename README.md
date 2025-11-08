# Tizen TV IPTV Player

A modern IPTV player application for Samsung Tizen TVs with Xtream Codes API support.

## Features

- ✅ **Modern User Interface** - Beautiful gradient design optimized for TV screens
- ✅ **Live TV** - Watch live channels with full EPG support
- ✅ **VOD (Movies)** - Browse and watch movies on demand
- ✅ **TV Series** - Access and watch TV series episodes
- ✅ **Categories** - Organized content by categories
- ✅ **Zap Buttons** - Quick channel switching with remote control CH+/CH- buttons
- ✅ **Subtitle Support** - Select and switch between subtitle tracks
- ✅ **Audio Track Selection** - Choose from multiple audio tracks
- ✅ **Remote Control Navigation** - Full support for TV remote control
- ✅ **Colored Button Support** - Quick access via Red/Green/Yellow/Blue buttons

## Architecture

The application consists of two main components:

### 1. Go Backend Server (`main.go`)
- Handles Xtream Codes API communication
- Provides RESTful API endpoints
- Proxies media streams
- Manages authentication

### 2. Tizen TV Web Application (`tizen-app/`)
- HTML5-based user interface
- JavaScript application logic
- Remote control navigation system
- Video player with subtitle/audio support

## Prerequisites

- Go 1.21 or higher
- Samsung Tizen TV (6.0+) or Tizen Studio for development
- Xtream Codes IPTV server credentials

## Installation & Setup

### 1. Backend Server Setup

```bash
# Clone the repository
git clone https://github.com/abdelsmith21/go.git
cd go

# Download dependencies
go mod download

# Run the server
go run main.go
```

The server will start on `http://localhost:8080`

### 2. Configuration

Edit `main.go` to configure your Xtream Codes credentials:

```go
baseURL := "http://your-xtream-server.com:8080"
username := "your_username"
password := "your_password"
```

Alternatively, you can set these via environment variables:

```bash
export XTREAM_URL="http://your-server.com:8080"
export XTREAM_USER="your_username"
export XTREAM_PASS="your_password"
```

### 3. Tizen TV Deployment

#### Option A: Using Tizen Studio

1. Install [Tizen Studio](https://developer.samsung.com/tizen/tizen-studio)
2. Import the `tizen-app` folder as a project
3. Connect to your TV (Tools → Device Manager)
4. Build and deploy the application

#### Option B: Manual Installation

1. Package the `tizen-app` folder as a `.wgt` file
2. Upload to your TV using the Tizen Studio or web-based installer
3. Install and run the application

## Usage

### First Time Setup

1. Launch the application on your Tizen TV
2. Enter your Xtream Codes credentials:
   - Server URL (e.g., `http://server.com:8080`)
   - Username
   - Password
3. Click "Login"

### Navigation

**Remote Control Keys:**
- **Arrow Keys** - Navigate between menu items and channels
- **Enter/OK** - Select item
- **Back** - Go back to previous screen
- **CH+ / CH-** - Quick channel zapping in live TV mode
- **Play/Pause** - Control video playback
- **Stop** - Stop playback and return
- **Red Button** - Toggle subtitles
- **Green Button** - Toggle audio tracks
- **Yellow Button** - Info/EPG (reserved for future)
- **Blue Button** - Favorites (reserved for future)

### Features

#### Live TV
- Browse channels by category
- Quick channel switching with CH+/CH- buttons
- Real-time streaming

#### Movies (VOD)
- Browse movies by category
- High-quality on-demand playback
- Movie information display

#### TV Series
- Browse series by category
- Episode selection
- Binge-watching support

#### Player Controls
- Play/Pause
- Previous/Next channel/episode
- Subtitle selection
- Audio track selection
- Return to menu

## API Endpoints

The Go backend provides the following REST API endpoints:

- `GET /api/auth` - Authenticate with Xtream Codes server
- `GET /api/live/categories` - Get live TV categories
- `GET /api/live/streams?category_id=X` - Get live streams
- `GET /api/vod/categories` - Get VOD categories
- `GET /api/vod/streams?category_id=X` - Get VOD streams
- `GET /api/series/categories` - Get series categories
- `GET /api/series/streams?category_id=X` - Get series
- `GET /api/stream/{id}?type=X&ext=Y` - Proxy stream

## Project Structure

```
.
├── main.go                 # Go backend server
├── go.mod                  # Go module dependencies
├── README.md              # This file
└── tizen-app/             # Tizen TV application
    ├── config.xml         # Tizen app configuration
    ├── index.html         # Main HTML file
    ├── icon.png           # App icon
    ├── css/
    │   └── style.css      # Application styles
    └── js/
        ├── navigation.js  # Remote control navigation
        ├── api.js         # API client
        └── app.js         # Application logic
```

## Development

### Testing Locally

You can test the web interface in a browser:

1. Start the Go server:
   ```bash
   go run main.go
   ```

2. Open a browser and navigate to:
   ```
   http://localhost:8080
   ```

### Building for Production

Build the Go server for your target platform:

```bash
# For Linux
GOOS=linux GOARCH=amd64 go build -o iptv-server main.go

# For Windows
GOOS=windows GOARCH=amd64 go build -o iptv-server.exe main.go

# For macOS
GOOS=darwin GOARCH=amd64 go build -o iptv-server main.go
```

## Security Considerations

**Credential Storage:**
The application stores IPTV credentials in the browser's localStorage for convenience. This is standard practice for IPTV applications as:
- Credentials are for streaming services, not sensitive personal/financial data
- The app runs locally on your TV in a sandboxed environment
- This prevents users from re-entering credentials on every launch

**Best Practices:**
- Only use this app on trusted devices (your personal TV)
- Don't share your TV with untrusted users
- Use strong, unique passwords for your IPTV service
- Consider using a separate IPTV account if sharing the TV

## Troubleshooting

### Issue: Cannot connect to Xtream Codes server
- Verify your server URL is correct and accessible
- Check firewall settings
- Ensure credentials are correct

### Issue: Video not playing
- Check if your TV supports the stream format
- Try a different stream/channel
- Check network connectivity

### Issue: Remote control not working
- Ensure you're using a compatible Samsung Tizen TV
- Check if the app has focus
- Try pressing the Back button to activate

### Issue: Subtitles not showing
- Ensure the stream has embedded subtitles
- Try selecting a different subtitle track
- Check if subtitles are enabled in TV settings

## Future Enhancements

- [ ] EPG (Electronic Program Guide) integration
- [ ] Favorites management
- [ ] Parental controls
- [ ] Multi-profile support
- [ ] Catch-up TV support
- [ ] Recording functionality
- [ ] Advanced search and filtering

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions, please open an issue on GitHub.

## Acknowledgments

- Built with Go and vanilla JavaScript
- Designed for Samsung Tizen TV platform
- Compatible with Xtream Codes API

---

**Note:** This application requires valid Xtream Codes IPTV credentials to function. Make sure you have legal access to the IPTV service you're using.