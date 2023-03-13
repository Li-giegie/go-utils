package go_utils

import (
	"io/fs"
	"os"
	"time"
)

type FileInfo struct {
	Name string
	Size int64
	Mode os.FileMode
	ModTime time.Time
	IsDir bool
	err error
}

// 获取给定目录内的所有文件、目录信息,入参_path 路径，filter 过滤器，过滤掉不想获取的目录、文件，返回值false表示过滤掉
func GetDirInfo(_path string,filter ...func(info os.FileInfo) (pass bool)) ([]*FileInfo, error) {
	var dirs = make([]*FileInfo,0)
	var info fs.FileInfo
	if _path[len(_path)-1] != '/' { _path += "/" }

	dirEntry,err := os.ReadDir(_path)
	if err != nil {
		return dirs,err
	}

	for _, entry := range dirEntry {
		info,err = entry.Info()
		if err != nil {
			dirs = append(dirs,&FileInfo{
				Name:    _path + entry.Name(),
				err:     err,
			} )
			continue
		}

		if len(filter) > 0 && !filter[0](info) {
			continue
		}

		dirs = append(dirs,&FileInfo{
			Name:    _path + info.Name(),
			Size:    info.Size(),
			Mode:    info.Mode(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
			err:     nil,
		})
		if entry.IsDir() {
			tmpInfo ,tmpErr  := GetDirInfo(dirs[len(dirs)-1].Name,filter...)
			if tmpErr != nil {
				return dirs,tmpErr
			}
			dirs = append(dirs, tmpInfo...)
			continue
		}
	}
	return dirs,nil
}