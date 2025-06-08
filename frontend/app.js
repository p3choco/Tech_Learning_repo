const chatDiv = document.getElementById('chat');

async function loadStarters() {
    try {
        const res = await fetch('/starters');
        const data = await res.json();
        const openings = document.getElementById('openings');
        data.openings.forEach(txt => openings.innerHTML += `<li>${txt}</li>`);
        const closings = document.getElementById('closings');
        data.closings.forEach(txt => closings.innerHTML += `<li>${txt}</li>`);
    } catch (err) {
        console.error('Błąd ładowania otwarć/zamknięć:', err);
    }
}

document.getElementById('chatForm').addEventListener('submit', async e => {
    e.preventDefault();
    const msg = document.getElementById('message').value;
    chatDiv.innerHTML += `<div class="user"><b>Ty:</b> ${msg}</div>`;
    document.getElementById('message').value = '';

    try {
        const res = await fetch('/chat', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ message: msg })
        });
        let data;
        try {
            data = await res.json();
        } catch (parseError) {
            chatDiv.innerHTML += `<div class="bot"><b>Bot:</b> Błąd parsowania odpowiedzi serwera.</div>`;
            return;
        }
        const { reply, sentiment } = data;
        chatDiv.innerHTML += `
            <div class="bot">
                <b>Bot:</b> ${reply}
                <div class="sentiment">Sentyment: ${sentiment || '-'}</div>
            </div>`;
    } catch (err) {
        chatDiv.innerHTML += `<div class="bot"><b>Bot:</b> Błąd sieci: ${err.message}</div>`;
    } finally {
        chatDiv.scrollTop = chatDiv.scrollHeight;
    }
});

loadStarters();