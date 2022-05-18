module github.com/MoraGames/StreamingScheduler/auth/internal/jwt

go 1.17

replace github.com/MoraGames/StreamingScheduler/auth/internal/utils => ../utils

require (
	github.com/MoraGames/StreamingScheduler/auth/internal/utils v0.0.0-20220518202815-3504a9a9dd15
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
)
