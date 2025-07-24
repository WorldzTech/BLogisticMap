package main

import (
	"blogistic/BLogisticMap"
	"fmt"
)

func main() {
	toEncrypt := []byte("This is a new way to create logistic map.")
	bm := BLogisticMap.NewBFlow(len(toEncrypt), 3.5, 0.56, 0.2, 0.8)

	streamer := BLogisticMap.NewStreamer(toEncrypt, bm)
	p := streamer.Encrypt()

	fmt.Println(p.Bytes)
	fmt.Println(p.AsString())

	unstreamer := BLogisticMap.NewStreamer(p.Bytes, bm)
	pr := unstreamer.Encrypt().AsString()
	fmt.Println(pr)
}