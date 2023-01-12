# golang-playground
Dedicated branch for practicing golang.

# Installing Go
## For windows, macOS, or Linux
Installation file can be download [here](https://go.dev/dl/)

## For WSL 2
- Donwload and install Go by executing this command below on your WSL terminal:
```
wget https://dl.google.com/go/<latest_version>.linux-amd64.tar.gz
sudo tar -xvf <latest_version>.linux-amd64.tar.gz
sudo mv go /usr/local
```
*NOTE*: Change the <latest_version> with Go latest version on [here](https://go.dev/dl/) e.g `go1.19.5`
- Add this line the your shell configuration file (.bashrc, .zshrc, etc)
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
