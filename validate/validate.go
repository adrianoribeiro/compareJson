package validate

import (
	"compareJson/compare"
	"errors"
	"math"
)

//It is just to protect.
const limitPercentageSizeDiff = 50
func Exec(fileName, fileBkpName string) error{
	file, err := compare.OpenFile(fileName)
	if err != nil {
		return err
	}
	fileBkp, err := compare.OpenFile(fileBkpName)
	if err != nil {
		return err
	}

	fileInfo, _ := file.Stat()
	fileBkpInfo, _ := fileBkp.Stat()

	if fileInfo.Size() == 0 || fileBkpInfo.Size() == 0 {
		return errors.New("there are at least one empty file")
	}

	return validateSize(fileInfo.Size(), fileBkpInfo.Size())
}

func validateSize(sizeFile int64, sizeFileBkp int64) error{
	bigger := sizeFile
	if sizeFileBkp > bigger {
		bigger = sizeFileBkp
	}
	diffSize := math.Abs(float64(sizeFile - sizeFileBkp))

	percentageAllowed := 1 - (float64(100-limitPercentageSizeDiff) / 100)
	if diffSize > (float64(bigger) * percentageAllowed) {
		return errors.New("the diff size of the values is significant so they are different")
	}

	return nil
}
