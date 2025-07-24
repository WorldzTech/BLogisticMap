package main

import (
	"blogistic/BLogisticMap"
	"fmt"
)

func main() {
//	generator := BLogisticMap.NewGenerator(3.5, 0.56, 0.2, 0.8, 10)
//	generator.WarmUp(100)
//
//	i := 100
//	for i > 0 {
//		i--
//		fmt.Println(generator.Next())
//	}

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