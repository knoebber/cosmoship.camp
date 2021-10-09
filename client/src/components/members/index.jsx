import React, { useState, useCallback } from 'react';
import Search from './Search';
import Signup from './Signup';

export default function Members() {
  const [query, setQuery] = useState('');
  const [updateCount, setUpdateCount] = useState(0);

  const afterCreate = useCallback(() => {
    setQuery('');
    setUpdateCount((c) => c + 1);
  }, []);

  return (
    <>
      <h2>Members</h2>
      <Signup afterCreate={afterCreate} />
      <input
        style={{ alignSelf: 'start', marginTop: '40px', marginBottom: '10px' }}
        onChange={({ target: { value } }) => setQuery(value)}
        placeholder="search"
        type="search"
        value={query}
      />
      <Search
        query={query}
        updateCount={updateCount}
        setUpdateCount={setUpdateCount}
      />
    </>
  );
}
