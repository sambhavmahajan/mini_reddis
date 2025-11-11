# Mini Redis
This is a redis clone, but with less features. It utilizes a TCP server running by default on 5233 port.
## Prerequisite
- Golang
- Netcat or Postman

## Featues:
- Reteriving value from key
- Inserting key value pair
- Deleting key value pair
- Auto Saving: Data is auto persisted every 1 minute
- Manual Saving
- Atomic Saving: File is first persisted to a temp file, then it is persisted to save file

## configuration file: constants/constants.go
```
const (
	Port = 5233
	FileName = "save.sav"
	TempFileName = "tmp.sav"
	UserName = "sambhav"
	Password = "sambhav"
)
```
## Running
```
go run main.go
```
## API calls (use netcat)
Request Format
```
METHOD /<username>/<password>/<args>
```
delimited by new line character
<args> are separated by -:
- For POST (set data): POST /user/pass/key-value
- For GET (fetch data): GET /user/pass/key
- For DELETE (delete data): DELETE /user/pass/key
- To persist manually: POST /user/pass/save
## Example request
On terminal use netcat
```
echo "POST /sambhav/sambhav/name-sambhav\n" | ncat localhost 5233
```
<img width="571" height="69" alt="image" src="https://github.com/user-attachments/assets/04b590a2-7624-4c09-8971-002485b530e6" /><br>
*Figure 1: Terminal Screenshot*

## Running using docker
```
docker build -t mini_reddis .
docker run -P mini_reddis
```
