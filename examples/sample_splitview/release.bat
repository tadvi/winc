rsrc -manifest app.manifest -ico=app.ico,add.ico,run.ico,edit.ico,error.ico -o rsrc.syso
go build -ldflags="-H windowsgui"
