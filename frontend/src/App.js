import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Produkty from './components/Produkty';
import Koszyk from './components/Koszyk';
import Platnosci from './components/Platnosci';
import Login from './components/Login';
import Register from './components/Register';
import AuthCallback from './components/AuthCallback';
import ProtectedRoute from './components/ProtectedRoute';
import { AuthProvider, useAuth } from './context/AuthContext';
import './App.css';

const NavBar = () => {
  const { currentUser, logout } = useAuth();
  
  return (
    <nav style={{ margin: '1em 0', display: 'flex', justifyContent: 'space-between' }}>
      <div>
        <Link to="/">Produkty</Link> |{' '}
        <Link to="/cart">Koszyk</Link> |{' '}
        <Link to="/checkout">Płatności</Link>
      </div>
      <div>
        {currentUser ? (
          <>
            <span>Hello, {currentUser.name} | </span>
            <button onClick={logout} className="btn-link">Logout</button>
          </>
        ) : (
          <>
            <Link to="/login">Login</Link> |{' '}
            <Link to="/register">Register</Link>
          </>
        )}
      </div>
    </nav>
  );
};

const AppRoutes = () => {
  return (
    <>
      <NavBar />
      <Routes>
        <Route path="/" element={<Produkty />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/auth/callback" element={<AuthCallback />} />
        <Route path="/cart" element={
          <ProtectedRoute>
            <Koszyk />
          </ProtectedRoute>
        } />
        <Route path="/checkout" element={
          <ProtectedRoute>
            <Platnosci />
          </ProtectedRoute>
        } />
      </Routes>
    </>
  );
};

export default function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <div className="app-container">
          <AppRoutes />
        </div>
      </AuthProvider>
    </BrowserRouter>
  );
}
