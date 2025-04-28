import axios from 'axios';

const api = axios.create({
  baseURL: process.env.NODE_ENV === 'development' 
    ? 'http://localhost:8080/api' 
    : '/api',
  timeout: 5000,
});

// Request interceptor for auth token
api.interceptors.request.use(config => {
  const token = localStorage.getItem('bank_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;