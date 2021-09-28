import React, { useState, useEffect } from 'react';
import { useFetch } from './hooks';

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
      <h1>Users</h1>
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
              <td>{u.phone}</td>
              <td>{new Date(u.createdAt).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}
