# Directory tree
Simple Go module to print directory tree structures in terminal.  

# Launch example

```
$ go run main.go testdata/ -f
├───project
│   ├───file.txt (19b)
│   └───gopher.png (70372b)
├───static
│   ├───a_lorem
│   │   ├───dolor.txt (empty)
│   │   ├───gopher.png (70372b)
│   │   └───ipsum
│   │       └───gopher.png (70372b)
│   ├───css
│   │   └───body.css (28b)
│   ├───empty.txt (empty)
│   ├───html
│   │   └───index.html (57b)
│   ├───js
│   │   └───site.js (10b)
│   └───z_lorem
│       ├───dolor.txt (empty)
│       ├───gopher.png (70372b)
│       └───ipsum
│           └───gopher.png (70372b)
├───zline
│   ├───empty.txt (empty)
│   └───lorem
│       ├───dolor.txt (empty)
│       ├───gopher.png (70372b)
│       └───ipsum
│           └───gopher.png (70372b)
└───zzfile.txt (empty)

```

# Project Goals
The code is written for educational purposes. Training course ["Разработка веб-сервисов на Go"](https://www.coursera.org/learn/golang-webservices-1)