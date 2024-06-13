
package main

/*
#include <stdlib.h>
#include "websocket.h"
*/
import "C"
import (
    "fmt"
    "sync"

    "github.com/gorilla/websocket"
)

var (
    conn  *websocket.Conn
    mutex sync.Mutex
)

//export ConnectWebSocket
func ConnectWebSocket(url *C.char) {
    u := C.GoString(url)
    var err error
    conn, _, err = websocket.DefaultDialer.Dial(u, nil)
    if err != nil {
        fmt.Println("Failed to connect:", err)
    }
    go readMessages()
}

//export SendMessage
func SendMessage(message *C.char) {
    mutex.Lock()
    defer mutex.Unlock()
    msg := C.GoString(message)
    err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
    if err != nil {
        fmt.Println("Failed to send message:", err)
    }
}

func readMessages() {
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("read:", err)
            return
        }
        fmt.Println("Received message:", string(message))
    }
}

func main() {
}
