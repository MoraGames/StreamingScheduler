module github.com/MoraGames/StreamingScheduler/auth

replace github.com/MoraGames/StreamingScheduler/auth/internal/jwt => ./internal/jwt

replace github.com/MoraGames/StreamingScheduler/auth/internal/mail => ./internal/mail

replace github.com/MoraGames/StreamingScheduler/auth/internal/utils => ./internal/utils

replace github.com/MoraGames/StreamingScheduler/auth/internal/password => ./internal/password

go 1.17

require (
	github.com/MoraGames/StreamingScheduler/auth/internal/jwt v0.0.0-00010101000000-000000000000
	github.com/MoraGames/StreamingScheduler/auth/internal/mail v0.0.0-00010101000000-000000000000
	github.com/MoraGames/StreamingScheduler/auth/internal/utils v0.0.0-20220518202815-3504a9a9dd15
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)
