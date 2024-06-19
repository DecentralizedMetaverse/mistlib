# fw Command Line Documentation

The `fw` executable can be run in two modes: Command Line Mode and HTTP Server Mode. This document focuses on the Command Line Mode, where various commands can be executed directly from the terminal.

## Command Line Mode

When executed with command-line arguments, `fw` performs specific operations as defined by the arguments. The general syntax for running `fw` in Command Line Mode is:

```
fw <command> [arguments...]
```

### Commands

Here is a list of available commands along with their descriptions and usage:

1. **init**
   - **Description**: Initializes the repository.
   - **Usage**: 
     ```
     fw init
     ```

2. **switch**
   - **Description**: Switches to a specified world.
   - **Usage**: 
     ```
     fw switch <worldName>
     ```
   - **Arguments**:
     - `<worldName>`: The name of the world to switch to.

3. **get**
   - **Description**: Retrieves a content identifier (CID).
   - **Usage**: 
     ```
     fw get <cid>
     ```
   - **Arguments**:
     - `<cid>`: The content identifier to retrieve.

4. **put**
   - **Description**: Puts content into the repository.
   - **Usage**: 
     ```
     fw put <args...>
     ```
   - **Arguments**:
     - `<args...>`: The arguments for the put operation.

5. **set-password**
   - **Description**: Sets a password.
   - **Usage**: 
     ```
     fw set-password <password>
     ```
   - **Arguments**:
     - `<password>`: The password to set.

6. **cat**
   - **Description**: Concatenates and displays the content of a file hash.
   - **Usage**: 
     ```
     fw cat <fileHash>
     ```
   - **Arguments**:
     - `<fileHash>`: The hash of the file to display.

7. **get-world-cid**
   - **Description**: Retrieves the CID of the current world.
   - **Usage**: 
     ```
     fw get-world-cid
     ```

8. **download-world**
   - **Description**: Downloads the specified world.
   - **Usage**: 
     ```
     fw download-world <cid>
     ```
   - **Arguments**:
     - `<cid>`: The CID of the world to download.

9. **get-world-data**
   - **Description**: Retrieves the data of the current world.
   - **Usage**: 
     ```
     fw get-world-data
     ```

### Example Usage

Here are some example commands and their expected operations:

- **Initialize the repository**:
  ```
  fw init
  ```

- **Switch to a world named "example_world"**:
  ```
  fw switch example_world
  ```

- **Retrieve content with CID "example_cid"**:
  ```
  fw get example_cid
  ```

- **Put content into the repository**:
  ```
  fw put arg1 arg2 arg3
  ```

- **Set the password to "example_password"**:
  ```
  fw set-password example_password
  ```

- **Display the content of a file with hash "example_file_hash"**:
  ```
  fw cat example_file_hash
  ```

- **Retrieve the CID of the current world**:
  ```
  fw get-world-cid
  ```

- **Download a world with CID "example_cid"**:
  ```
  fw download-world example_cid
  ```

- **Retrieve the data of the current world**:
  ```
  fw get-world-data
  ```

By following these command structures, you can efficiently use the `fw` executable in Command Line Mode to perform various operations related to repository and world management.