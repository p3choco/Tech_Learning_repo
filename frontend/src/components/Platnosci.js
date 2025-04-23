// src/components/Platnosci.js
import React, { useState } from 'react';
import axios from 'axios';

const Platnosci = () => {
    const [name, setName]   = useState('');
    const [email, setEmail] = useState('');
    const [status, setStatus] = useState(null);

    const handleSubmit = async e => {
        e.preventDefault();
        try {
            const payload = {
                customer: { name, email },
                items: []
            };
            await axios.post('http://localhost:8080/payments', payload);
            setStatus('success');
        } catch (err) {
            console.error(err);
            setStatus('error');
        }
    };

    return (
        <div>
            <h2>Płatności</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Imię:</label><br/>
                    <input
                        type="text"
                        value={name}
                        onChange={e => setName(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Email:</label><br/>
                    <input
                        type="email"
                        value={email}
                        onChange={e => setEmail(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Wyślij płatność</button>
            </form>

            {status === 'success' && (
                <p style={{ color: 'green' }}>Płatność zrealizowana pomyślnie!</p>
            )}
            {status === 'error' && (
                <p style={{ color: 'red' }}>Błąd przy wysyłaniu płatności.</p>
            )}
        </div>
    );
};

export default Platnosci;
