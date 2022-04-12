# SWkshp
SWkshp installs mods from the Steam Workshop to mod folders of games owned via platforms other than steam. I don't endorse piracy, for use only if you own the game.

## Installation
### Through GitHub Releases
Head to the [latest release](https://github.com/Breadinator/swkshp/releases/latest) and download `swkshp.exe` if you're on Windows.

### From source
It can be installed directly from github using go's built-in tools.
```bash
go install github.com/breadinator/swkshp
```

## Usage
Assumes you're using Windows and the place it is installed to is in the PATH. If this is not the case, adjust accordingly.

### Setting download path
Each game has to have a path to download to manually set. This is done using the command `swkshp.exe config game <gamename> <path/to/dir>`.

An example of this would be
```bash
swkshp.exe config game rimworld "D:/Program Files/RimWorld/Mods"
```

### Installing a mod
After setting the download path for a game, it is easy to install a mod. The command is as follows:
```bash
swkshp.exe <url to workshop page> [-g <gamename>]
```
The -g flag is used to define the game. By default it reads the web page to find out what game it's for, so this flag is often unneccesary.

Let's say we want to install [HugsLib](https://steamcommunity.com/workshop/filedetails/?id=818773962). The command to install it would be:
```bash
swkshp.exe https://steamcommunity.com/workshop/filedetails/?id=818773962
```

### Installing a collection
Installing entire collections can be done in the same way as installing a mod. The following example would let you install the [Vanilla Expanded collection](https://steamcommunity.com/workshop/filedetails/?id=1884025115) for RimWorld.
```bash
swkshp.exe https://steamcommunity.com/workshop/filedetails/?id=1884025115
```
This command looks through the web page, makes a list of every listed item on it then downloads them.

I'm not sure how reliable the check to see if it's a collection is, so I might have to improve it or at least add a manual flag if necessary.

### Removing a mod
To remove a mod, use the command:
```bash
shwkshp.exe remove <url> [-g <gamename>]
```
This command will delete all associated files and remove its entry from the internal 

## To-do
* Improve the version tracking, and test more extensively
    * Improve how files are updated
* Compile to Linux & Mac
* Add extra config options
* Streamline