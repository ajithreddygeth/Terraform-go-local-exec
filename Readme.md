This is the example for using go binary to run using terraform 

Advantages:
1. No need for any dependencies
2. You can incluse anything to run in the binary

How to use:
1. Put all of the thing that needs to run while doing terrraform apply
2. build the main file and store it on /bin/{utils}
3. yiou can build the binary using build.sh with current directory as *utils*
```
$ ls
bin  build.sh src
./build.sh
```

Dirctory structure
```
.
├── Readme.md
├── main.tf
├── utils
│   ├── bin
│   │   └── utils
│   ├── build.sh
│   └── src
│       └── main.go
└── variables.tf
```

Feature ehnacement :
1. Use the docker command to build the binary (refer the build file )
#docker run -v  $(pwd):/app/ -e CGO_ENABLED=0 -w /app/src/  golang:1.14 go build -a -o /app/bin/utils .
