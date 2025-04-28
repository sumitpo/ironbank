export default function AccountList({ accounts }) {
    return (
      <div className="account-list">
        {accounts.length === 0 ? (
          <p>No accounts found</p>
        ) : (
          <table>
            <thead>
              <tr>
                <th>Account Number</th>
                <th>Type</th>
                <th>Balance</th>
              </tr>
            </thead>
            <tbody>
              {accounts.map(account => (
                <tr key={account.id}>
                  <td>{account.number}</td>
                  <td>{account.type}</td>
                  <td>${account.balance.toFixed(2)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    );
  }