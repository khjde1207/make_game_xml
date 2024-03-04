package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/djherbis/atime"
	"github.com/itchyny/timefmt-go"
	"gopkg.in/ini.v1"
)

type InfoData struct {
	Text              string `xml:",chardata"`
	TargetGameFolder  string `xml:"target_game_folder"`
	TargetImageFolder string `xml:"target_image_folder"`
	TargetVideoFolder string `xml:"target_video_folder"`
}
type GameData struct {
	Name       string
	FilePath   string
	VideoPath  string
	ImagePath  string
	CreateTime string
}

// export GOOS=windows
// export GOARCH=amd64
// go build -o makegamexml.exe && zip makegamecml.zip makegamexml.exe setting.ini template.xml

// GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o makegamexml.exe && zip makegamexml.zip makegamexml.exe setting.ini template.xml
func main() {

	cfg, err := ini.Load("./setting.ini")

	if err != nil {
		panic(err)
	}

	game_folder := cfg.Section("info").Key("game_folder").String()
	image_folder := cfg.Section("info").Key("image_folder").String()
	video_folder := cfg.Section("info").Key("video_folder").String()
	createtime := cfg.Section("info").Key("CreateTime").String()
	exs := strings.Split(cfg.Section("info").Key("ex").String(), ",")

	println(game_folder)
	println(image_folder)
	println(video_folder)

	// dir
	gameDir, err := os.ReadDir(game_folder)
	if err != nil {
		panic("게임 폴더가 업습니다.")
	}

	// imageDir, _ := os.ReadDir(image_folder)
	// videoDir, _ := os.ReadDir(video_folder)
	// if imageDir != nil {

	// }
	// if videoDir != nil {

	// }
	gameDataList := []GameData{}
	for _, gamefile := range gameDir {

		if gamefile.IsDir() {
			continue
		}
		if !checkEx(exs, filepath.Ext(gamefile.Name())) {
			continue
		}
		gameData := GameData{}
		fullpath, _ := filepath.Abs(fmt.Sprint(game_folder, "/", gamefile.Name()))
		gameData.Name = fileNameWithoutExtTrimSuffix(gamefile.Name())
		gameData.FilePath = fmt.Sprint(gamefile.Name())
		at, atErr := atime.Stat(fullpath)
		if atErr == nil {
			str := timefmt.Format(at, createtime)
			gameData.CreateTime = str
		}

		// fmt.Println(info.ModTime())

		imgExList := []string{"png", "PNG", "jpg", "PNG", "bmp", "BMP"}
		for _, ex := range imgExList {
			var fpath = fmt.Sprint(image_folder, "/", gameData.Name, ".", ex)
			if Exists(fpath) {
				gameData.ImagePath = fmt.Sprint(gameData.Name, ".", ex)
				break
			}
		}

		VideoExList := []string{"mp4", "MP4", "MOV", "mov", "AVI", "wmv", "WMV", "avi"}
		for _, ex := range VideoExList {
			var fpath = fmt.Sprint(video_folder, "/", gameData.Name, ".", ex)

			if Exists(fpath) {
				gameData.VideoPath = fmt.Sprint(gameData.Name, ".", ex)

				break
			}
		}

		gameDataList = append(gameDataList, gameData)

	}

	templateStr, errTemplate := os.ReadFile("./template.xml")
	if errTemplate != nil {
		panic("template.xml 에 읽기에 실패 하였습니다.")
	}

	t := template.New("")
	t1, err := t.Parse(string(templateStr))

	if err != nil {
		panic("template.xml 에 오류가 생겼습니다.")
	}

	ofile, err := os.Create("./gamelist.xml")
	fmt.Print(err)

	t1.Execute(ofile, gameDataList)

}

func checkEx(esList []string, fileEx string) bool {
	fileEx = strings.ReplaceAll(fileEx, ".", "")
	for _, ex := range esList {
		if strings.Compare(strings.ToLower(ex), strings.ToLower(fileEx)) == 0 {
			return true
		}
	}
	return false
}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
func Exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}
