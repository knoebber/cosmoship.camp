import React, { useState } from 'react';

export default function Signup() {
  const [email, setEmail] = useState('');

  function onSubmit(e) {
    e.preventDefault();
  }

  function onChange({ target: { value } }) {
    setEmail(value);
  }

  return (
    <>
      <h1>Signup</h1>
      <form onSubmit={onSubmit}>
        <label>
          Email
          <input
            onChange={onChange}
            required
            type="email"
            value={email}
          />
        </label>
        <button type="submit">Submit</button>
      </form>
    </>
  );
}
