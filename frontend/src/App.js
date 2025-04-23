import React from 'react';
import Produkty   from './components/Produkty';
import Platnosci from './components/Platnosci';

function App() {
    return (
        <div style={{ padding: 20 }}>
            <h1>Moja Aplikacja Sklep</h1>
            <Produkty />
            <hr/>
            <Platnosci />
        </div>
    );
}

export default App;
