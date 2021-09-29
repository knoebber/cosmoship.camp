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
    let strippedPhone = user.phone.replaceAll(/\D/g, '');
    if (strippedPhone.length === 10) strippedPhone = `1${strippedPhone}`;
    if (strippedPhone) strippedPhone = `+${strippedPhone}`;

    const result = await fetch('/api/v1/users', {
      method: 'POST',
      body: JSON.stringify({
        ...user,
        phone: strippedPhone,
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
      <h2>Signup</h2>
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
