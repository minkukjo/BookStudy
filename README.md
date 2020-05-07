# Go + Vue.js 로 만들어보는 게시판 프로젝트

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