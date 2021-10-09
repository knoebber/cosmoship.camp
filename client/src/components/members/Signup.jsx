import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { useFetch, useNotification } from '../hooks';

const emptyMember = {
  email: '',
  name: '',
  phone: '',
};

export default function Signup(props) {
  const [member, setMember] = useState(emptyMember);
  const fetch = useFetch();
  const setNotification = useNotification();

  const { afterCreate } = props;

  async function onSubmit(e) {
    e.preventDefault();
    let strippedPhone = member.phone.replaceAll(/\D/g, '');
    if (strippedPhone.length === 10) strippedPhone = `1${strippedPhone}`;
    if (strippedPhone) strippedPhone = `+${strippedPhone}`;

    const result = await fetch('members', {
      method: 'POST',
      body: JSON.stringify({
        ...member,
        phone: strippedPhone,
      }),
    });
    if (result) {
      setMember(emptyMember);
      setNotification(`added "${member.email}"`);
      afterCreate();
    }
  }

  function onChange({ target: { name, value } }) {
    setMember({ ...member, [name]: value });
  }

  return (
    <form onSubmit={onSubmit}>
      <label>
        Name:
        <input
          name="name"
          onChange={onChange}
          value={member.name}
        />
      </label>
      <label>
        Email:
        <input
          name="email"
          onChange={onChange}
          required
          type="email"
          value={member.email}
        />
      </label>
      <label>
        Phone:
        <input
          name="phone"
          onChange={onChange}
          value={member.phone}
        />
      </label>
      <button type="submit">Submit</button>
    </form>
  );
}

Signup.propTypes = {
  afterCreate: PropTypes.func.isRequired,
};
