package com.example.bot

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.coroutines.runBlocking

object DiscordClient {

    private val httpClient = HttpClient {
        install(ContentNegotiation) {
            json()
        }
    }

    fun sendMessage(botToken: String, channelId: String, content: String) = runBlocking {
        val response: HttpResponse = httpClient.post("https://discord.com/api/channels/$channelId/messages") {
            header("Authorization", botToken)
            contentType(ContentType.Application.Json)
            setBody(mapOf("content" to content))
        }
        if (response.status.isSuccess()) {
            println("Wiadomość wysłana: $content")
        } else {
            println("Błąd wysyłania wiadomości: ${response.status} - ${response.bodyAsText()}")
        }
    }
}
