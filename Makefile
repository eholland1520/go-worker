BINARY_NAME=go-worker
 
build_test:
	#Download TwistCLI
	curl --progress-bar -L -k --header "authorization: Bearer [prisma-bearer-token]" [prisma-console-url]/api/v1/util/osx/twistcli > twistcli; chmod a+x twistcli;
	chmod a+x twistcli
	docker build -t prismaclouddev/go-worker:sa .
	./twistcli images scan --details --address [prisma-console-url] --u [user-access-token] -p [access-token-password] prismaclouddev/go-worker:sa
build:
	docker build -t prismaclouddev/go-worker:sa .
run:
	docker-compose up -d
build_test_run:
	docker build -t prismaclouddev/go-worker:sa .
	./twistcli images scan --details --address [prisma-console-url] --u [user-access-token] -p [access-token-password] prismaclouddev/go-worker:sa
	docker-compose up -d
test:
	./twistcli images scan --details --address [prisma-console-url] --u [user-access-token] -p [access-token-password] prismaclouddev/go-worker:sa
clean:
	docker-compose down
install_twistcli:
	#Download TwistCLI
	curl --progress-bar -L -k --header "authorization: Bearer [prisma-bearer-token]" [prisma-console-url]/api/v1/util/osx/twistcli > twistcli; chmod a+x twistcli;
	chmod a+x twistcli;