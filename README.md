# 휴아봇 백엔드(Golang)

## 개발 환경
* Windows 10 19043
* Go 1.15.11
* fiber framework

## 실행
* Docker 사용시
```console
$ sudo docker build --tag <이미지 이름> .
$ sudo docker run --name <컨테이너 이름> -d -p 8080:8080 <직전에 정한 이미지 이름>
```

* 직접 실행시
```console
$ go run .\main.go
```

## URL
### /kakao : 카카오 i 챗봇을 위한 url prefix
* /shuttle : 사용자의 정류장 입력을 받아 해당 정류장의 도착 정보를 반환
* /shuttle/stop : 사용자의 정류장 입력을 받아 해당 정류장의 첫/막차 및 로드뷰 반환
* /bus : 10-1, 707-1, 3102 노선의 GBIS 도착 정보 및 종점 시간표 반환
* /subway : 한대앞역(4호선 및 수인분당선)의 실시간 도착 정보 및 시간표 반환
* /food : ERICA 교내 식당에 대한 메뉴를 Firebase 를 통해 반환
* /library : ERICA 학술정보관 내 열람실 좌석 정보를 Firebase 를 통해 반환

### /app : 휴아봇 앱을 위한 url prefix
* /shuttle(GET) : 전첸 정류장의 도착 정보를 반환
* /shuttle(POST) : 사용자의 정류장 입력을 받아 해당 정류장의 도착 정보를 반환
* /shuttle/stop : 사용자의 정류장 입력을 받아 해당 정류장의 첫/막차 및 로드뷰 반환
* /bus : 10-1, 707-1, 3102 노선의 GBIS 도착 정보 및 종점 시간표 반환
* /subway : 한대앞역(4호선 및 수인분당선)의 실시간 도착 정보 및 시간표 반환
* /food : ERICA 교내 식당에 대한 메뉴를 Firebase 를 통해 반환
* /library : ERICA 학술정보관 내 열람실 좌석 정보를 Firebase 를 통해 반환

### /common : 공통 기능을 위한 url prefix
* /food : ERICA 교내 식당에 대한 메뉴를 Firebase 에 저장
* /library : ERICA 학술정보관 내 열람실 좌석 정보를 Firebase 에 저장