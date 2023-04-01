
build_windows:
	@echo "Generating windows executable file"
	env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" go build -o bin/cmdhelp.exe
	@echo "Generated the cmdhelp.exe file"

build_linux:
	@echo "Generating linux binary"
	go build -o bin/cmdhelp
	@echo "Generated cmdhelp binary file"

build_all:
	env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" go build -o bin/cmdhelp.exe
	go build -o bin/cmdhelp

mod-update:
	go get -u
	go mod tidy

mod-tidy:
	go mod tidy

clean:
	rm bin/cmdhelp bin/cmdhelp.exe 
