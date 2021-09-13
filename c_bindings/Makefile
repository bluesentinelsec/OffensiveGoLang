target: static-library

static-library:
	go build -o libDiscovery.a -buildmode=c-archive discovery.go

shared-library:
	go build -o libDiscovery.so -buildmode=c-shared discovery.go

clean:
	rm libDiscovery.a
	rm libDiscovery.h
	rm libDiscovery.so

example:
	go build -o libDiscovery.a -buildmode=c-archive discovery.go
	cp libDiscovery.a ../examples/use_c_bindings/
	cp libDiscovery.h ../examples/use_c_bindings/
	go build -o libDiscovery.so -buildmode=c-shared discovery.go
	cp libDiscovery.so ../examples/use_with_python/
	cp libDiscovery.h ../examples/use_with_python/
