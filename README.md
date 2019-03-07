# Restrictor

## Get Golang
<img src="https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fih1.redbubble.net%2Fimage.38303832.2081%2Fraf%2C360x360%2C075%2Ct%2Cfafafa%3Aca443f4786.jpg&f=1" width=300>

To install **GO** v1.12.0, follow <a href="https://golang.org/doc/install" target="_blank">this instructions</a> or you can use this bash code:
### Linux
```bash
wget -i https://dl.google.com/go/go1.12.linux-amd64.tar.gz\
&& sudo mv go1.12.linux-amd64.tar.gz /usr/local\
&& sudo tar -xvf /usr/local/go1.12.linux-amd64.tar.gz\
&& sudo rm /usr/local/go1.12.linux-amd64.tar.gz\
&& export PATH=/usr/local/go/bin:$PATH\ 
&& sudo echo "export PATH=/usr/local/go/bin:$PATH" >> ~/.bashrc\
&& mkdir -p $HOME/go/src/test $$ cd $HOME/go/src/test\
&& echo "package main
 import \"fmt\"
 func main() {
 fmt.Println(\"Hello world\")
 }" > test.go \
&& go run test.go
```

### MacOS
[Homebrew](https://brew.sh/) is required to install from terminal
```bash
brew install wget
```
```bash
wget -i https://dl.google.com/go/go1.12.darwin-amd64.pkg\
&& sudo installer -pkg go1.12.darwin-amd64.pkg -target /usr/local/go\
&& sudo rm go1.12.darwin-amd64.pkg
```
Or get [this file](https://dl.google.com/go/go1.11.5.darwin-amd64.pkg) and double-click it


### Windows
Get [this file](https://dl.google.com/go/go1.11.5.windows-amd64.msi) and double-click it



## Check-list:
- [x] Agent closes/opens restricted apps according to data
- [x] [API](https://github.com/alzaburetz/restrictor-api.git)
- [x] Database intagration
- [ ] Admin pannel

### Pre-pre-pre alpha stage
This is basic applcation, that runs silently, without any notifications and/or signals
Made easily with [this package](https://github.com/shirou/gopsutil)
```bash
go get -u https://github.com/shirou/gopsutil/...
```

### API test
Go to [API repository](https://github.com/alzaburetz/restrictor-api.git)

