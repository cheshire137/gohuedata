import Group from '../models/Group';

type GroupsResult = {
  groups: Group[];
  page: number;
  totalPages: number;
  totalCount: number;
}

export default GroupsResult;
