run: 
	go run main.go

test:
	go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof ./lists

inspect_cpu:
	go tool pprof cpu.prof

inspect_memory:
	go tool pprof mem.prof	

inspect_cpu_web: 
	go tool pprof -http=:8080 cpu.prof

inspect_mem_web: 
	go tool pprof -http=:8080 mem.prof	