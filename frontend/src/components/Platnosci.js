import React, { useState } from 'react';
import axios from 'axios';
import { useCart, useCartDispatch } from '../context/CartContext';

const Platnosci = () => {
    const cart = useCart();
    const dispatch = useCartDispatch();
    const [name, setName]   = useState('');
    const [email, setEmail] = useState('');
    const [status, setStatus] = useState(null);

    const total = cart.reduce((sum, i) => sum + i.price * i.qty, 0);

    const handleSubmit = async e => {
        e.preventDefault();
        try {
            await axios.post('http://localhost:8080/payments', {
                customer: { name, email },
                items: cart
            });
            setStatus('success');
            dispatch({ type: 'CLEAR_CART' });  // czyszczenie koszyka po sukcesie
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
                <button type="submit">
                    Zapłać {total.toFixed(2)} zł
                </button>
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
