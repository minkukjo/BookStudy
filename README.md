# Go + Vue.js 로 만들어보는 게시판 프로젝트

## 데모

[데모](http://34.67.130.46:9090/)

배포환경 GCP + Docker

## 메인 화면

<img width="1429" alt="스크린샷 2020-05-10 오전 5 08 00" src="https://user-images.githubusercontent.com/43809168/81484159-839f2100-927e-11ea-9f7a-6b5567876193.png">

## Main

![main](https://user-images.githubusercontent.com/43809168/81484162-8863d500-927e-11ea-87f9-0a7d18974333.gif)

## Create

![Create](https://user-images.githubusercontent.com/43809168/81484232-e690b800-927e-11ea-86cc-bd9996e15529.gif)

## Read

![Read](https://user-images.githubusercontent.com/43809168/81484234-e85a7b80-927e-11ea-896e-433d34c099bb.gif)

## Update

![Update](https://user-images.githubusercontent.com/43809168/81484281-4edf9980-927f-11ea-91cd-95227c2ddb60.gif)

## Delete

![삭제](https://user-images.githubusercontent.com/43809168/81484238-edb7c600-927e-11ea-8b69-1f033cb4993c.gif)

## 사용자 정보 조회 & 로그아웃

![last](https://user-images.githubusercontent.com/43809168/81484241-f3151080-927e-11ea-912c-a2b1a22278ff.gif)

## Task List

- [ ] 글 종류 수정하고 Nav Bar는 이전 상태 그대로 남아있는 이슈있음
- [ ] 사용자 정보 보여주기 더 깔끔하게 수정
- [ ] Gitflow 관리 (rebase 사용해서)
- [ ] 백엔드 중요 로직 테스트 코드 짜기
- [ ] 배포 하기 (GCP? AWS?)
- [ ] 도커 컴포즈 이용
- [ ] GCP CI/CD 구축해보기

## 기술 스택

- Go (Back)

- Go-chi (웹 프레임워크)

- Gorm (ORM)

- Vue.js (Front) + Vuex (상태 관리) + Vuex Persistedstate (Store 영구 저장)

- Redis (DB)

- MySQL (DB)

## 실행 방법

**서버**

```bash
go run ./main.go
```

**프론트**

```bash
vue-cli-service build --watch
```

변경 사항 저장할 때 마다 자동으로 빌드해줘서 개꿀
