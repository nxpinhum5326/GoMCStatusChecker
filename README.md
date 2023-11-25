# GoMCStatusChecker
Minecraft server status / online players checker written in Golang with mc-rcon library

## Installation 
``` git clone https://github.com/nxpinhum5326/GoMCStatusChecker.git ```
after cloning repo, configure rcon settings:
- err := conn.Open("IP:PORT", "rconPassword")

then ``` go build main.go ```
