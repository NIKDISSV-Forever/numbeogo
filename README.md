# Numbeo GUI

A desktop graphical user interface (GUI) application written in Go using the Fyne toolkit. It allows users to fetch,
filter, sort, and compare country-specific metrics and statistics from Numbeo.

## Features

- **Asynchronous Data Loading**: Concurrently fetches rankings and categories from the parser.
- **Dynamic Sorting**: Supports customizable positive and negative sorting weights (e.g., Crime Index vs. Quality of
  Life).
- **Country Filter**: Interactive checklist with search capability to focus on specific countries.
- **Automatic Region Detection**: Resolves the user's current region using the `ipinfo.io` API to set initial filters.
- **Lazy Loading**: Paginated data renderer with customizable page size for improved scrolling performance.

## Prerequisites

To build and run this application, you need:

- Go 1.18 or higher
- Graphics drivers (OpenGL) configured on your system (required by Fyne)

Refer to the [Fyne Installation Guide](https://developer.fyne.io/started/) for OS-specific dependencies (e.g., build
tools on Windows/Linux).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ndenissov/numbeogo.git
   cd numbeogo
   ```

2. Download Go modules:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run cmd/fyne/main.go
   ```

## Configuration

The application saves user preferences using the native Fyne Preferences API. You can configure:

- **Page Size**: Number of rendered items per view.
- **Year/Region**: Filter indicators retrieved from the Numbeo data structure.
- **Active Countries**: Filter out regions or nations you do not wish to compare.

## Dependencies

- [Fyne v2](https://fyne.io/) - Cross-platform GUI toolkit.
- [hashicorp/go-set](https://github.com/hashicorp/go-set) - Type-safe Set implementation for Go.
