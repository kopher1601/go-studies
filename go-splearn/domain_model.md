## 도메인 모델 만들기
1. 듣고 배우기
	- Domain Expert 로 부터
2. `중요한 것`들 찾기
	- Concept Identification (개념 식별)
3. `연결고리` 찾기
	- 관계 정의
4. `것` 들을 설명하기 (속성 및 기본 행위 명시)
5. 그려보기 (시각화)
6. 이야기 하고 다듬기 (반복)

## 도메인 모델

### 회원(`Member`)
_Entity_

#### 속성
- `email`: 이메일 - ID
- `nickname`: 닉네임
- `passwordHash`: 비밀번호
- `status`: 회원 상태(`MemberStatus`)

#### 행위
- `private constructor()`: 회원 생성: email, nickname, passwordHash
- `static create()`: 회원 생성: email, nickname, password, passwordEncoder
- `activate()`: 가입을 완료 시킨다
- `deactivate()`: 탈퇴 시킨다
- `verifyPassword()`: 비밀번호를 검증한다

#### 규칙
- 회원 생성후 상태는 가입 대기
- 일정 조건을 만족하면 가입 완료가 된다
- 가입 대기 상태에서만 가입 완료가 될 수 있다
- 가입 완료 상태에서는 탈퇴할 수 있다
- 회원의 비밀번호는 해시를 만들어서 저장한다

### 회원 상태(`MemberStatus`)
_Enum_
#### 상수
- `PENDING`: 가입 대기
- `ACTIVE`: 가입 완료
- `DEACTIVATED`: 

### 비밀번호 인코더(`PasswordEncoder`)
_Domain Service_

#### 행위
- `encode()`: 비밀번호 암호화하기
- `matches()`: 비밀번호가 일치하는지 확인


### 강사

### 강의

### 수업

### 섹션

### 수강

### 진도

