FROM ubuntu:22.04 as chiavdf

RUN apt-get update && \
    apt-get install -y cmake libgmp-dev libboost-system-dev build-essential git

RUN git clone https://github.com/Chia-Network/chiavdf
RUN mkdir /chiavdf/build

WORKDIR /chiavdf/build

RUN git reset --hard 2844974ff81274060778a56dfefd2515bc567b90

RUN cmake -DBUILD_CHIAVDFC=ON -DBUILD_PYTHON=OFF ../src
RUN make


FROM golang:1.22-bookworm as vdf-golang

RUN apt-get update && \
    apt-get install -y libgmp-dev build-essential libboost-system-dev

COPY --from=chiavdf /chiavdf/build/lib /chiavdf/build/lib
COPY --from=chiavdf /chiavdf/src/c_bindings /chiavdf/src/c_bindings

ENV CGO_CFLAGS="-I/chiavdf/src/c_bindings"
ENV CGO_LDFLAGS="-L/chiavdf/build/lib/shared -L/usr/lib/x86_64-linux-gnu -lgmp"
ENV LD_LIBRARY_PATH="/chiavdf/build/lib/shared"

COPY . /app

WORKDIR /app

CMD ["go", "run", "main.go"]