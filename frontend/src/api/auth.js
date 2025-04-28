import api from './client';

export const login = async (email, password) => {
  try {
    const response = await api.post('/login', { email, password });
    localStorage.setItem('bank_token', response.data.token);
    return response.data.user;
  } catch (error) {
    throw error.response?.data?.message || 'Login failed';
  }
};

export const logout = () => {
  localStorage.removeItem('bank_token');
};