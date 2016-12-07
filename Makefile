CPD_PLATFORM=rpi3
CPCD_PLATFORM=linux_amd64

ifeq ($(CPD_PLATFORM),rpi3)
	CPD_ENV= GOOS=linux GOARCH=arm GOARM=7
endif
ifeq ($(CPD_PLATFORM),linux_amd64)
	CPD_ENV= GOOS=linux GOARCH=amd64
endif

ifeq ($(CPCD_PLATFORM),linux_amd64)
	CPCD_ENV= GOOS=linux GOARCH=amd64
endif

CPCTL_ENV=

clean:
	rm -f build/*

deps:
	cd cmd/ && go get ...

cpd:
	cd cmd/cpd && $(CPD_ENV) go build -o ../../build/cpd

cpcd:
	cd cmd/cpcd && $(CPCD_ENV) go build -o ../../build/cpcd

cpctl:
	cd cmd/cpctl && $(CPCTL_ENV) go build -o ../../build/cpctl

build: clean cpd cpcd cpctl

test_ctl_cd: clean cpcd cpctl
	killall -q cpcd || true
	build/cpcd &
	sleep 1
	build/cpctl -c 127.0.0.1:3334 nodes
	killall -q cpcd || true
