# Findr

Findr is a location-based web application that helps users discover and connect with others in their vicinity. Built with Go and modern web technologies, it provides an interactive map interface for exploring nearby locations.

## Features

- Interactive map interface using Leaflet.js
- User authentication system

## Prerequisites

- Go 1.16 or higher
- Supabase account and credentials

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
SUPABASE_URL=your_supabase_url
SUPABASE_KEY=your_supabase_key
SERVER_PORT=8080  # Optional, defaults to 8080
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/findr.git
cd findr
```

2. Install Go dependencies:
```bash
go mod download
```

3. Set up environment variables (see above)

4. Run the application:
```bash
go run server.go
```

The application will be available at `http://localhost:8080`

## Project Structure

```
findr/
├── client/             # Frontend files
│   ├── index.html     # Main application page
│   ├── login.html     # Login page
│   ├── styles.css     # CSS styles
│   └── script.js      # Frontend JavaScript
├── server.go          # Main Go server file
├── go.mod            # Go module file
└── README.md         # This file
```

## API Endpoints

- `GET /` - Serves the main application page
- `GET /login` - Serves the login page
- `POST /login` - Handles user login
- `GET /api/users` - Retrieves nearby users
- `POST /api/users` - Updates search radius

## Development

The project uses Air for live reloading during development. To start the development server:

```bash
air
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Acknowledgments

- [Leaflet.js](https://leafletjs.com/) for the interactive map
- [Gin](https://gin-gonic.com/) for the Go web framework
- [Supabase](https://supabase.io/) for the backend services 