# GO VDF
*Running Verifiable Delay Function in GO*
Execution example using [ChiaVDF](https://github.com/Chia-Network/chiavdf.git).

## Run using Docker
Raise the docker container using `docker compose`. It produces a one-shot execution.
```
docker compose up --build
```


## Run Locally (Ubuntu)
1. Load ChiaVDF submodule bu running:
```
git submodule update --init
```

2. Install in your device the required compiling tools.
```
sudo apt-get install cmake libgmp-dev libboost-system-dev build-essential -y
```

3. Create directory `chiavdf/build`, and then move into it.
```
mkdir chiavdf/build && cd chiavdf/build
```

4. Run CMake to configure the project, establishing flags to compile `chiavdfc` and ommit python library. Pass as parameter the directory in which `CMakeLists.txt` is located.
```
cmake -DBUILD_CHIAVDFC=ON -DBUILD_PYTHON=OFF ../src
```

5. Build Chia VDF Library.
```
make
```
By following the previous commands, we obtain a shared library: `libchiavdfc.so`, located inside `chiavdf/build/shared`.

6. Go back to root project's folder.
```
cd ../..
```

7. Export path of C wrappers, ChiaVDF shared library and libgmp.
```sh
export CGO_CFLAGS="-I/path/to/chiavdf/src/c_bindings"
export CGO_LDFLAGS="-L/path/to/chiavdf/build/lib/shared -L/usr/lib/x86_64-linux-gnu -lgmp"
export LD_LIBRARY_PATH="/path/to/chiavdf/build/lib/shared"
```

8. Run main program:
```sh
go run main.go
```
