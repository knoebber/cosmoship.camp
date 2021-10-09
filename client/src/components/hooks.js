import {
  useCallback,
  useContext,
  useEffect,
  useRef,
} from 'react';
import { NotificationContext } from '../providers/NotificationProvider';

const apiVersion = 'v1';

export function useNotification() {
  const { setNotification } = useContext(NotificationContext);
  return setNotification;
}

export function useFetch() {
  const { setNotification } = useContext(NotificationContext);
  const abortControllerRef = useRef(null);

  // Abort ongoing requests when the component unmounts.
  useEffect(() => {
    abortControllerRef.current = new AbortController();
    return () => abortControllerRef.current.abort();
  }, []);

  return useCallback(async (route, options) => {
    try {
      const response = await fetch(`/api/${apiVersion}${route}`, {
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
      return json.data;
    } catch (error) {
      if (error.name === 'AbortError') return undefined;
      setNotification('Unexpected error');
      console.error('useFetch:', error);
      return null;
    }
  }, [setNotification]);
}
