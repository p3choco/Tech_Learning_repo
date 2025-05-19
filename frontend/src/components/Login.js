import React, { useState } from 'react';
import { useAuth } from '../context/AuthContext';
import { Link, useNavigate } from 'react-router-dom';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { login, loginWithGoogle, loginWithGithub, loading } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    
    try {
      await login(email, password);
      navigate('/');
    } catch (err) {
      setError('Failed to sign in');
      console.error(err);
    }
  };

  const handleGoogleLogin = (e) => {
    e.preventDefault();
    console.log("Google login button clicked");
    loginWithGoogle();
  };

  const handleGithubLogin = (e) => {
    e.preventDefault();
    console.log("GitHub login button clicked");
    loginWithGithub();
  };

  return (
    <div className="auth-container">
      <h2>Login</h2>
      {error && <div className="alert alert-error">{error}</div>}
      
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Email</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        
        <div className="form-group">
          <label>Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        
        <button type="submit" disabled={loading} className="btn btn-primary">
          {loading ? 'Loading...' : 'Login'}
        </button>
      </form>
      
      <div className="oauth-buttons">
        <button onClick={handleGoogleLogin} className="btn btn-google">
          Login with Google
        </button>
        <button onClick={handleGithubLogin} className="btn btn-github">
          Login with GitHub
        </button>
      </div>
      
      <p>
        Don't have an account? <Link to="/register">Register</Link>
      </p>
    </div>
  );
};

export default Login; 