package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	file := `E:\CodeHere\qtp.provision\qtp\QTP\FrontEnd\SdFrontEnd\App\UsageReporter_Server\UsageDataMerge\usagedata\UFT_14.52\1ccea399-2601-4a58-a179-315e76930974\470eb787-d850-463d-86ec-580e6347a82d.dat`
	fmt.Println(filepath.Dir(file) + "\n")
	fmt.Println(filepath.Dir(filepath.Dir(file)))
	fmt.Println(filepath.Base(filepath.Dir(filepath.Dir(file))))
}
