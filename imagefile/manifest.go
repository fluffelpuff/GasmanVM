package imagefile

import (
	"strconv"
	"strings"
)

func (o *Manifest) GetManifestSourceFiles() []*FilePlayground {
	retrive := make([]*FilePlayground, 0)
	for _, i := range o.SourceFiles.Files {
		// Der Hash wird Gesplittet
		splited_hash := strings.Split(i.Hash, ":")
		if len(splited_hash) != 2 {
			return nil
		}

		// Die Nummer wird eingelesen
		num, err := strconv.ParseUint(splited_hash[0], 10, 32)
		if err != nil {
			return nil
		}

		// Der Eintrag wird hinzugefügt
		retrive = append(retrive, &FilePlayground{i.Path, i.Playground, splited_hash[1], uint32(num)})
	}
	return retrive
}
