import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Produkty = () => {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        axios.get('http://localhost:8080/products')
            .then(res => {
                setProducts(res.data);
                setLoading(false);
            })
            .catch(err => {
                console.error(err);
                setError('Nie udało się pobrać produktów');
                setLoading(false);
            });
    }, []);

    if (loading) return <p>Ładowanie produktów…</p>;
    if (error)   return <p style={{ color: 'red' }}>{error}</p>;

    return (
        <div>
            <h2>Produkty</h2>
            <ul>
                {products.map(p => (
                    <li key={p.id}>
                        <strong>{p.name}</strong> — {p.price.toFixed(2)} zł
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Produkty;
