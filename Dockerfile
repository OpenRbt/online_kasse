FROM ubuntu:22.04 as kassebuilder

# update everything
RUN apt update -y && apt upgrade -y && apt install git sudo wget gcc libc-dev -y 
# clone some code
RUN git clone https://github.com/OpenRbt/online_kasse


RUN wget https://golang.org/dl/go1.17.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz

RUN chmod 777 /usr/local/go/bin/go

RUN export PATH=$PATH:/usr/local/go/bin && cd /online_kasse/cmd/web && go build
RUN wget -O driver.zip "https://f1.atoldriver.ru/10/10.10.7.0/10.10.7.0.zip"
RUN apt install -y  p7zip-full
RUN 7z x driver.zip


FROM ubuntu:24.04
RUN apt update -y && apt upgrade -y && apt install -y p7zip-full gcc libc-dev wget libusb-1.0-0 && apt clean

COPY --from=kassebuilder /10.10.7.0/installer/deb/libfptr10_10.10.7.0_amd64.deb /10.10.7.0/installer/deb/libfptr10_10.10.7.0_amd64.deb

RUN chmod 777 "/10.10.7.0/installer/deb/libfptr10_10.10.7.0_amd64.deb"
RUN apt install -y "./10.10.7.0/installer/deb/libfptr10_10.10.7.0_amd64.deb"

COPY --from=kassebuilder /online_kasse/cmd/web/web /kasse
COPY --from=kassebuilder /online_kasse/cmd/web/key.pem /
COPY --from=kassebuilder /online_kasse/cmd/web/cert.pem /

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# let's build everything
CMD [ "/entrypoint.sh" ]
