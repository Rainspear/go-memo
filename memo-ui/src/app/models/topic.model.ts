import { User } from "./user.model";

export interface Topic {
  id: string;
  title: string;
  author_id: string;
  author: User;
  description: string;
  created_date?: number;
  last_update?: number;
  // repetition?: Schedule[];
}

export interface ParamsCreateTopic {
  title: string;
  description: string;

}

export interface IFilterTopic {
  name: string;
  value: {
    from_date: number;
    to_date: number;
  }
}
