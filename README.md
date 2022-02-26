# Timer Convert
## Download
Download [latest release](https://github.com/Blueil/Timer_Convert/releases/latest).
## Use
1. Export from twisty timer to other timers and get the file on the pc or connect the phone and browse to it.
2. Start `converter.exe` and select the file or use
```posh
converter.exe path_to_file.txt
```
3. It will ask in the console: `How much solves do you want to convert from x solves?`. Answer how much of the latest solves you want to import, if you want to import all just type the x.
4. Open CubeDesk, then goto Settings > Data > (`Recommended`: Click export all data in case something goes wrong) > Import Data > Import From CsTimer.
5. Select cube type, import the generated `output.csv` file and then click proccess file and import data
6. Goto Sessions and select as the current one the one you want to add the solves to.
7. Click on `csTimer Session` and then merge into current.

You are done!

## Building
Install go and use
```posh
go build main.go
```
To build the exe file.
