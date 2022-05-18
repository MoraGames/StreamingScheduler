module mail

replace github.com/MoraGames/StreamingScheduler/auth/internal/utils => ../utils

go 1.17

require github.com/MoraGames/StreamingScheduler/auth/internal/utils v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
)
