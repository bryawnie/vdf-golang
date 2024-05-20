# GO VDF
*Running Verifiable Delay Function in GO*

## Compile Chia VDF
1. Clone [ChiaVDF official repostory](https://github.com/Chia-Network/chiavdf.git). To do so, we can run:
```
git submodule update --init
```

2. Install compiling tools.
```
sudo apt-get install cmake libgmp-dev libboost-python-dev $PYTHON_DEV_DEPENDENCY libboost-system-dev build-essential -y
```

3. Edit `chiavdf/src/CMakeLists.txt`, by turning on the `chiavdfc` option, and turning off the `python` option.
```
option(BUILD_CHIAVDFC "Build the chiavdfc shared library" ON)
option(BUILD_PYTHON "Build the python bindings for chiavdf" OFF)
```

4. Create directory `chiavdf/build`. Then move into it.
```
mkdir chiavdf/build
cd chiavdf/build
```

5. Run CMake to configure the project. Pass as parameter the directory in which `CMakeLists.txt` is located.
```
cmake ../src
```

6. Build the project.
```
make
```

By following these commands, we obtain a compiled C VDF library: `libchiavdfc.a`, located inside `chiavdf/build/static`.

## RUN GO FILE
1. Export absolute path of C Wrappers for ChiaVDF.
```
export CGO_CFLAGS="-I/path/to/chiavdf/src/c_bindings"
```

2. Export absolute path ob libchiavdfc and libgmp.
```
export CGO_LDFLAGS="-L/path/to/chiavdf/build/lib/static -L/usr/lib/x86_64-linux-gnu -lgmp"
```

3. (⚠️ Not working) Run main program.
```
go run main.go
```