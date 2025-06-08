import React, { createContext, useState, useContext, useEffect } from 'react';
import axios from '../axiosConfig';

const AuthContext = createContext(null);

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
  const [currentUser, setCurrentUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    
    console.log('AuthContext initialized, path:', window.location.pathname);
    console.log('URL params:', window.location.search);
    
    if (token) {
      console.log('Found existing token in localStorage');
      fetchUserData(token);
    } else {
      setLoading(false);
    }
    
    const urlParams = new URLSearchParams(window.location.search);
    let tokenParam = urlParams.get('token');
    
    console.log('Token from URL (encoded):', tokenParam);
    
    if (tokenParam) {
      try {
        tokenParam = decodeURIComponent(tokenParam);
        console.log('Processing OAuth callback, token decoded:', tokenParam.substring(0, 15) + '...');
        localStorage.setItem('token', tokenParam);
        fetchUserData(tokenParam);
        window.history.replaceState({}, document.title, '/');
      } catch (err) {
        console.error('Error decoding token:', err);
        setError('Authentication failed: Invalid token format');
        setLoading(false);
      }
    }
  }, []);

  const fetchUserData = async (token) => {
    try {
      const res = await axios.get('/user', {
        headers: { Authorization: token }
      });
      setCurrentUser(res.data);
      setError(null);
    } catch (err) {
      console.error('Error fetching user data:', err);
      setError('Failed to authenticate user');
      localStorage.removeItem('token');
    } finally {
      setLoading(false);
    }
  };

  const login = async (email, password) => {
    setLoading(true);
    try {
      const res = await axios.post('/login', { email, password });
      localStorage.setItem('token', res.data.token);
      setCurrentUser(res.data.user);
      return res.data.user;
    } catch (err) {
      setError('Login failed. Please check your credentials.');
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const register = async (name, email, password) => {
    setLoading(true);
    try {
      const res = await axios.post('/register', { name, email, password });
      localStorage.setItem('token', res.data.token);
      setCurrentUser(res.data.user);
      setError(null);
      return res.data.user;
    } catch (err) {
      setError('Registration failed. Please try again.');
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const loginWithGoogle = () => {
    console.log("loginWithGoogle called");
    console.log("Redirecting to:", `${axios.defaults.baseURL}/auth/google/login`);
    window.location.href = `${axios.defaults.baseURL}/auth/google/login`;
    console.log("Redirection attempted");
  };

  const loginWithGithub = () => {
    window.location.href = `${axios.defaults.baseURL}/auth/github/login`;
  };

  const logout = () => {
    localStorage.removeItem('token');
    setCurrentUser(null);
  };

  const value = {
    currentUser,
    loading,
    error,
    login,
    register,
    loginWithGoogle,
    loginWithGithub,
    logout
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}; 