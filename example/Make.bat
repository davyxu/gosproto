set CURR_DIR=%cd%

: Build generator
cd ..\..\..\..\..
set GOPATH=%cd%
go build -o %CURR_DIR%\sprotogen.exe github.com/davyxu/gosproto/sprotogen
cd %CURR_DIR%

: Generate go source file by sproto
sprotogen --type=go --out=addressbook.go --package=example --cellnet_reg=true addressbook.sp

: Convert to standard sproto file
sprotogen --type=sproto --out=addressbook.sproto addressbook.sp

: Generate c# source file by sproto
sprotogen --type=cs --out=addressbook.cs --package=example addressbook.sp

: Generate lua source file by sproto
sprotogen --type=lua --out=addressbook.lua --package=example addressbook.sp