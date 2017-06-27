package common

import (
    "github.com/gnampfelix/pub"
)

func FailureMessage() pub.Message{
    result := NewMessage("result")
    result.Write([]byte("FAILURE"))
    return result
}

func ErrorMessage() pub.Message {
    result := NewMessage("result")
    result.Write([]byte("ERROR"))
    return result
}

func SuccessMessage() pub.Message {
    result := NewMessage("result")
    result.Write([]byte("SUCCESS"))
    return result
}
