FROM mysql:8-debian

RUN rm /etc/apt/sources.list.d/mysql.list && \
    apt update && apt install -y curl && apt clean && \
    curl -sSfL https://repo.mysql.com/RPM-GPG-KEY-mysql-2023 | gpg --import && \
    gpg --batch --export "B7B3B788A8D3785C" > /etc/apt/keyrings/mysql.gpg && \
    echo 'deb [ signed-by=/etc/apt/keyrings/mysql.gpg ] http://repo.mysql.com/apt/debian/ bullseye mysql-8.0' > /etc/apt/sources.list.d/mysql.list

RUN apt-get update && \
    apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8

ENV TZ=Asia/Tokyo

ENV LC_ALL ja_JP.UTF-8
