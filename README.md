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

We will assume that you have git and Golang properly set up, and that you are logged into Github.

**Do not run `git clone` on this repository.** Instead perform the following:

1. First create a repository in Github (**not** git) via this link: https://github.com/new

2. Create a new repository using git by running the following commands in a directory of your choosing:
```
git init
git branch -M master
git remote add origin <YOUR GITHUB REPO LINK HERE>
git push -u origin master
```
A variation of these commands are present in the "...or create a new repository on the command line" section after you create a repository on Github.

3. Add this project's repository as remote and alias it to `source`:
```
git remote add source https://github.com/BearCloud/proj0.git
```

4. Run the following command to pull starter code:
```
git pull source master
```

5. Run the following command to install the required Golang packages:
```
go mod download
```

6. You are done with setup.
