export GO111MODULE=on

aikor-nginx: *.go builder/*.go command/*.go configure/*.go module3rd/*.go openresty/*.go util/*.go
	go build -ldflags "-X main.NginxBuildVersion=`git rev-list HEAD -n1`" -o $@

build-example: aikor-nginx
	./aikor-nginx -c config/configure.example -m config/modules.cfg.example -d work -clear

check:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf aikor-nginx
