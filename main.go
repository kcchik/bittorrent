package main

import (
	"fmt"
	"bittorrent/metainfo"
)

func main() {
	meta, err := metainfo.DecodeMetainfo("./testdata/Welcome to the NHK.torrent")
	if err != nil {
		fmt.Errorf("Error decoding metainfo:", err)
		return
	}
	fmt.Printf("Announce: %s\n", meta.Announce)
	fmt.Printf("Info Name: %s\n", meta.Info.Name)
	fmt.Printf("Info Piece Length: %d\n", meta.Info.PieceLength)
	fmt.Printf("Info Length: %d\n", meta.Info.Length)
	for i, file := range meta.Info.Files {
		fmt.Printf("File %d: Path: %s, Length: %d\n", i+1, file.Path, file.Length)
	}

	info, err := metainfo.EncodeInfo(meta.Info)
	if err != nil {
		fmt.Errorf("Error encoding info:", err)
		return
	}
	fmt.Println(info)
}
