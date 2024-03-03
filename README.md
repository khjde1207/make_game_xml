# make game xml 

일부 휴대용 게임기는 게임명, 게임경로, 비디오 파일 경로, 이미지경로를 하나하나 입력 해 줘야 합니다. 
이 프로젝트를 이를 자동으로 만들어 주기 위해 제작 하였습니다. 

`template.xml` 파일은 결과물로 나오게 될 xml 템플릿 입니다. 

```
<contents>
    <gameList>
    {{range $val := .}} 
        <!--▼ 이 부분만 수정이 가능합니다 ▼-->
        <game>
            <path>./{{.FilePath}}</path>
            <name>{{.Name}}</name>
            <ImagePath>./downloaded_images/{{.ImagePath}}</ImagePath>
            <VideoPath>./downloaded_videos/{{.VideoPath}}</VideoPath> 
            <rating>0</rating>
            <releasedate>{{.CreateTime}}</releasedate>
            <playcount>0</playcount>
            <lastplayed>{{.CreateTime}}</lastplayed>
        <game>
        <!--▲ 이 부분만 수정이 가능합니다 ▲-->
    {{end}}
    </gameList> 
</contents>
```

Name : 파일명

FilePath : 파일 경로 

ImagePath : 이미지 파일 경로 

VideoPath : 비디오 파일 경로 

CreateTime : 파일 생성 일자 

---

`setting.ini` 파일은 게임, 비디오파일, 이미지 폴더 경로를 설정 할 수 있습니다. 

```
[info]
game_folder = ./psp
image_folder = ./psp/downloaded_images
video_folder = ./psp/downloaded_videos
CreateTime= %Y%m%dT%H%M%S
```

game_folder : 게임폴더 경로

image_folder : 이미지폴더 경로

video_folder : 비디오폴더 경로

CreateTime : 파일 생성일 포멧 형식 

시간 포멧은 링크를 참고 해주세요 : https://github.com/itchyny/timefmt-go