# Trace - A Simple Version Control System

[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8.svg?logo=go)](https://golang.org)

**Trace: Learn the heart of version control through a lightweight, Git-inspired system.**


> **Note**: Trace is not a production-ready replacement for Git. It’s a learning tool to explore VCS concepts.



## Overview
Trace is a minimalistic version control system (VCS) built in Go, designed for educational and experimental purposes. It mimics core Git functionality, allowing users to initialize repositories, stage files, and create commits using SHA-1 content hashing. Trace stores all data in a human-readable JSON format within a `.trace/` directory, making it ideal for understanding how version control systems work under the hood.

## Features
- **Repository Initialization**: Set up a new VCS repository with `trace init`.
- **File Staging**: Add files or directories to the staging area with `trace add`.
- **Commit System**: Create commits with metadata and file hashes using `trace commit`.
- **JSON-Based Storage**: Stores data in `.trace/` for transparency and debugging.
- **Planned Status Command**: (In development) Check changes since the last commit with `trace status`.


## Installation
### Requirements 
- Go 
### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/trace.git
   ```
2. Navigate to the project directory:
   ```bash
   cd trace
   ```
3. Build and install Trace:
   ```bash
   go build -o trace
   go install
   ```
4. Verify installation:
   ```bash
   trace --help
   ```

## Usage
### Initialize a Repository
Start tracking a project by creating a `.trace/` directory:
```bash
trace init
```

### Stage Files
Add files or directories to the staging area (stored in `.trace/index.json`):
```bash
trace add .          # Stage all files in the current directory
trace add file.txt   # Stage a specific file
```

### Create a Commit
Save staged changes with a commit message:
```bash
trace commit -m "Initial commit"
```


## Directory Structure
```
trace/
├── .trace/           # Trace metadata directory
│   ├── objects/      # Hashed file content blobs
│   ├── commits/      # Commit objects with metadata
│   └── index.json    # Staging area for tracked files
├── main.go           # Source code for Trace
└── README.md         # This file
```

## How It Works
- **Initialization**: `trace init` creates a `.trace/` folder with `objects/`, `commits/`, and `index.json`.
- **Staging**: `trace add` hashes file contents using SHA-1 and stores them in `objects/`, updating `index.json`.
- **Committing**: `trace commit` creates a commit object(.json) in `commits/`, linking to staged file hashes and metadata.
- **Storage**: All data is stored in JSON for readability, making Trace a great tool for learning VCS internals.
