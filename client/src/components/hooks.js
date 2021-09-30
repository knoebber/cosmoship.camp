import {
  useCallback,
  useContext,
  useEffect,
  useRef,
} from 'react';
import { NotificationContext } from '../providers/NotificationProvider';

export function useNotification() {
  const { setNotification } = useContext(NotificationContext);
  return setNotification;
}

export function useFetch() {
  const { setNotification } = useContext(NotificationContext);
  const abortControllerRef = useRef(null);

  // Abort any current requests when the component unmounts.
  useEffect(() => {
    abortControllerRef.current = new AbortController();
    return () => abortControllerRef.current.abort();
  }, []);

  return useCallback(async (url, options) => {
    try {
      const response = await fetch(`/api/v1/${url}`, {
        ...options,
        signal: abortControllerRef.current.signal,
        headers: { 'Content-Type': 'application/json' },
      });
      if (response.status !== 200) {
        setNotification(response.statusText);
        return null;
      }
      const json = await response.json();
      if (json.message) setNotification(json.message);
      console.log(json);
      return json.data;
    } catch (error) {
      if (error.name === 'AbortError') return undefined;
      setNotification('Unexpected error');
      console.error('fetch:', error);
      return null;
    }
  }, [setNotification]);
}
