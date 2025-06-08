import re

def getcategories():
    return "Bluzy, Spodnie, Koszulki"

def getproducts(category):
    data = {
        'Bluzy': "Bluza Adidas[250zł], Bluza Nike[220zł]",
        'Spodnie': "Jeansy Levi's[300zł], Spodnie dresowe Puma[180zł]",
        'Koszulki': "Koszulka Tommy Hilfiger[150zł], Koszulka Reserved[60zł]"
    }
    return data.get(category, "Brak produktów w tej kategorii.")

def analyze_sentiment(text):
    pos = ['świetnie','super','dziękuję','miło','cieszę','zadowolony','pomóc']
    neg = ['problem','nie działa','zły','reklamacja','nie polecam','niezadowolony']
    low = text.lower()
    if any(w in low for w in pos): return 'pozytywny'
    if any(w in low for w in neg): return 'negatywny'
    return 'neutralny'

def extract_function_call(text):
    m = re.search(r"\[call:(getcategories|getproducts)(?::([^\]]+))?\]", text)
    if not m: return None
    return (m.group(1), m.group(2).strip() if m.group(2) else None)