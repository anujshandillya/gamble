# Use the official Go image as the base image
FROM golang:1.24.3-bookworm

# Install Air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Expose port 4000
EXPOSE 4000

# Command to run the application with Air for hot-reloading
CMD ["air", "-c", ".air.toml"] 