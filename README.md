# TAXII 서버 구축

MITRE ATT&CK 데이터를 기반으로 동작하는 TAXII 서버 구축 엔진
Go 언어(Golang)로 직접 구현, MITRE의 `cti` 저장소를 주기적으로 동기화

## 기능

- MITRE ATT&CK 최신 데이터 GitHub에서 자동 클론 / 업데이트
- 10분 주기 자동 pull & 커밋 비교 (테스트용으로 1분)
- 엔진 형태로 항상 실행 유지 (스케줄러 역할)
- 향후 확장 예정:
  - STIX JSON 파싱
  - API 제공 (TAXII Discovery, Collections 등)
  - Elasticsearch / ClickHouse 연동

## 실행 방법

```bash
go run main.go
