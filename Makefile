test:
	go test
bench:
	go test -bench=. -benchtime=5s
devtest:
	go test -cfv=dev.conf
devbench:
	go test -bench=. -benchtime=5s -cfg=dev.conf
