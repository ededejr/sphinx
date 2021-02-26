# Sphinx
A CLI Token Generator implemented in Go. The generated token is automatically copied to your clipboard.

<p align="center">
  <img src="./sphinx.png" width="200px" height="200px" />
</p>

## Usage
The binary is included in source so you can download just the `sphinx` file if you choose.

### Flags
- `--characters`: The types of characters allowed in the token, see below for more
- `--length`: The length of the token
- `--display`: Print the token in the terminal instead of putting it in your clipboard
- `--separator`: Determines what to use the delineator for separating chunks
- `--chunks`: The number of chunks the Token is split into

### Platform Specific
#### Unix
1. Open terminal and go to the project folder
2. Run with `./sphinx`

#### Windows
1. You need to build the program for windows using `go`
2. Run with `sphinx.exe`

> You can also just click on the file. The output is copied to your clipboard by default.

## Appendix

### Characters
- `A`: Uppercase
- `a`: Lowercase
- `0`: Numerics
- `!`: Special characters
- `*`: Any of the above
