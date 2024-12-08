build: 
	g++ -shared -o libbpewrapper.so -fPIC src/*.cpp wrapper/*.cpp

install:
	mkdir -p /usr/local/include/youtokentogo
	cp libbpewrapper.so /usr/local/lib/
	cp wrapper/wrapper.h /usr/local/include/youtokentogo/
	ldconfig
