package lib

import (
    "strings"
    glog "github.com/golang/glog"
    "github.com/nu7hatch/gouuid"
    "time"
    "strconv"
)

func RenameFileToUniqueId(s string) (n string, err error){
    splits := strings.Split(s, ".")

    randomUuid, err := uuid.NewV4()
    if err != nil {
        glog.Fatalf("RenameFileToUniqueId failed: %s", err)
        return
    }

    glog.Info(string(time.Now().Unix()))
    glog.Info(splits[len(splits)-1])

    n = randomUuid.String() + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "." + splits[len(splits)-1]

    return
}