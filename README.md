# b2h (Bytes to Human)

`b2h` is a simple command-line tool to convert bytes into human-readable formats like KB and MB.

## Installation

1. Build the binary:
   ```bash
   make build
   ```
2. Install it to `/usr/local/bin`:
   ```bash
   sudo make install
   ```

## Usage

```bash
b2h <numOfBytes> -u [KB|MB|GB|TB]
```

### Examples
```bash
# Convert 123456789 bytes to MB
b2h 123456789
# Output: 117.74 MB

# Convert 123456789 bytes to GB
b2h 123456789 -u GB
# Output: 0.1149780946 GB
```
