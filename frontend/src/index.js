import React from 'react';
import { createRoot } from 'react-dom/client';

import './axiosConfig';

import App from './App';
import './index.css';
import { CartProvider } from './context/CartContext';

const container = document.getElementById('root');
const root = createRoot(container);
root.render(
    <CartProvider>
        <App />
    </CartProvider>
);
