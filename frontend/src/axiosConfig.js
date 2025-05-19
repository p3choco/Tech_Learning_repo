import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080';
console.log("Axios baseURL configured as:", axios.defaults.baseURL);

axios.defaults.headers.common['Content-Type'] = 'application/json';

axios.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
      console.log("Added Authorization header:", `Bearer ${token.substring(0, 15)}...`);
    }
    return config;
  },
  error => {
    console.error("Axios request error:", error);
    return Promise.reject(error);
  }
);

axios.interceptors.response.use(
  response => {
    console.log(`Response from ${response.config.url}:`, response.status);
    return response;
  },
  error => {
    console.error("Axios response error:", error.response?.status, error.response?.data);
    return Promise.reject(error);
  }
);

export default axios;
