package com.example.bot

import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.application.*
import io.ktor.server.routing.*
import io.ktor.server.response.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.request.*
import kotlinx.serialization.Serializable


@Serializable
data class SlackEvent(
        val token: String? = null,
        val challenge: String? = null,
        val type: String? = null
)

fun main() {
    val botToken = System.getenv("DISCORD_BOT_TOKEN") ?: error("Missing token")
    val channelId = System.getenv("DISCORD_CHANNEL_ID") ?: "414517520188440586"
    val slackToken = System.getenv("SLACK_BOT_TOKEN") ?: error("Missing token")
    val slackChannel = System.getenv("SLACK_CHANNEL_ID") ?: "#general"

    DiscordClient.sendMessage("Bot $botToken", channelId, "Hello from Ktor client!")

    val server = embeddedServer(Netty, port = 8080) {
        install(ContentNegotiation) {
            json()
        }
        routing {
            get("/") {
                call.respondText("Ktor Discord&Slack Bot is running (sending + receiving) ...")
            }
            post("/slack/events") {
                val event = call.receive<SlackEvent>()

                if (event.type == "url_verification" && event.challenge != null) {
                    call.respondText(event.challenge)
                } else {

                    println("Otrzymano event Slack: $event")
                    call.respondText("OK")
                }
            }

            get("/slack/send") {
                SlackIntegration.sendMessage(slackToken, slackChannel, "Hello from Kotlin Slack Bot!")
                call.respondText("Wysłano wiadomość na Slack!")
            }
        }


    }

    server.start(wait = false)



    Bot.startBot(botToken)
}
