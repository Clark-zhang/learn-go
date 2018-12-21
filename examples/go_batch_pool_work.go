package main

import (
	"github.com/golang/glog"
)

type data struct {
}

func main() {
	offset := 0
	pageSize := 1000
	sema := make(chan struct{}, 100)

	for {
		dataset, err := getData(offset, pageSize)
		if err != nil {
			glog.Info("fetch data failed")
			continue
		}
		if len(dataset) == 0 {
			glog.Info("finished")
			break
		}

		for _, item := range dataset {
			// acquire token
			sema <- struct{}{}
			go func(item data) {
				// do job
				// release semaphone
				<-sema
			}(item)
		}

		offset += pageSize
	}
}

func getData(offset int, pageSize int) (items []data, err error) {
	return items, nil
}
