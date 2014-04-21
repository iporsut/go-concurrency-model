package main

import (
    "fmt"
)

type StackServerMessage struct {
    PushMessage chan int
    GetAllMessage chan (chan []int)
}

type StackServerState struct {
    List []int
    Top int
}

func ServerLoop(message StackServerMessage, state StackServerState) {
    for {
        select {
            case val := <-message.PushMessage:
                state.List[state.Top] = val
                state.Top = state.Top + 1
            case replyChan := <-message.GetAllMessage:
                replyChan <- state.List
        }
    }
}

func main() {
    pushMessage := make(chan int)
    getAllMessage := make(chan (chan []int))
    state := StackServerState { make([]int, 100), 0}
    message := StackServerMessage {pushMessage, getAllMessage}

    go ServerLoop(message, state)

    message.PushMessage <- 10
    message.PushMessage <- 20
    replyGetAll := make(chan []int, 100)
    message.GetAllMessage <- replyGetAll
    allValue := <-replyGetAll
    fmt.Println(allValue)
}
