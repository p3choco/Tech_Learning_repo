package com.example.bot

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.coroutines.runBlocking

object SlackIntegration {
    private val httpClient = HttpClient {
        install(ContentNegotiation) {
            json()
        }
    }

    fun sendMessage(token: String, channel: String, text: String) = runBlocking {
        val response: HttpResponse = httpClient.post("https://slack.com/api/chat.postMessage") {
            header("Authorization", "Bearer $token")
            contentType(ContentType.Application.Json)
            setBody(
                    mapOf(
                            "channel" to channel,
                            "text" to text
                    )
            )
        }
        if (response.status.isSuccess()) {
            println("Wysłano wiadomość na Slack: $text")
        } else {
            println("Błąd wysyłania na Slack: ${response.status} - ${response.bodyAsText()}")
        }
    }
}
