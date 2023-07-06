# Doorbird Video Capture

This is a Go script that captures a video stream from a Doorbird Doorstation and saves it to a file. It is triggered via an HTTP GET request and runs in a Docker container.

## Prerequisites

- Docker
- Go (if running without Docker)
- ffmpeg (if running without Docker)

## Usage

### Running with Docker

1. Clone the repository: `git clone https://github.com/j0hannesr0th/doorbird-video-capture.git`
2. Navigate to the project directory: `cd doorbird-video-capture`
3. Build the Docker image: `docker build -t doorbird-video-capture .`
4. Run the Docker container: `docker run -p 8787:8080 -v $(pwd)/videos:/videos doorbird-video-capture`

### Running without Docker

1. Clone the repository: `git clone https://github.com/j0hannesr0th/doorbird-video-capture.git`
2. Navigate to the project directory: `cd doorbird-video-capture`
3. Install dependencies: `go mod tidy`
4. Run the Go script: `go run main.go`

## Configuration

The configuration file for the application is `config.json`. Edit this file to set the following parameters:

- `doorbird.ip`: The IP address of your Doorbird Doorstation.
- `doorbird.username`: Your Doorbird username.
- `doorbird.password`: Your Doorbird password.
- `record.duration`: The duration (in seconds) to record the video stream.
- `record.path`: The path where the recorded video will be saved.

## API Endpoint

To trigger the video recording, make an HTTP GET request to the following endpoint:

GET http://localhost:8787/record

## Recorded Videos

The recorded videos will be saved in the `videos` folder.

## License

See LICENSE file.
