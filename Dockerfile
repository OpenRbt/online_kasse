FROM ubuntu:22.04 AS kassebuilder

# update everything
RUN apt update -y && apt upgrade -y && apt install git sudo wget gcc libc-dev -y 
# clone some code
WORKDIR /online_kasse
COPY . .
# RUN git clone https://github.com/OpenRbt/online_kasse


RUN GOARCH=$(case "$(dpkg --print-architecture)" in \
        amd64) echo amd64;; \
        arm64) echo arm64;; \
        armhf) echo armv6l;; \
        *) echo amd64;; \
    esac) && \
    wget "https://golang.org/dl/go1.24.11.linux-${GOARCH}.tar.gz" && \
    tar -C /usr/local -xzf "go1.24.11.linux-${GOARCH}.tar.gz"

RUN chmod 777 /usr/local/go/bin/go

RUN export PATH=$PATH:/usr/local/go/bin && ls && cd cmd && cd web && go build
RUN wget -O driver.zip "https://f1.atoldriver.ru/10/10.10.7.0/10.10.7.0.zip"
RUN apt install -y  p7zip-full
RUN 7z x driver.zip
# normalize the arch-specific driver deb name so the next stage can COPY a fixed path
RUN cp "/online_kasse/10.10.7.0/installer/deb/libfptr10_10.10.7.0_$(dpkg --print-architecture).deb" /libfptr10.deb


FROM ubuntu:22.04
RUN apt update -y && apt upgrade -y && apt install -y p7zip-full gcc libc-dev wget libusb-1.0-0 && apt clean

COPY --from=kassebuilder /libfptr10.deb /tmp/libfptr10.deb

RUN apt-get install -y /tmp/libfptr10.deb
COPY --from=kassebuilder /online_kasse/cmd/web/web /kasse
COPY --from=kassebuilder /online_kasse/cmd/web/key.pem /
COPY --from=kassebuilder /online_kasse/cmd/web/cert.pem /

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# let's build everything
CMD [ "/entrypoint.sh" ]
