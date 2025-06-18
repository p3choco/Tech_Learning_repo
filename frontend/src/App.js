import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchItems = async () => {
      try {
        console.log('Fetching items from backend...');
        const API = process.env.REACT_APP_API_URL || "http://34.116.182.229:8000";
        const response = await axios.get(`${API}/items`);
        console.log('Received items:', response.data);
        setItems(response.data);
        setLoading(false);
      } catch (err) {
        console.error('Error fetching items:', err);
        setError('Error fetching items. Please make sure the backend is running.');
        setLoading(false);
      }
    };

    fetchItems();
  }, []);

  if (loading) {
    return (
      <div className="loading-container">
        <div className="loading-spinner"></div>
        <p>Loading items...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="error-container">
        <h2>Error</h2>
        <p>{error}</p>
        <button onClick={() => window.location.reload()}>Retry</button>
      </div>
    );
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>Welcome to Our App</h1>
        <p className="subtitle">Here are our featured items:</p>
        
        <div className="items-grid">
          {items.map((item) => (
            <div key={item.id} className="item-card">
              <h2>{item.name}</h2>
              <p>{item.description}</p>
              <div className="item-footer">
                <span className="item-id">ID: {item.id}</span>
              </div>
            </div>
          ))}
        </div>

        {items.length === 0 && (
          <div className="no-items">
            <p>No items available at the moment.</p>
          </div>
        )}
      </header>
    </div>
  );
}

export default App; 