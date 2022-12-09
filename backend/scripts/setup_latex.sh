#!/bin/bash

#### Download latex class from the author Git repository
wget -q "https://raw.githubusercontent.com/liantze/AltaCV/main/altacv.cls"

### Download TinyTex for minimal latex installation
wget -qO- "https://yihui.name/gh/tinytex/tools/install-unx.sh" | sh

### Set PATH variable
export PATH=$PATH:$HOME/bin

### Install the necessary packages
### Packages based on my installation in an Arch linux machine
tlmgr install pgf fontawesome5 koma-script cmap ragged2e everysel tcolorbox \
    enumitem ifmtarg dashrule changepage multirow environ paracol lato \
    fontaxes accsupp extsizes pdfx colorprofiles xmpincl adjustbox collectbox
