package metainfo

import (
	"bittorrent/bencode"
	"fmt"
	"os"
)

type Metainfo struct {
	Announce string
	Info     Info
}

type Info struct {
	Name        string
	PieceLength int
	Pieces      string
	Length      int
	Files       []File
}

type File struct {
	Path   []interface{}
	Length int
}

func DecodeMetainfo(path string) (*Metainfo, error) {
	encoded, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", path, err)
	}

	decoded, err := bencode.Decode(string(encoded))
	if err != nil {
		return nil, fmt.Errorf("error decoding bencode data: %v", err)
	}

	metainfo, err := validateMetainfo(decoded)
	if err != nil {
		return nil, fmt.Errorf("error validating metainfo: %v", err)
	}

	return metainfo, nil
}

func validateMetainfo(metainfoData interface{}) (*Metainfo, error) {
	metainfo, ok := metainfoData.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("metainfo is invalid")
	}

	announce, ok := metainfo["announce"].(string)
	if !ok || announce == "" {
		return nil, fmt.Errorf("announce field is empty or invalid")
	}

	info, err := validateInfo(metainfo["info"])
	if err != nil {
		return nil, fmt.Errorf("error validating info field: %v", err)
	}

	metainfoStruct := &Metainfo{
		Announce: announce,
		Info:     *info,
	}
	return metainfoStruct, nil
}

func validateInfo(infoData interface{}) (*Info, error) {
	info, ok := infoData.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("info field is invalid")
	}

	name, ok := info["name"].(string)
	if !ok || name == "" {
		return nil, fmt.Errorf("name field is empty or invalid")
	}

	pieceLength, ok := info["piece length"].(int)
	if !ok || pieceLength <= 0 {
		return nil, fmt.Errorf("piece length field is empty or invalid")
	}

	pieces, ok := info["pieces"].(string)
	if !ok || pieces == "" || len(pieces) % 20 != 0 {
		return nil, fmt.Errorf("pieces field is empty or invalid")
	}

	length := 0
	files := []File{}
	if info["length"] != nil && info["files"] == nil {
		length, ok := info["length"].(int)
		if !ok || length <= 0 {
			return nil, fmt.Errorf("length field is empty or invalid")
		}
	} else if info["length"] == nil && info["files"] != nil {
		var err error
		files, err = validateFiles(info["files"])
		if err != nil {
			return nil, fmt.Errorf("error validating files field: %v", err)
		}
	} else {
		return nil, fmt.Errorf("exactly one of length or files must be specified")
	}

	return &Info{
		Name:        name,
		PieceLength: pieceLength,
		Pieces:      pieces,
		Length:      length,
		Files:       files,
	}, nil

}

func validateFiles(filesData interface{}) ([]File, error) {
	filesDataSlice, ok := filesData.([]interface{})
	if !ok {
		return nil, fmt.Errorf("files field is empty or invalid")
	}

	files := []File{}
	for i, fileData := range filesDataSlice {
		file, ok := fileData.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("file[%d] is empty or invalid", i)
		}

		path, ok := file["path"].([]interface{})
		if !ok || len(path) == 0 {
			return nil, fmt.Errorf("file[%d] path is empty or invalid", i)
		}

		length, ok := file["length"].(int)
		if !ok || length <= 0 {
			return nil, fmt.Errorf("file[%d] length is empty or invalid", i)
		}

		files = append(files, File{
			Path:   path,
			Length: length,
		})
	}
	return files, nil
}
