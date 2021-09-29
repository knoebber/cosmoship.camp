import React, { useState, useEffect } from 'react';
import { useFetch } from './hooks';

// Input: +17076412600
// Output: +1 (707) 641-2600
function formatPhone(phone) {
  const result = [];

  for (let i = 0; i < phone.length; i += 1) {
    result.unshift(phone[phone.length - 1 - i]);
    if (i === 3) result.unshift('-');
    else if (i === 6) result.unshift(') ');
    else if (i === 9) result.unshift(' (');
  }

  console.log(result);
  return result.join('');
}

export default function Signup() {
  const [users, setUsers] = useState([]);
  const fetch = useFetch();

  useEffect(() => {
    (async () => {
      const result = await fetch('/api/v1/users');
      console.log(result);
      if (result) setUsers(result);
    })();
  }, [fetch]);

  return (
    <>
      <h2>Users</h2>
      <table>
        <thead>
          <tr>
            <th>Email</th>
            <th>Name</th>
            <th>Phone</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          {users.map((u) => (
            <tr key={u.id}>
              <td>{u.email}</td>
              <td>{u.name}</td>
              <td>{formatPhone(u.phone)}</td>
              <td>{new Date(u.createdAt).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}
