## Perintah Go Test

```bash
go test ./... -v -coverprofile=cover.out && go tool cover -html=cover.out
```

```bash
go test ./... --coverprofile cover.out
go tool cover -func cover.out
```

## Generate Mock data via mockery
```bash
mockery --dir=features/users --name=UserDataInterface --filename=UserData.go --structname=UserData
```