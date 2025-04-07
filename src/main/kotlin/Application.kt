package com.example.bot

import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.application.*
import io.ktor.server.routing.*
import io.ktor.server.response.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.request.*

fun main() {
    val botToken = System.getenv("DISCORD_BOT_TOKEN") ?: "Bot MTM1ODkzMTc3NTMzNzc5MTU3OQ.G1QnpC.S4TD_3WrQWbsgJS6rP7EcA2v7c2mmiHfl-fV6s"
    val channelId = System.getenv("DISCORD_CHANNEL_ID") ?: "414517520188440586"

    DiscordClient.sendMessage(botToken, channelId, "Hello from Ktor client!")

    embeddedServer(Netty, port = 8080) {
        install(ContentNegotiation) {
            json()
        }
        routing {
            get("/") {
                call.respondText("Ktor Discord Client – wysłano wiadomość testową.")
            }
        }
    }.start(wait = true)
}
