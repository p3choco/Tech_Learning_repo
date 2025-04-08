package com.example.bot

import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import kotlinx.coroutines.runBlocking


object Bot {
    fun startBot(botToken: String) = runBlocking {
        val kord = Kord(botToken)

        kord.on<MessageCreateEvent> {
            if (message.author?.isBot != false) return@on

            val content = message.content

            if (content.startsWith("!bot")) {
                message.channel.createMessage("Hej, otrzymałem wiadomość: $content")
            }
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }
}
