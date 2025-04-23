import React from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { useCart, useCartDispatch } from '../context/CartContext';

function Koszyk() {
    const cart = useCart();
    const dispatch = useCartDispatch();

    const saveCart = async () => {
        try {
            await axios.post('http://localhost:8080/carts', { items: cart });
            alert('Koszyk zapisany na serwerze');
        } catch (e) {
            console.error(e);
            alert('Błąd przy przesyłaniu koszyka');
        }
    };

    if (cart.length === 0) return <p>Twój koszyk jest pusty</p>;

    const total = cart.reduce((sum, i) => sum + i.price * i.qty, 0);

    return (
        <div>
            <h2>Twój Koszyk</h2>
            <ul>
                {cart.map(i => (
                    <li key={i.id}>
                        {i.name} x {i.qty} – {(i.price * i.qty).toFixed(2)} zł
                    </li>
                ))}
            </ul>
            <p><strong>Razem: {total.toFixed(2)} zł</strong></p>
            <button onClick={saveCart}>Zapisz koszyk</button>{' '}
            <Link to="/checkout">
                <button>Przejdź do płatności</button>
            </Link>
        </div>
    );
}

export default Koszyk;
