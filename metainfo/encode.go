package metainfo

import (
	"fmt"
	"bittorrent/bencode"
)

func EncodeInfo(info Info) (string, error) {
	infoMap := map[string]interface{}{
		"name":         info.Name,
		"piece length": info.PieceLength,
		// "pieces":       info.Pieces,
	}

	if info.Length > 0 {
		infoMap["length"] = info.Length
	} else if len(info.Files) > 0 {
		files := []interface{}{}
		for _, file := range info.Files {
			path := []interface{}{}
			for _, part := range file.Path {
				path = append(path, part)
			}
			files = append(files, map[string]interface{}{
				"path":   path,
				"length": file.Length,
			})
		}
		infoMap["files"] = files
	}

	fmt.Printf("path: %T\n", infoMap["files"].([]interface{})[0].(map[string]interface{})["path"])

	fmt.Printf("files: %T\n\n", infoMap["files"])

	return bencode.Encode(infoMap)
}
