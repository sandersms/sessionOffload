FROM rockylinux:9
USER root

RUN mkdir -p /home/grpc/local

WORKDIR /home/grpc/local

ENV GRPC_INSTALL /home/grpc/local

RUN echo $GRPC_INSTALL && \
    yum -y install wget && \
    yum -y install which && \
    yum -y groupinstall "Development tools" && \
    yum clean all

RUN wget --no-check-certificate -q -O cmake-linux.sh https://github.com/Kitware/CMake/releases/download/v3.23.4/cmake-3.23.4-linux-x86_64.sh && \
    sh ./cmake-linux.sh -- --skip-license --prefix=/home/grpc/local && \
    rm cmake-linux.sh

ENV PATH="$GRPC_INSTALL/bin:${PATH}"

RUN echo $PATH && \
    git config --global http.sslVerify false && \
    git clone --recurse-submodules -b v1.49.1 https://github.com/grpc/grpc grpc && \
    mkdir -p /home/grpc/local/grpc/cmake/build

WORKDIR /home/grpc/local/grpc/cmake/build

RUN echo $PATH

#CMD ["sh", "-c", "tail -f /dev/null"]

RUN pwd && \
    cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=/home/grpc/local -DCMAKE_POSITION_INDEPENDENT_CODE=TRUE -DCMAKE_CXX_STANDARD=17 -DABSL_PROPAGATE_CXX_STD=TRUE ../.. && \
    make -j 4 && \
    make install && \
    mkdir -p /home/grpc/tmp

WORKDIR /home/grpc/tmp

RUN wget -nv http://hyperrealm.github.io/libconfig/dist/libconfig-1.7.3.tar.gz && \
    tar xvzf libconfig-1.7.3.tar.gz

WORKDIR /home/grpc/tmp/libconfig-1.7.3

RUN ./configure && \
    make && \
    make install

WORKDIR /home/grpc/
#
CMD ["sh", "-c", "tail -f /dev/null"]
