import { Topic } from "./topic.model";
import { User } from "./user.model";

export interface Schedule {
  topic_id: string;
  author_id: string;
  author: User;
  topic: Topic
  time: number;
  level: string;
  status: string;
}