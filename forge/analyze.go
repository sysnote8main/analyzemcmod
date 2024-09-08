package forge

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sysnote8main/analyzemcmod/common"
)

var (
	infoFileName = "mcmod.info" // TODO needs to change this if support other version
)

func ExtractInfoJsonFromJar(src string) (string, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return "", err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == infoFileName {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(rc)
			return buf.String(), nil
		}
	}
	return "", fmt.Errorf("not found %s in %s", infoFileName, src)
}

func GetModidFromJar(jarPath string) (*common.ModInfo, error) {
	s, err := ExtractInfoJsonFromJar(jarPath)
	if err != nil {
		return nil, err
	}
	var modInfo []common.ModInfo
	if err := json.Unmarshal([]byte(s), &modInfo); err != nil {
		return nil, err
	}
	if len(modInfo) != 1 {
		return nil, fmt.Errorf("failed to parse mod data")
	}
	return &modInfo[0], nil
}
