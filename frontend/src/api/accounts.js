import api from './client';

export const getAccounts = async () => {
  const response = await api.get('/accounts');
  return response.data.accounts;
};

export const transferFunds = async (fromAccount, toAccount, amount) => {
  await api.post('/transfer', { fromAccount, toAccount, amount });
};