package com.example.bot

import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import kotlinx.coroutines.runBlocking

private val categories = listOf("Elektronika", "RTV/AGD", "Książki")
private val productsMap = mapOf(
        "Elektronika" to listOf("Laptop", "Smartfon", "Tablet"),
        "RTV/AGD" to listOf("Telewizor", "Lodówka", "Pralka"),
        "Książki" to listOf("Harry Potter", "Władca Pierścieni", "Pan Tadeusz")
)

object Bot {
    fun startBot(botToken: String) = runBlocking {
        val kord = Kord(botToken)

        kord.on<MessageCreateEvent> {
            if (message.author?.isBot == true) return@on

            val content = message.content
            when {
                content.startsWith("!hello") -> {
                    message.channel.createMessage("Cześć, otrzymałem Twoją wiadomość: $content")
                }
                content.startsWith("!kategorie") -> {
                    val response = "Dostępne kategorie: ${categories.joinToString(", ")}"
                    message.channel.createMessage(response)
                }
                content.startsWith("!produkty") -> {
                    val parts = content.split(" ", limit = 2)
                    if (parts.size < 2) {
                        message.channel.createMessage("Podaj kategorię: !produkty <nazwa_kategorii>")
                        return@on
                    }
                    val categoryName = parts[1]
                    val products = productsMap[categoryName]
                    if (products == null) {
                        message.channel.createMessage("Nieznana kategoria '$categoryName'!")
                    } else {
                        val response = "Produkty w '$categoryName': ${products.joinToString(", ")}"
                        message.channel.createMessage(response)
                    }
                }
            }
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }
}
