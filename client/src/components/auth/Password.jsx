import React, { useState } from 'react';
import { useFetch, useNotification } from '../hooks';

const initialState = {
  password: '',
  passwordConfirm: '',
  email: '',
};

export default function Password() {
  const [state, setState] = useState(initialState);
  const notify = useNotification();
  const fetch = useFetch();

  const {
    password,
    passwordConfirm,
    email,
  } = state;

  function onSubmit(e) {
    e.preventDefault();
    if (password !== passwordConfirm) {
      notify("Passwords don't match");
    }

    fetch('/auth/password', { method: 'POST', body: { email, password } });
  }

  function onChange({ target: { name, value } }) {
    setState({ ...state, [name]: value });
  }

  return (
    <>
      <h2>Set Password</h2>
      <form onSubmit={onSubmit}>
        <label>
          Email
          <input
            name="email"
            onChange={onChange}
            type="email"
            value={email}
          />
        </label>
        <label>
          Password
          <input
            name="password"
            onChange={onChange}
            type="password"
            value={password}
          />
        </label>
        <label>
          Password Confirm
          <input
            name="password"
            onChange={onChange}
            type="password"
            value={password}
          />
        </label>
        <button type="submit">Submit</button>
      </form>
    </>
  );
}
