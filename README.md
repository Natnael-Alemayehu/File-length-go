# File Length Reader

This Go program reads a file and returns the length (number of bytes) of its contents.

## How It Works

The program takes the file name as a command-line argument, opens the file, and reads its content. It then returns the number of bytes read. If the file is not specified or cannot be read, the program will return an error.

### Key Functions

- **getFile(name string)**: Opens a file and returns a pointer to the file along with a closure function for closing the file.
- **fileLen(name string)**: Reads the content of the file and returns the number of bytes.

### Error Handling

- If the file is not provided or cannot be opened, the program will return an error message.
- If an error occurs while reading the file, the program will stop execution and log the error.

## Installation

1. Clone the repository:

   ```bash
   git clone [https://github.com/your-username/your-repo-name.git](https://github.com/Natnael-Alemayehu/File-length-go.git)


Here's the complete README file in markdown style for your Go program:

2. Navigate to the project directory:

   ```bash
   cd File-length-go
   ```

3. Build the program:

   ```bash
   go build
   ```

## Usage

To use the program, specify the file you want to check as a command-line argument:

```bash
./filelength <file_name>
```

### Example

```bash
./filelength test.txt
```

This will output the number of bytes read from the file `test.txt`.

### Error Handling

- File not specified: If no file is provided as an argument, the program will print:
  
  ```bash
  File not specified
  ```

- **File cannot be opened**: If the specified file cannot be opened, the program will print:

  ```bash
  file not specified
  ```

- **File reading error**: If there is an issue while reading the file (other than reaching EOF), the program will exit with the error.

## Testing

You can test the program by creating a simple text file:

```bash
echo "Hello, world!" > test.txt
./filelength test.txt
```

This should output the number of bytes in the file (in this case, `13`).


## License

This project is licensed under the MIT License.

## Acknowledgments

- The Go programming language [documentation](https://golang.org/doc/)
