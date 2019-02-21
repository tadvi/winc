rsrc -manifest app.manifest -ico=app.ico,add.ico,application_lightning.ico,application_edit.ico,application_error.ico -o rsrc.syso
go build -ldflags="-H windowsgui"
