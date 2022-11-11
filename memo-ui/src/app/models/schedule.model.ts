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
  time_date?: Date; // custom field not from api
}

export interface ParamsFilterSchedule {
  from_date?: number;
  to_date?: number;
  topic_id: string;
}

export interface ParamsCreateSchedule {
  time: number;
  level: "essential" | "important" | "critical" | "major" | "minor";
  status: "success" | "untouch" | "failure" | "skipped";
  topic_id: string;
}
