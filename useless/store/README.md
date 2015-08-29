# go-store

[![go-store](https://godoc.org/github.com/TheCreeper/go-store?status.png)](http://godoc.org/github.com/TheCreeper/go-store)

Package store attempts to open the currently running or specified executable as a zip file. It also provides a means to easily find and access files within the archive.

## Examples

Open the currently running executable file.
```Go
s, err := Open()
if err != nil {
	panic(err)
}
defer s.Close()

file, err := s.Load("helloworld.txt")
if err != nil {
	panic(err)
}
println(string(file))
```

Open a specified executable file.
```Go
s, err := OpenFile("myfile")
if err != nil {
	panic(err)
}
defer s.Close()

file, err := s.Load("helloworld.txt")
if err != nil {
	panic(err)
}
println(string(file))
```

## How to Use

Fist create a zip archive
```
zip my.zip my.file
```

Compile your code
```
go build -o myprogram
```

Append the zip archive to the end of the executable
```
cat my.zip >> myprogram
```

Fix the zip offset in the file
```
zip -q -A myprogram
```