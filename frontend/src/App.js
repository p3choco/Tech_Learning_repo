import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Produkty  from './components/Produkty';
import Koszyk    from './components/Koszyk';
import Platnosci from './components/Platnosci';

export default function App() {
    return (
        <BrowserRouter>
            <nav style={{ margin: '1em 0' }}>
                <Link to="/">Produkty</Link> |{' '}
                <Link to="/cart">Koszyk</Link> |{' '}
                <Link to="/checkout">Płatności</Link>
            </nav>
            <Routes>
                <Route path="/"         element={<Produkty />} />
                <Route path="/cart"     element={<Koszyk />} />
                <Route path="/checkout" element={<Platnosci />} />
            </Routes>
        </BrowserRouter>
    );
}
