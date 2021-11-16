package main

import (
	"encoding/json"
	"fmt"
)

const (
	Large = iota
	Meduim
	Small
)

type Size int

func (s Size) marshalText() ([]byte, error) {
	switch s {
	case Large:
		return []byte("large"), nil
	case Meduim:
		return []byte("mediual"), nil
	case Small:
		return []byte("small"), nil
	default:
		return []byte("unkonwn"), nil
	}
}

func (s *Size) UnmarshalText(bytes []byte) error {
	switch string(bytes) {
	case "large":
		*s = Large
	case "medium":
		*s = Meduim
	case "small":
		*s = Small
	default:
		*s = Small
	}
	return nil
}

func main() {
	var size Size = Meduim
	bytes, _ := json.MarshalIndent(size, "", "\t")
	fmt.Println(string(bytes))					// 1


	var size02 Size
	json.Unmarshal(bytes,&size02)
	fmt.Println(size02)								// 0
	sizes := []Size{Large,Large,Small,Meduim}
	bytes ,_ = json.Marshal(sizes)

	fmt.Println(string(bytes))				//[0,0,2,1]





}
