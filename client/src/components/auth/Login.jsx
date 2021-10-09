import React, { useState } from 'react';

export default function Login() {
  const [state, setState] = useState({ username: '', value: '' });

  function onSubmit(e) {
    e.preventDefault();
  }

  function onChange({ target: { name, value } }) {
    setState({ ...state, [name]: value });
  }

  return (
    <>
      <h2>Login</h2>
      <form onSubmit={onSubmit}>
        <label>
          Username
          <input
            name="username"
            onChange={onChange}
            type="text"
            value={state.username}
          />
        </label>
        <label>
          Password
          <input
            name="password"
            onChange={onChange}
            type="password"
            value={state.password}
          />
        </label>
        <button type="submit">Login</button>
      </form>
    </>
  );
}
