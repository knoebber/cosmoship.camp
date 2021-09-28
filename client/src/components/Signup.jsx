import React, { useState } from 'react';
import { useFetch, useNotification } from './hooks';

const initialState = {
  email: '',
  name: '',
  phone: '',
};

export default function Signup() {
  const fetch = useFetch();
  const setNotification = useNotification();
  const [user, setUser] = useState(initialState);

  async function onSubmit(e) {
    e.preventDefault();
    const result = await fetch('/api/v1/users', {
      method: 'POST',
      body: JSON.stringify({
        ...user,
        phone: user.phone.replaceAll(/\D/g, ''),
      }),
    });
    if (result) {
      setNotification(`added "${user.email}"`);
    }
  }

  function onChange({ target: { name, value } }) {
    setUser({ ...user, [name]: value });
  }

  return (
    <>
      <h1>Signup</h1>
      <form onSubmit={onSubmit}>
        <label>
          Email:
          <input
            name="email"
            onChange={onChange}
            required
            type="email"
            value={user.email}
          />
        </label>
        <label>
          Name:
          <input
            name="name"
            onChange={onChange}
            value={user.name}
          />
        </label>
        <label>
          Phone:
          <input
            name="phone"
            onChange={onChange}
            value={user.phone}
          />
        </label>
        <button type="submit">Submit</button>
      </form>
    </>
  );
}
