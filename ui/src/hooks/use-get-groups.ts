import { useState, useEffect } from 'react';
import GoHueDataApi from '../models/GoHueDataApi';
import Group from '../models/Group';

interface Results {
  groups?: Group[];
  fetching: boolean;
  error?: string;
}

function useGetGroups(): Results {
  const [results, setResults] = useState<Results>({ fetching: true });

  useEffect(() => {
    async function fetchGroups() {
      try {
        const result = await GoHueDataApi.getLiveGroups();
        setResults({ groups: result.groups, fetching: false })
      } catch (err: any) {
        console.error('failed to fetch groups', err);
        setResults({ fetching: false, error: err.message });
      }
    }

    fetchGroups()
  }, []);

  return results;
}

export default useGetGroups;
