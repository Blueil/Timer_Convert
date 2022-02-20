# Timer Convert
## Download
Download [latest release](https://github.com/Blueil/Timer_Convert/releases/latest).
## Use
1. Export from twisty timer to other timers and get the file on the pc or connect the phone and browse to it.
2. Start `converter.exe` and select the file or use
```posh
converter.exe path_to_file.txt
```
3. Open CubeDesk, then goto Settings > Data > (`Recommended`: Click export all data in case something goes wrong) > Import Data > Import From CsTimer.
4. Select cube type, import the generated `output.csv` file and then click proccess file and import data
5. Goto Sessions and select as the current one the one you want to add the solves to.
6. Click on `csTimer Session` and then merge into current.

You are done!

## Building
Install go and use
```posh
go build main.go
```
To build the exe file.
