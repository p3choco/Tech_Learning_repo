plugins {
    kotlin("jvm") version "1.8.20"
    id("io.ktor.plugin") version "2.2.4"
}

application {
    mainClass.set("com.example.bot.ApplicationKt")
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor:ktor-client-auth:2.2.4")
    implementation("dev.kord:kord-core:0.8.0")

    implementation("io.ktor:ktor-server-core:2.2.4")
    implementation("io.ktor:ktor-server-netty:2.2.4")

    implementation("io.ktor:ktor-serialization-kotlinx-json:2.2.4")
    implementation("io.ktor:ktor-server-content-negotiation:2.2.4")

    implementation("io.ktor:ktor-client-core:2.2.4")
    implementation("io.ktor:ktor-client-cio:2.2.4")

    implementation("dev.kord:kord-core:0.8.0")

    implementation("ch.qos.logback:logback-classic:1.2.11")

    testImplementation("io.ktor:ktor-server-tests:2.2.4")
    testImplementation("org.jetbrains.kotlin:kotlin-test-junit:1.8.20")
}
