package com.example.bot

import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import kotlinx.coroutines.runBlocking

private val categories = listOf("Elektronika", "RTV/AGD", "Książki")

object Bot {
    fun startBot(botToken: String) = runBlocking {
        val kord = Kord(botToken)

        kord.on<MessageCreateEvent> {
            if (message.author?.isBot == true) return@on

            val content = message.content
            when {
                content.startsWith("!bot") -> {
                    message.channel.createMessage("Hej, otrzymałem wiadomość: $content")
                }
                content.startsWith("!kategorie") -> {
                    val response = "Dostępne kategorie: ${categories.joinToString(", ")}"
                    message.channel.createMessage(response)
                }
            }
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }
}
