from flask import Flask, request, jsonify, send_from_directory
from dotenv import load_dotenv
from utils import getcategories, getproducts, analyze_sentiment
import os
import openai
import random
import json

load_dotenv()

app = Flask(__name__, static_folder='frontend', static_url_path='', template_folder='frontend')

OPENAI_MODEL = os.getenv("OPENAI_MODEL", "gpt-4")
openai.api_key = "PLACEHOLDER"

OPENINGS = [
    "Witamy w naszym sklepie z ubraniami!",
    "Miło Cię widzieć w naszym sklepie odzieżowym!",
    "Wspaniale, że odwiedzasz nasz sklep z modą!",
    "Witaj w sklepie z ubraniami, gdzie styl jest najważniejszy!",
    "Cieszymy się, że wybrałeś nasz sklep z odzieżą!"
]
CLOSINGS = [
    "Dziękujemy za rozmowę! Zapraszamy ponownie!",
    "Miłego dnia i udanych zakupów!",
    "Do zobaczenia ponownie w naszym sklepie!",
    "Cieszymy się, że mogliśmy pomóc!",
    "W razie pytań – jesteśmy do dyspozycji!"
]

SYSTEM_PROMPT = (
    "Jesteś asystentem sklepu z ubraniami. "
    "Odpowiadaj wyłącznie na pytania dotyczące sklepu oraz produktów odzieżowych. "
    "Używaj funkcji getcategories() i getproducts(category) do pobierania danych w razie potrzeby. "
    "Zawsze uprzejmie i z szacunkiem. "
)

@app.route('/chat', methods=['POST'])
def chat():
    try:
        data = request.get_json(force=True)
        user_message = data.get('message', '').strip()
        lower_msg = user_message.lower()

        if analyze_sentiment(user_message) == 'negatywny':
            return jsonify(
                {'reply': 'Proszę, unikaj negatywnych sformułowań. Skupmy się na pozytywnych aspektach naszej oferty.',
                 'sentiment': 'neutralny'})

        if any(kw in lower_msg for kw in ['koniec','do widzenia','pa']):
            return jsonify({'reply': random.choice(CLOSINGS), 'sentiment': 'neutralny'})

        greeting_keywords = ['hej','cześć','witaj']
        if any(lower_msg.startswith(kw) for kw in greeting_keywords):
            return jsonify({'reply': random.choice(OPENINGS), 'sentiment': 'neutralny'})

        allowed = [
            'sklep','ubranie','ubrania','odzież','koszulka','koszulki',
            'spodnie','buty','bluza','bluzy','bluzka','kategorie'
        ]
        if not any(w in lower_msg for w in allowed):
            return jsonify({'reply': 'Możesz pytać wyłącznie o asortyment i usługi naszego sklepu z odzieżą.', 'sentiment': 'neutralny'})

        messages = [
            {"role": "system", "content": SYSTEM_PROMPT},
            {"role": "user", "content": user_message}
        ]
        functions = [
            {"name": "getcategories", "description": "Lista kategorii odzieży", "parameters": {"type": "object", "properties": {}}},
            {"name": "getproducts", "description": "Produkty w podanej kategorii", "parameters": {"type": "object","properties": {"category":{"type":"string"}},"required":["category"]}}
        ]

        response = openai.chat.completions.create(
            model=OPENAI_MODEL,
            messages=messages,
            functions=functions,
            function_call="auto"
        )

        choice = response.choices[0]
        msg = choice.message

        if msg.function_call:
            fname = msg.function_call.name
            args = {}
            if msg.function_call.arguments:
                args = json.loads(msg.function_call.arguments)
            result = getcategories() if fname == 'getcategories' else getproducts(args.get('category',''))

            followup = openai.chat.completions.create(
                model=OPENAI_MODEL,
                messages=[
                    {"role": "system", "content": SYSTEM_PROMPT},
                    {"role": "assistant", "name": fname, "content": result},
                    {"role": "user", "content": user_message}
                ]
            )
            ai_reply = followup.choices[0].message.content.strip()
        else:
            ai_reply = msg.content.strip()

        if analyze_sentiment(ai_reply) == 'negatywny':
            ai_reply = 'Asystent unika negatywnych odpowiedzi. Skupmy się na pozytywnych aspektach sklepu.'

        return jsonify({'reply': ai_reply, 'sentiment': analyze_sentiment(ai_reply)})

    except Exception as e:
        return jsonify({'reply': f'Wystąpił błąd serwera: {str(e)}', 'sentiment': 'neutralny'})

@app.route('/starters', methods=['GET'])
def starters():
    return jsonify({'openings': OPENINGS, 'closings': CLOSINGS})

@app.route('/')
def index():
    return send_from_directory('frontend', 'index.html')

if __name__ == '__main__':
    app.run(debug=True)