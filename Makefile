PREFIX := /usr
DESTDIR := /
BINARY_NAME := apx

GO := go

all: clean build

build: ${BINARY_NAME}

${BINARY_NAME}:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -a -tags netgo -ldflags '-w -extldflags "-static"' -o $@ ./cmd

install: build
	install -Dm755 ${BINARY_NAME} ${DESTDIR}${PREFIX}/bin/${BINARY_NAME}
	mkdir -p ${DESTDIR}/etc/apx
	sed -i 's|/usr/share/apx/distrobox|${PREFIX}/share/apx/distrobox|g' config/apx.json
	install -Dm644 config/apx.json ${DESTDIR}/etc/apx/apx.json
	mkdir -p ${DESTDIR}${PREFIX}/share/apx/distrobox
	sh distrobox/install --prefix ${DESTDIR}${PREFIX}/share/apx/distrobox
	mv ${DESTDIR}${PREFIX}/share/apx/distrobox/bin/distrobox* ${DESTDIR}${PREFIX}/share/apx/distrobox/.

install-manpages:
	mkdir -p ${DESTDIR}${PREFIX}/share/man/man1
	cp -r man/* ${DESTDIR}${PREFIX}/share/man/.
	chmod 644 ${DESTDIR}${PREFIX}/share/man/man1/apx*

uninstall: uninstall-manpages
	rm ${DESTDIR}${PREFIX}/bin/${BINARY_NAME}
	rm -rf ${DESTDIR}/etc/apx
	rm -rf ${DESTDIR}${PREFIX}/share/apx

uninstall-manpages:
	rm -rf ${DESTDIR}${PREFIX}/share/man/man1/apx*

clean:
	rm -f ${BINARY_NAME}
	${GO} clean

# dist: build multi-OS binaries into dist/<OS> folders
.PHONY: dist

dist:
	mkdir -p dist/Mac dist/Windows-10-11 dist/Linux
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GO} build -o dist/Mac/${BINARY_NAME}-darwin-amd64 ./cmd
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 ${GO} build -o dist/Mac/${BINARY_NAME}-darwin-arm64 ./cmd
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o dist/Linux/${BINARY_NAME}-linux-amd64 ./cmd
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 ${GO} build -o dist/Linux/${BINARY_NAME}-linux-arm64 ./cmd
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GO} build -o dist/Windows-10-11/${BINARY_NAME}-windows-amd64.exe ./cmd
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 ${GO} build -o dist/Windows-10-11/${BINARY_NAME}-windows-arm64.exe ./cmd
