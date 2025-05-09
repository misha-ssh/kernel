# build and start ssh server with default port
up-ssh:
	docker build ./build/ssh -t ssh-host
	docker run -d --name ssh-default -p 22:22 ssh-host

# stop and rm ssh-default container
down-ssh:
	docker stop ssh-default
	docker rm ssh-default

# build and start ssh server with 2222 port
up-ssh-port:
	docker build ./build/ssh -t ssh-host
	docker run -d --name ssh-port -p 2222:22 ssh-host

# stop and rm ssh-port container
down-ssh-port:
	docker stop ssh-port
	docker rm ssh-port