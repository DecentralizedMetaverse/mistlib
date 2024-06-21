# FW (File World) CLI Documentation

## Introduction

FW (File World) is a command-line interface tool designed to manage and manipulate world data, objects, and content in a structured format using IPFS for distributed storage. This document provides an overview of the commands available in FW and their usage.

## Commands

### 1. Initialization

#### `fw init`

Initializes a new FW repository. This command sets up the necessary directory structure and configuration files.

**Usage:**
```bash
fw init
```

### 2. Handling Worlds

#### `fw get-world-info`

Fetches and consolidates the world information from the current world specified in the HEAD file. The consolidated data is saved as a YAML file in the content directory.

**Usage:**
```bash
fw get-world-info
```

#### `fw download-world <cid>`

Downloads world data and its associated content from IPFS using the specified CID. The downloaded data is decrypted, decompressed, and saved locally.

**Usage:**
```bash
fw download-world <cid>
```

#### `fw get-world-cid`

Calculates the SHA-256 hash of the current world file, compresses, encrypts it, and uploads it to IPFS. The resulting CID is then used to rename the local file and update the world data.

**Usage:**
```bash
fw get-world-cid
```

#### `fw switch <world-name>`

Switches to a different world by updating the HEAD file. If multiple worlds with the same name exist, the user is prompted to select one.

**Usage:**
```bash
fw switch <world-name>
```

### 3. Content Management

#### `fw get <cid> [filename]`

Downloads and decrypts a file from IPFS using the specified CID. Optionally, a filename can be provided to specify the local save location.

**Usage:**
```bash
fw get <cid> [filename]
```

#### `fw put <file> [x y z rx ry rz sx sy sz]`

Adds a new file to the world with specified coordinates, rotation, and scale. If only the file is specified, default values are used for the coordinates, rotation, and scale.

**Usage:**
```bash
fw put <file> [x y z rx ry rz sx sy sz]
```

### 4. Metadata Management

#### `fw update <meta-cid> <x y z rx ry rz sx sy sz>`

Updates the metadata for an existing content identified by `meta-cid` with new coordinates, rotation, and scale values.

**Usage:**
```bash
fw update <meta-cid> <x y z rx ry rz sx sy sz>
```

#### `fw set-custom-data <meta-cid> <value> [delete=true]`

Sets custom data for a specified metadata CID. Optionally, the old metadata can be retained or deleted.

**Usage:**
```bash
fw set-custom-data <meta-cid> <value> [delete=true]
```

#### `fw set-parent <child-cid> <parent-cid> [delete=true]`

Sets a parent-child relationship between two CIDs. Optionally, the old metadata can be retained or deleted.

**Usage:**
```bash
fw set-parent <child-cid> <parent-cid> [delete=true]
```

#### `fw remove <cid>`

Removes a specified content identified by CID from the world data and deletes the local file.

**Usage:**
```bash
fw remove <cid>
```

### 5. Utility Commands

#### `fw set-password <password>`

Sets a password for encrypting and decrypting files.

**Usage:**
```bash
fw set-password <password>
```

#### `fw cat <path>`

Displays the decrypted and decompressed content of a file.

**Usage:**
```bash
fw cat <path>
```

#### `fw unpack`

Unpacks the world data, decrypts, and decompresses all associated content files, and saves them locally.

**Usage:**
```bash
fw unpack
```

## Detailed Examples

### Example: Initializing a Repository
```bash
fw init
```
This command initializes a new FW repository by creating the necessary directories and files.

### Example: Adding a File with Default Coordinates
```bash
fw put example.txt
```
This command adds the file `example.txt` to the world with default coordinates, rotation, and scale values.

### Example: Switching to a Specific World
```bash
fw switch my-world
```
This command switches to the world named "my-world". If multiple worlds with this name exist, you will be prompted to select one.