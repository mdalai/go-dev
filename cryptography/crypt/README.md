# crypt

```sh
./crypt hash test.txt main.go
./crypt hash -sha256 test.txt
./crypt hash -str "hello world"
./crypt hash -str "hello world" -sha256

# Encrypt a string to string output
./crypt encrypt -str "hello world" 
# Encrypt a string to bytes output
./crypt encrypt -str "hello world" -toByte
# Encrypt a file
./crypt encrypt hello.txt hello.txt.enc key.file

# Decrypt a byte string, the output is "hello world". Can also output to bytes by using "-toByte".
./crypt decrypt -str "[119 171 45 78 180 7 48 205 149 76 24 168 109 74 246 254 196 218 103 218 250 144 126 139 62 110 196 82 126 145 144 234 167 218 212 199 148 103 229]"
# Decrypt a file to std output
./crypt decrypt hello.txt.enc key.file
# Decrypt a file to another file
./crypt decrypt hello.txt.enc hello.txt.dec key.file
```

## FAQ
1. Read file with `os.open()` VS `ioutil.ReadFile()`? 