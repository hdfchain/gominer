CC ?= gcc -fPIC
CXX ?= g++ -fPIC
NVCC ?= nvcc -Xcompiler -fPIC
AR ?= ar
# -o is gnu only so this needs to be smarter; it does work because on darwin it
#  fails which is also not windows.
ARCH:=$(shell uname -o)

.DEFAULT_GOAL := build

ifeq ($(ARCH),Msys)
nvidia:
endif

# Windows needs additional setup and since cgo does not support spaces in
# in include and library paths we copy it to the correct location.
#
# Windows build assumes that CUDA V7.0 is installed in its default location.
#
# Windows gominer requires nvml.dll and hdfchain.dll to reside in the same
# directory as gominer.exe.
ifeq ($(ARCH),Msys)
obj: nvidia
	mkdir nvidia
	cp -r /c/Program\ Files/NVIDIA\ GPU\ Computing\ Toolkit/* nvidia
	cp -r /c/Program\ Files/NVIDIA\ Corporation/NVSMI nvidia
else
obj:
endif
	mkdir obj

ifeq ($(ARCH),Msys)
obj/hdfchain.dll: obj sph/blake.c hdfchain.cu
	$(NVCC) --shared --optimize=3 --compiler-options=-GS-,-MD -I. -Isph hdfchain.cu sph/blake.c -o obj/hdfchain.dll
else
obj/hdfchain.a: obj sph/blake.c hdfchain.cu
	$(NVCC) --lib --optimize=3 -I. hdfchain.cu sph/blake.c -o obj/hdfchain.a
endif

ifeq ($(ARCH),Msys)
build: obj/hdfchain.dll
else
build: obj/hdfchain.a
endif
	go build -tags 'cuda'

ifeq ($(ARCH),Msys)
install: obj/hdfchain.dll
else
install: obj/hdfchain.a
endif
	go install -tags 'cuda'

clean:
	rm -rf obj
	go clean
ifeq ($(ARCH),Msys)
	rm -rf nvidia
endif
