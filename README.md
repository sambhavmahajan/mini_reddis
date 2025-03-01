# Mini Reddis
This is a reddis clone, but with less features. It utilizes a TCP server running by default on 5233 port.

## Featues:
- Reteriving value from key: Get Request with endpoint /username/password/key
- Inserting key value pair: Post Request with endpont /username/password/key-value
- Deleting key value pair: Delete Request with endpoint /username/password/key
- Auto Saving: Data is auto persisted every 1 minute
- Manual Saving: Post request with endpoint /username/password/save
- Atomic Saving: File is first persisted to a temp file, then it is persisted to save file

## configuration file: constants/constants.go
- Port = 5233
- FileName = "save.sav"
- TempFileName = "tmp.sav"
- UserName = "sambhav"
- Password = "sambhav"

## Running
```
go run main.go
```
