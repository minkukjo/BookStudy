# Go + Vue.js 로 만들어보는 게시판 프로젝트

## 메인 화면

<img width="1429" alt="스크린샷 2020-05-10 오전 5 08 00" src="https://user-images.githubusercontent.com/43809168/81484159-839f2100-927e-11ea-9f7a-6b5567876193.png">

## Main

![main](https://user-images.githubusercontent.com/43809168/81484162-8863d500-927e-11ea-87f9-0a7d18974333.gif)

## Task List

-[ ] 글 종류 수정하고 Nav Bar는 이전 상태 그대로 남아있는 이슈있음
-[ ] 사용자 정보 보여주기 더 깔끔하게 수정
-[ ] Gitflow 관리 (rebase 사용해서)
-[ ] 백엔드 중요 로직 테스트 코드 짜기
-[ ] 배포 하기 (GCP? AWS?)
-[ ] 젠킨스 CI/CD 구축해보기 ( Terraform? Ansible? )
-[ ] 도커로 패키징해보기 (쿠버네티스 까진 필요 없을듯 추후에 서버가 마이크로하게 확장된다면..? 설마 ㅋㅋ)

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