package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"os"
	"path/filepath"
)

const picDit = `\Packages\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\LocalState\Assets`

var sourceDict = fmt.Sprintf("%s%s", os.Getenv("LOCALAPPDATA"), picDit)

var distDict = fmt.Sprintf("%s%s", os.Getenv("USERPROFILE"), `\Pictures\`)

//var distDict = `C:\Users\X\Pictures\Saved Pictures`

func main() {

	creatDic(distDict + "big")
	creatDic(distDict + "small")
	getfile(sourceDict)
	fmt.Println("win10 聚焦图片提取结束")

	//tests := `C:\Users\X\AppData\Local\Packages\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\LocalState\Assets\1`
	//dists := `C:\Users\X\Pictures\Saved Pictures\big\1.jpeg`
	//copyFile(tests, dists)
}

func creatDic(path string) {

	exist, err := PathExists(path)
	if err != nil {
		fmt.Println(err)
	}

	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func getfile(path string) {
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			//fmt.Println(path)
			picHandle(path)
		}
		return nil
	})
}

func picHandle(source string) {
	file, err := os.Open(source)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	imageAttr, typeAttr, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
	}

	_, fileName := filepath.Split(source)
	f := fmt.Sprintf("%s.%s", fileName, typeAttr)
	mill := ""
	if imageAttr.Height > imageAttr.Width {
		mill = `small\`

	} else {
		mill = `big\`

	}
	dist := fmt.Sprintf("%s%s%s", distDict, mill, f)
	//dist := fmt.Sprintf("%s%s", mill, f)
	//fmt.Println(dist)

	copyFile(source, dist)

}

func copyFile(source, dist string) {
	img, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(dist, img, 0777)

	if err != nil {
		fmt.Println(err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
