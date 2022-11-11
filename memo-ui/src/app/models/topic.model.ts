import { User } from "./user.model";

export interface Topic {
  id: string;
  title: string;
  author_id: string;
  author: User;
  description: string;
  // repetition?: Schedule[];
}



export interface FilterParamsTopic {
  from_date: number;
  to_date: number;
}

export interface IFilterTopic {
  name: string;
  value: {
    from_date: number;
    to_date: number;
  }
}
