import { Topic } from "./topic.model";
import { User } from "./user.model";

export interface Memo {
  id: string;
  author_id: string;
  author: User;
  topic_id: string;
  topic: Topic;
  question: string;
  content: string;
  answer?: string[]
}

export interface ParamsCreateMemo {
  answer?: string[];
  question: string;
  content: string;
  topic_id: string;
}

export interface ParamsFilterMemo {
  topic_id: string;
}