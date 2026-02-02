FROM ubuntu:22.04 as kassebuilder

# update everything
RUN apt update -y && apt upgrade -y && apt install git sudo wget gcc libc-dev -y 
# clone some code
WORKDIR /online_kasse
COPY . .
# RUN git clone https://github.com/OpenRbt/online_kasse


RUN wget https://golang.org/dl/go1.24.11.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.24.11.linux-amd64.tar.gz

RUN chmod 777 /usr/local/go/bin/go

RUN export PATH=$PATH:/usr/local/go/bin && ls && cd cmd && cd web && go build
RUN wget -O driver.zip "https://f1.atoldriver.ru/10/10.9.3.1/10.9.3.1.zip"
RUN apt install -y  p7zip-full
RUN 7z x driver.zip


FROM ubuntu:22.04
RUN apt update -y && apt upgrade -y && apt install -y p7zip-full gcc libc-dev wget libusb-1.0-0 && apt clean

COPY --from=kassebuilder /online_kasse/10.9.3.1/installer/deb/libfptr10_10.9.3.1_amd64.deb /10.9.3.1/installer/deb/libfptr10_10.9.3.1_amd64.deb

RUN chmod 777 "/10.9.3.1/installer/deb/libfptr10_10.9.3.1_amd64.deb" && \
    apt-get install -y "./10.9.3.1/installer/deb/libfptr10_10.9.3.1_amd64.deb"
COPY --from=kassebuilder /online_kasse/cmd/web/web /kasse
COPY --from=kassebuilder /online_kasse/cmd/web/key.pem /
COPY --from=kassebuilder /online_kasse/cmd/web/cert.pem /

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# let's build everything
CMD [ "/entrypoint.sh" ]
