const API_URL = 'http://localhost:8080/api';

export const authService = {
  login: async (email, password) => {
    const response = await fetch(`${API_URL}/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });
    return response.json();
  },

  getAccounts: async (token) => {
    const response = await fetch(`${API_URL}/accounts`, {
      headers: { 'Authorization': `Bearer ${token}` }
    });
    return response.json();
  },

  transfer: async (transferData, token) => {
    const response = await fetch(`${API_URL}/transfer`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(transferData)
    });
    return response.json();
  }
};
