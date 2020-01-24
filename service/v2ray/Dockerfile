FROM ubuntu:18.04

MAINTAINER KallyDev kallydev@gmail.com

RUN apt update && \
    apt install wget python3.8 python3-pip unzip -y && \
    pip3 install argparse

RUN wget https://github.com/v2ray/v2ray-core/releases/download/v4.22.1/v2ray-linux-64.zip && \
    mkdir v2ray && \
    unzip v2ray-linux-64.zip -d v2ray && \
    rm v2ray-linux-64.zip

ADD run.py ./v2ray

ADD bbr.sh ./v2ray

ADD config.json ./v2ray

WORKDIR v2ray

CMD bash bbr.sh

ENTRYPOINT [ "python3.8", "run.py" ]
