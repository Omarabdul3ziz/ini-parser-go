### Tasks
- [x] Read
- [x] Write
- [x] Test
- [ ] Handle the syntax error in the file


---
### Test sample
```ini
[Profile]
name = jarvis

# credential
password = secret

[Deployment]
project = peertube
name = peertest
public_ip = true
cpu = 4
memory = 8192
; rootfs = 1
```

```go
/*
    - parse function takes reader
    - read function takes (string/file) raturn reader
*/

/*
    - bufio.NewScanner() > return a scanner
    - scanner.Split(split function) > devide the tokens in the stream
    - scanner.Scan() > advance to the next token unless stoped by err/EOF
    - scanner.Text() > return the most recent token
*/
/*
    The returned value of type *os.File implements the io.Reader interface.
*/
```