# Restrictor

## Get Golang
<img src="https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn-images-1.medium.com%2Fmax%2F1200%2F1*mUjcwJ7INewkUIVWFJVRUA.jpeg&f=1">
### To install go, follow [this instructions] (https://golang.org/doc/install) or you can use this bash code:
- Linux
```
wget -i https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz\
&& mv go1.11.5.linux-amd64.tar.gz /usr/local\
&&  tar -xvf go1.11.5.linux-amd64.tar.gz\
&& sudo echo "export PATH=/usr/local/go/bin:$PATH" >> ~/.bashrc\
&& mkdir $HOME/go
```

### Pre-pre-pre alpha stage
This is basic applcation, that runs silently, without any notifications and/or signals
Made easily with [This package] (https://github.com/shirou/gopsutil)
```
go get -u https://github.com/shirou/gopsutil
```
