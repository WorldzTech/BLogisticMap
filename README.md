# BLogisticMap Stream Cipher

A simple yet powerful stream cipher implementation in Go, based on a **modified logistic map** — a deterministic chaotic system adapted to improve the uniqueness and unpredictability of cryptographic sequences.

> 🚀 This project is backed by the paper
> _"Improving the Security of Cryptographic Keys Using Modified Chaotic Models"_
> by Daniil Baikalov (PDF available in repo).
>
> [📚 Read the paper](./Improving_the_Security_of_Cryptographic_Keys_Using_Modified_Chaotic_Models.pdf)

---

## 📌 Description

This project demonstrates the use of a **chaotically-modified logistic map** to generate byte streams for encryption. Unlike traditional chaotic systems, the model introduces dynamic multipliers and alternating modulation to drastically increase the uniqueness index of generated sequences, making them resistant to cryptanalysis.

It includes:

- Modified logistic map (`BFlow`)
- Byte-stream generator
- Simple stream cipher (`BStreamer`) using XOR encryption

---

## 🔐 Cryptographic Foundation

The algorithm is based on:

- **Logistic Map**:
  A classic chaotic model described by the recurrence
  `xₙ₊₁ = r * xₙ * (1 - xₙ)`

- **Modified with Alternating Multiplier**:
  `xₙ₊₁ = r * xₙ * (1 - xₙ) * dₙ`,
  where `dₙ` dynamically changes based on position `n`, and additional parameters `p`, `q`.

These modifications result in a **uniqueness index > 0.9**, compared to ~0.04 for the standard logistic map (on sequences of length 1000+).

---

## 🧠 How It Works

1. Generate a modified chaotic flow with `NewBFlow(n, r, x0, p, q)`
2. Encrypt message with `Encrypt()`
3. Get bytes stream or string representation from product by `Bytes` or `AsString()`

---

## ✈️ Example

```go
package main

import (
	"github.com/WorldzTech/BLogisticMap/BLogisticMap"
	"fmt"
)

func main() {
  // Creating a byte message and BLogistic Map
  toEncrypt := []byte("This is a new way to create logistic map.")
  bm := BLogisticMap.NewBFlow(len(toEncrypt), 3.5, 0.56, 0.2, 0.8)

  // Creating a stream cipher with specific BLogisticMap
  streamer := BLogisticMap.NewStreamer(toEncrypt, bm)
  p := streamer.Encrypt()

  fmt.Println(p.Bytes)
  fmt.Println(p.AsString())

  // Decrypting a message with XOR again
  unstreamer := BLogisticMap.NewStreamer(p.Bytes, bm)
  pr := unstreamer.Encrypt().AsString() // XOR again
  fmt.Println(pr)
}
```

---

## 📦 Installation

```bash
go get -u github.com/WorldzTech/BLogisticMap
```

---

## ✍️ Author

Daniil Baikalov

Email: felix.trof@gmail.com

Location: Moscow, Russia