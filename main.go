package main

import "fmt"

import "bittorrent/bencode"

func main() {
	fmt.Println(bencode.Encode("spam"))
	fmt.Println(bencode.Encode(42))
	fmt.Println(bencode.Encode([]interface{}{"foo", 42, "bar"}))
	fmt.Println(bencode.Encode(map[string]interface{}{"key1": "value1", "key2": 2}))

	fmt.Println(bencode.Decode("4:spam"))
	fmt.Println(bencode.Decode("i42e"))
	fmt.Println(bencode.Decode("l4:spami42e4:eggse"))
	fmt.Println(bencode.Decode("le"))
	fmt.Println(bencode.Decode("d4:eggs5:toast4:spami42ee"))
}
