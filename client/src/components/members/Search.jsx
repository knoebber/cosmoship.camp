import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { useFetch, useNotification } from '../hooks';

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

export default function Search(props) {
  const [members, setMembers] = useState([]);
  const fetch = useFetch();
  const notify = useNotification();

  const {
    query,
    setUpdateCount,
    updateCount,
  } = props;

  useEffect(() => {
    (async () => {
      const result = await fetch('/members');
      if (result) setMembers(result);
    })();
  }, [fetch, updateCount]);

  async function deleteMember(id) {
    const result = await fetch(`/members/${id}`, { method: 'DELETE' });
    if (result) {
      setUpdateCount(updateCount + 1);
      notify('deleted member');
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
    <table>
      <thead>
        <tr>
          <th>Email</th>
          <th>Name</th>
          <th>Phone</th>
          <th>Source</th>
          <th>Created</th>
          <th>&nbsp;</th>
        </tr>
      </thead>
      <tbody>
        {members
          .filter(filterRows)
          .map((m) => (
            <tr key={m.id}>
              <td>{m.email}</td>
              <td>{m.name}</td>
              <td>{formatPhone(m.phone)}</td>
              <td>{m.source}</td>
              <td>{new Date(m.createdAt).toLocaleString()}</td>
              <td>
                <button
                  className="link-button"
                  type="button"
                  onClick={() => deleteMember(m.id)}
                >
                  x
                </button>
              </td>
            </tr>
          ))}
      </tbody>
    </table>
  );
}

Search.propTypes = {
  query: PropTypes.string.isRequired,
  updateCount: PropTypes.number.isRequired,
  setUpdateCount: PropTypes.func.isRequired,
};
