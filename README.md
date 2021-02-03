# Sphinx
A CLI Token Generator implemented in Go. The generated token is automatically copied to your clipboard.

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
1. Open terminal
2. `./sphinx`

#### Windows
1. Open CMD to wherever you put the downloaded file
2. `sphinx.exe`

> You can also just click on the file. The output is copied to your clipboard by default.

### Character Index
- `A`: Uppercase
- `a`: Lowercase
- `0`: Numerics
- `!`: Special characters
- `*`: Any of the above
