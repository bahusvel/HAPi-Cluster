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
CONTROLLER_IP=192.168.7.28

clean:
	rm -f build/*

deps:
	cd cmd/ && go get ...

ssh_id:
	if [ ! -f ~/.ssh/id_rsa ]; then ssh-keygen; fi
	ssh-copy-id root@$(CONTROLLER_IP)

cpd:
	cd cmd/cpd && $(CPD_ENV) go build -o ../../build/cpd

cpcd:
	cd cmd/cpcd && $(CPCD_ENV) go build -o ../../build/cpcd

cpctl:
	cd cmd/cpctl && $(CPCTL_ENV) go build -o ../../build/cpctl

build: clean cpd cpcd cpctl

test_cluster:
	make build CPD_PLATFORM=rpi3 CPCD_PLATFORM=linux_amd64
	ssh root@$(CONTROLLER_IP) "killall cpcd" || true
	scp build/cpcd root@$(CONTROLLER_IP):/usr/bin/
	scp build/cpd root@$(CONTROLLER_IP):/nfs/share/bin/
	ssh root@$(CONTROLLER_IP) "bash distribute_cpd.sh" &
	ssh root@$(CONTROLLER_IP) "/usr/bin/cpcd -i"

test_ctl_cd: build
	make build CPD_PLATFORM=current CPCD_PLATFORM=current
	killall cpcd || true
	build/cpcd &
	sleep 1
	build/cpctl -c 127.0.0.1:3334 nodes
	killall cpcd || true

test_controller:
	make build CPD_PLATFORM=current CPCD_PLATFORM=current
	killall cpcd || true
	killall cpd || true
	build/cpcd &
	sleep 1
	build/cpd -i 127.0.0.1 -c 127.0.0.1:3334  &
	sleep 5
	killall cpcd || true
	killall cpd || true
