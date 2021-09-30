import React, { useState, useEffect } from 'react';
import { useFetch, useNotification } from './hooks';

// Precondition: phone is in e164 format.
// Input: '+17076412600'
// Output: '+1 (707) 641-2600'
function formatPhone(phone) {
  const result = [];

  for (let i = 0; i < phone.length; i += 1) {
    result.unshift(phone[phone.length - 1 - i]);
    if (i === 3) result.unshift('-');
    else if (i === 6) result.unshift(') ');
    else if (i === 9) result.unshift(' (');
  }

  return result.join('');
}

const emptyMember = {
  email: '',
  name: '',
  phone: '',
};

export default function Signup() {
  const [query, setQuery] = useState('');
  const [updateCount, setUpdateCount] = useState(0);
  const [member, setMember] = useState(emptyMember);
  const [members, setMembers] = useState([]);
  const fetch = useFetch();
  const setNotification = useNotification();

  useEffect(() => {
    (async () => {
      const result = await fetch('members');
      if (result) setMembers(result);
    })();
  }, [fetch, updateCount]);

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
      setUpdateCount(updateCount + 1);
      setQuery('');
      setNotification(`added "${member.email}"`);
    }
  }

  function onChange({ target: { name, value } }) {
    setMember({ ...member, [name]: value });
  }

  async function deleteMember(id) {
    const result = await fetch(`members/${id}`, { method: 'DELETE' });
    if (result) {
      setUpdateCount(updateCount + 1);
      setNotification('deleted member');
    }
  }

  function filterRows(row) {
    const q = query.toLowerCase();
    if (!query) return true;
    return row.phone.replaceAll(/D/g, '').includes(q)
           || row.email.includes(q)
           || row.name.toLowerCase().includes(q);
  }

  return (
    <>
      <h2>Members</h2>
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
      <div style={{ height: '40px' }} />
      <input
        style={{ alignSelf: 'start', marginBottom: '10px' }}
        onChange={({ target: { value } }) => setQuery(value)}
        placeholder="search"
        type="search"
        value={query}
      />
      <table>
        <thead>
          <tr>
            <th>Email</th>
            <th>Name</th>
            <th>Phone</th>
            <th>Created</th>
            <th>&nbsp;</th>
          </tr>
        </thead>
        <tbody>
          {members
            .filter(filterRows)
            .map((u) => (
              <tr key={u.id}>
                <td>{u.email}</td>
                <td>{u.name}</td>
                <td>{formatPhone(u.phone)}</td>
                <td>{new Date(u.createdAt).toLocaleString()}</td>
                <td>
                  <button
                    className="link-button"
                    type="button"
                    onClick={() => deleteMember(u.id)}
                  >
                    x
                  </button>
                </td>
              </tr>
            ))}
        </tbody>
      </table>
    </>
  );
}
