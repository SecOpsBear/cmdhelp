# CMD helper cli

This is a simple cli tool to store commands in a sqlite database which is written in GO.   

## A simple cli tool to store and retrieve data such as example commands  

The commands are stored in a sqlite database. User can add various example commands for each individual tool. 

## Demo  

Below is a sample demo on storing commands to the database. More detailed usage is available at [SecOpsBear cmdhelp](https://secopsbear.com/tools/bear_tools/cmdhelp) 

![Add commands to database](https://secopsbear.com/assets/images/cmdhelp-add-ae8a088a089b900d405849930cbe7bc1.gif)  

## How to build   

### Install in linux cmdhelp

```console
go install github.com/secopsbear/cmdhelp@latest
```

To generate binaries Makefile is provided

### Build for linux

```console
make build_linux
```

### Build for window

Generate an executable **`cmdhelp.exe`** for windows environment.   

```console
make build_windows
```


## Find a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above.