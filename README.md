# Setup

### git

Install git: https://git-scm.com/downloads

Choose your operating system and install git by following the installation instructions. You may need to restart your terminal for changes to take effect.

### Golang

Please choose the latest stable version (1.15.2) for your operating system here: https://golang.org/dl/

##### `GOROOT` and `GOPATH`

If you are having difficulties with downloading and importing libraries, the most likely reason is a misconfigured `GOROOT` and/or `GOPATH`. 

The `GOROOT` points to the Golang installation location and defaults to `/usr/local/go` for Unix and `C:\Go` for Windows if it is not set. If you install Golang anywhere else, you must set the `GOROOT` environmental variable.

`GOPATH` lists locations where Golang looks for Go code. It is used to resolve import statements. Directories listed as `GOPATH` must have the subdirectories `src`, `bin`, and `pkg`, and the `src` contains the Golang modules.

Follow this link for more information about `GOPATH`: https://golang.org/cmd/go/#hdr-GOPATH_environment_variable

Follow this link for more information about setting `GOPATH`: https://github.com/golang/go/wiki/SettingGOPATH

### Text editors

You may choose any text editor of your choosing:

 - VS Code: https://code.visualstudio.com/download
 - Sublime Text: https://www.sublimetext.com/3
 - Notepad++: https://notepad-plus-plus.org/downloads/
 - Atom: https://atom.io/
 - Vim: https://www.vim.org/download.php

You may also use a text editor not present in this list. Choose one that you find most comfortable to use.

### Starting

Open terminal to a directory where you would like to download this repository and run the following command:

```
git clone https://github.com/BearCloud/proj0.git
```
