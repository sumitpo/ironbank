import { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';
import { getAccounts } from '../api/accounts';
import AccountList from '../components/AccountList';
import TransferForm from '../components/TransferForm';

export default function Dashboard() {
  const { user, logout } = useAuth();
  const [accounts, setAccounts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchAccounts = async () => {
      try {
        const data = await getAccounts();
        setAccounts(data);
      } catch (error) {
        console.error('Failed to fetch accounts:', error);
      } finally {
        setLoading(false);
      }
    };
    fetchAccounts();
  }, []);

  if (loading) return <div>Loading accounts...</div>;

  return (
    <div className="dashboard">
      <header>
        <h1>Welcome, {user?.name}</h1>
        <button onClick={logout}>Logout</button>
      </header>
      
      <section className="accounts-section">
        <h2>Your Accounts</h2>
        <AccountList accounts={accounts} />
      </section>
      
      <section className="transfer-section">
        <h2>Transfer Funds</h2>
        <TransferForm accounts={accounts} />
      </section>
    </div>
  );
}