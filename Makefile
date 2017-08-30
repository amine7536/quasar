all: quasar

GO ?= go
GOTEST = go test -v -bench\=.

test:
	$(GOTEST) github.com/amine7536/quasar/quasar
	$(GOTEST) github.com/amine7536/quasar/conf
	$(GOTEST) github.com/amine7536/quasar/output/stdout
	$(GOTEST) github.com/amine7536/quasar/output/file
	$(GOTEST) github.com/amine7536/quasar/output/logstash
	$(GOTEST) github.com/amine7536/quasar/event
	$(GOTEST) github.com/amine7536/quasar/cmd
	
quasar: test
	mkdir -p build
	$(GO) env
	$(GO) build -ldflags="-s -w" $(EXTRA_BUILD_FLAGS) -o build/quasar

tmp/quasar.tar.gz: quasar
	mkdir -p tmp/
	rm -rf tmp/quasar
	mkdir -p tmp/quasar/
	cp build/quasar tmp/quasar/quasar
	cp deploy/quasar.json tmp/quasar/quasar.json
	cp deploy/quasar.sysconfig tmp/quasar/quasar.sysconfig
	cp deploy/quasar.service tmp/quasar/quasar.service
	cd tmp && tar czf quasar.tar.gz quasar/

rpm: tmp/quasar.tar.gz
	chmod +x deploy/buildrpm.sh
	cp deploy/buildrpm.sh tmp/buildrpm.sh
	cd tmp && ./buildrpm.sh ../deploy/quasar.spec.centos `../build/quasar version`
	cp tmp/rpm/RPMS/x86_64/quasar-*.rpm build/

clean-rpm:
	rm -fr tmp

clean:
	rm -f build/quasar

clean-all: clean clean-rpm

image:
	CGO_ENABLED=0 GOOS=linux $(MAKE) quasar
	docker build -t quasar -f quasar.Dockerfile .