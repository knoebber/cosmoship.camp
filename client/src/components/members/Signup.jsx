import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { useFetch, useNotification } from '../hooks';

const emptyMember = {
  email: '',
  name: '',
  phone: '',
  source: '',
};

export default function Signup(props) {
  const [sources, setSources] = useState([]);
  const [member, setMember] = useState(emptyMember);
  const fetch = useFetch();
  const setNotification = useNotification();

  const { afterCreate } = props;

  useEffect(() => {
    (async () => {
      const result = await fetch('/members/sources');
      if (result && result.length) {
        setSources(result);
        setMember((m) => ({ ...m, source: result[0] }));
      }
    })();
  }, [fetch]);

  async function onSubmit(e) {
    e.preventDefault();
    let strippedPhone = member.phone.replaceAll(/\D/g, '');
    if (strippedPhone.length === 10) strippedPhone = `1${strippedPhone}`;
    if (strippedPhone) strippedPhone = `+${strippedPhone}`;

    const result = await fetch('/members', {
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
      <label>
        Source:
        <select
          style={{ float: 'right' }}
          name="source"
          value={member.source}
          onChange={onChange}
        >
          {sources.map((s) => <option key={s} value={s}>{s}</option>)}
        </select>
      </label>
      <button type="submit">Submit</button>
    </form>
  );
}

Signup.propTypes = {
  afterCreate: PropTypes.func.isRequired,
};
