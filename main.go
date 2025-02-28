package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"mini_reddis/constants"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	mu   sync.Mutex
	data map[string]string
)

func loadData() {
	_, err := os.Stat(constants.FileName)
	if err != nil {
		return
	}
	dat, err := os.ReadFile(constants.FileName)
	if err != nil {
		fmt.Println("Error: Could not read file.")
		return
	}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		fmt.Println("Error: Could not convert data")
		data = make(map[string]string)
	}
}

func saveData() {
	mu.Lock()
	defer mu.Unlock()
	dat, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting data to string: ", err)
		return
	}
	t := string(dat)
	err = os.WriteFile(constants.TempFileName, []byte(t), 0644)
	if err != nil {
		fmt.Println("Error persisting the data")
		return
	}
	err = os.Rename(constants.TempFileName, constants.FileName)
	if err != nil {
		fmt.Println("Error renaming persistance file")
	}
}

func execute(args []string, method string) string {
	mu.Lock()
	defer mu.Unlock()
	if len(args) == 1 && args[0] == "save" && method == "POST" {
		saveData()
		return "Data Saved"
	}
	if len(args) == 1 && method == "GET" {
		val, exists := data[args[0]]
		if !exists {
			fmt.Println(args[0], "key does not exists")
			return args[0] + " key does not exists"
		}
		return val
	}
	if len(args) == 1 && method == "DELETE" {
		_, exists := data[args[0]]
		if !exists {
			fmt.Println(args[0], "key does not exists")
			return args[0] + " key does not exists"
		}
		delete(data, args[0])
		return args[0] + " key deleted"
	}
	if len(args) == 2 && method == "POST" {
		data[args[0]] = args[1]
		return args[1]
	}
	return "Invalid Command"
}

func handler(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()
	reader := bufio.NewReader(conn)
	req, _ := reader.ReadString('\n')
	req = strings.TrimSpace(req)
	parts := strings.Split(req, " ")
	if len(parts) < 2 {
		return
	}
	fmt.Println(parts)
	method := parts[0]
	cmds := strings.Split(parts[1], "/")
	cmds = cmds[1:]
	if len(cmds) < 3 {
		fmt.Println("Invalid request format, requires atleast three args but", len(cmds), "were sent")
		conn.Write([]byte("Invalid request format, requires atleast three args and atmost 4 arguments"))
		return
	}
	if strings.TrimSpace(cmds[0]) != constants.UserName || strings.TrimSpace(cmds[1]) != constants.Password {
		fmt.Println("Invalid credentials!")
		conn.Write([]byte("Invalid credentials!"))
		return
	}
	args := strings.Split(cmds[2], "-")
	res0 := execute(args, method)
	res := "HTTP/1.1 200 OK\r\n" + "Content-Length: " + strconv.Itoa(len(res0)) + "\r\n" + "Content-Type: text/plain\r\n\r\n" + res0
	conn.Write([]byte(res))
}

func persistData() {
	for {
		time.Sleep(time.Minute)
		saveData()
	}
}

func main() {
	data = make(map[string]string)
	loadData()
	go persistData()
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(constants.Port))
	if err != nil {
		log.Fatal("Could not create server")
		return
	}
	defer ln.Close()
	fmt.Println("Server Created sucessfully, listening at ", constants.Port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Bad Request: ", err)
			continue
		}
		wg.Add(1)
		go handler(conn)
	}
	fmt.Println("Closing requests if any remaining")
	wg.Wait()
}
