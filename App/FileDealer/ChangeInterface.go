package FileDealer

import "sync"

func SaveInterface(ListOfInterfaces *[]SWCInterfaces, wg *sync.WaitGroup, currentInterface chan SWCInterfaces, NUMBER_OF_FILE int) {
	i := 0
	for i < NUMBER_OF_FILE {

		*ListOfInterfaces = append(*ListOfInterfaces, <-currentInterface)
		i++
	}
	wg.Done()
}

//func ChangeInterface() {
//
//}
