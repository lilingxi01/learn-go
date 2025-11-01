# File I/O in Go

Learn to read, write, and manipulate files in Go.

## Prerequisites

- Completed [11-testing](../11-testing/)
- Understanding of error handling

## What You'll Learn

- Reading files (entire file, line by line, buffered)
- Writing files
- Working with directories
- File permissions
- JSON encoding/decoding
- CSV operations
- Temporary files
- File metadata

## Reading Files

### Read Entire File

```go
data, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

### Read Line by Line

```go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

### Buffered Reading

```go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

reader := bufio.NewReader(file)
for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
        break
    }
    fmt.Print(line)
}
```

## Writing Files

### Write Entire File

```go
data := []byte("Hello, World!")
err := os.WriteFile("file.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

### Write with File Handle

```go
file, err := os.Create("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Hello, World!\n")
if err != nil {
    log.Fatal(err)
}
```

### Buffered Writing

```go
file, err := os.Create("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Buffered write\n")
writer.Flush()
```

## Working with Directories

```go
// Create directory
err := os.Mkdir("mydir", 0755)

// Create directory and parents
err := os.MkdirAll("path/to/mydir", 0755)

// Remove directory
err := os.Remove("mydir")

// Remove directory and contents
err := os.RemoveAll("mydir")

// List directory
entries, err := os.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name())
}
```

## File Operations

```go
// Check if file exists
_, err := os.Stat("file.txt")
if os.IsNotExist(err) {
    // File does not exist
}

// Rename/move file
err := os.Rename("old.txt", "new.txt")

// Copy file
src, _ := os.Open("source.txt")
defer src.Close()
dst, _ := os.Create("dest.txt")
defer dst.Close()
io.Copy(dst, src)

// Delete file
err := os.Remove("file.txt")
```

## JSON Operations

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Encode to JSON
person := Person{Name: "Alice", Age: 30}
data, err := json.Marshal(person)

// Decode from JSON
var p Person
err := json.Unmarshal(data, &p)

// Write JSON to file
file, _ := os.Create("data.json")
defer file.Close()
encoder := json.NewEncoder(file)
encoder.Encode(person)

// Read JSON from file
file, _ := os.Open("data.json")
defer file.Close()
decoder := json.NewDecoder(file)
var p Person
decoder.Decode(&p)
```

## Running the Examples

```bash
go run read-write.go
go run json.go
```

## File Permissions

Unix-style permissions (octal):
- `0644`: Owner read/write, group/others read
- `0755`: Owner all, group/others read+execute
- `0600`: Owner read/write only

## Best Practices

1. **Always defer Close()**: Ensure files are closed
2. **Check errors**: File operations can fail
3. **Use buffered I/O**: For large files
4. **Atomic writes**: Write to temp, then rename
5. **Use filepath.Join()**: For cross-platform paths

## Common Patterns

### Read Config File

```go
data, err := os.ReadFile("config.json")
if err != nil {
    log.Fatal(err)
}

var config Config
json.Unmarshal(data, &config)
```

### Safe Write (Atomic)

```go
tmp, err := os.CreateTemp("", "file-*.txt")
if err != nil {
    return err
}
defer os.Remove(tmp.Name())

if _, err := tmp.Write(data); err != nil {
    return err
}

if err := tmp.Close(); err != nil {
    return err
}

return os.Rename(tmp.Name(), "file.txt")
```

### Walk Directory Tree

```go
filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path)
    return nil
})
```

## Common Mistakes

1. **Not closing files**: Use defer
2. **Ignoring errors**: Check all error returns
3. **Not using buffered I/O**: Slow for large files
4. **Hardcoded paths**: Use filepath package
5. **Race conditions**: Multiple writers to same file

## Quick Reference

```go
// Read
data, err := os.ReadFile("file.txt")

// Write
err := os.WriteFile("file.txt", data, 0644)

// Open
file, err := os.Open("file.txt")
defer file.Close()

// Create
file, err := os.Create("file.txt")
defer file.Close()

// Append
file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)

// Check exists
_, err := os.Stat("file.txt")
if os.IsNotExist(err) {
    // Doesn't exist
}

// JSON
json.Marshal(v)
json.Unmarshal(data, &v)
```

## Next Steps

1. Run the file I/O examples
2. Try reading and writing different formats
3. Practice directory operations
4. Move to **13-http-basics** for HTTP programming

## Further Reading

- [os Package](https://pkg.go.dev/os)
- [io Package](https://pkg.go.dev/io)
- [bufio Package](https://pkg.go.dev/bufio)
- [encoding/json Package](https://pkg.go.dev/encoding/json)

