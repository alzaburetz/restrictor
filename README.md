# Restrictor

## Get Golang
<img src="https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn-images-1.medium.com%2Fmax%2F1200%2F1*mUjcwJ7INewkUIVWFJVRUA.jpeg&f=1" width=300>

To install **GO** v1.11.5, follow <a href="https://golang.org/doc/install" target="_new">this instructions</a> or you can use this bash code:
### Linux
```bash
wget -i https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz\
&& sudo mv go1.11.5.linux-amd64.tar.gz /usr/local\
&& sudo tar -xvf /usr/local/go1.11.5.linux-amd64.tar.gz\
&& export PATH=/usr/local/go/bin:$PATH 
&& sudo echo "export PATH=/usr/local/go/bin:$PATH" >> ~/.bashrc\
&& mkdir -p $HOME/go/src $$ cd $HOME/go/src\
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
```
wget -i https://dl.google.com/go/go1.11.5.darwin-amd64.pkg
sudo installer -pkg go1.11.5.darwin-amd64.pkg -target /usr/local/go
```
Or get [this file](https://dl.google.com/go/go1.11.5.darwin-amd64.pkg) and double-click it


### Windows
Get [this file](https://dl.google.com/go/go1.11.5.windows-amd64.msi) and double-click it



## Check-list:
- [x] Agent closes/opens restricted apps according to data
- [ ] API
- [ ] Database intagration
- [ ] Admin pannel

### Pre-pre-pre alpha stage
This is basic applcation, that runs silently, without any notifications and/or signals
Made easily with [this package](https://github.com/shirou/gopsutil)
```
go get -u https://github.com/shirou/gopsutil/...
```
To test this app, you can edit JSON Mocking-data file:
```json
{
 "app": "name_of_process_from terminal",
 "windows": 1,
 "rule": "Close for restricted Open for required",
 "time": "working for working days weekend otherwise",
 "hourfrom": "minimum hour of a day",
 "hourto": "maximum",
 "executable": "path/to/executable"
}
```

