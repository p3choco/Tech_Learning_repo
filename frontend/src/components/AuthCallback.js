import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const AuthCallback = () => {
  const navigate = useNavigate();
  const { loading, currentUser, error } = useAuth();
  const [processed, setProcessed] = useState(false);

  useEffect(() => {
    console.log('AuthCallback component mounted, loading:', loading);
    console.log('URL params:', window.location.search);
    console.log('Current user:', currentUser);

    const timeoutId = setTimeout(() => {
      if (!processed) {
        console.log('Timeout reached, navigating to home');
        setProcessed(true);
        navigate('/');
      }
    }, 3000);

    return () => clearTimeout(timeoutId);
  }, []);

  useEffect(() => {
    if (!loading && !processed) {
      console.log('Auth loading complete, user:', currentUser);
      setProcessed(true);
      navigate('/');
    }
  }, [loading, currentUser, navigate, processed]);

  return (
    <div className="auth-callback">
      <h2>Processing authentication...</h2>
      <p>Please wait while we log you in.</p>
      {error && <div className="alert alert-error">{error}</div>}
    </div>
  );
};

export default AuthCallback; 