FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    wget \
    curl \
    unzip \
    software-properties-common \
    python3.12 python3.12-dev python3-pip \
    openjdk-8-jdk \
    ca-certificates

RUN wget https://github.com/JetBrains/kotlin/releases/download/v1.8.22/kotlin-compiler-1.8.22.zip \
    && unzip kotlin-compiler-1.8.22.zip -d /opt \
    && ln -s /opt/kotlinc/bin/kotlinc /usr/bin/kotlinc \
    && rm kotlin-compiler-1.8.22.zip

ARG GRADLE_VERSION=8.1.1
RUN wget https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-bin.zip \
    && unzip gradle-${GRADLE_VERSION}-bin.zip -d /opt \
    && ln -s /opt/gradle-${GRADLE_VERSION}/bin/gradle /usr/bin/gradle \
    && rm gradle-${GRADLE_VERSION}-bin.zip

WORKDIR /app

COPY build.gradle settings.gradle ./
COPY src ./src

RUN gradle --no-daemon build || return 0

CMD ["gradle", "run", "--no-daemon"]

