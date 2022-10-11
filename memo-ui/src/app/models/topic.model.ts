export interface Topic {
  _id: string;
  title: string;
  description: string;
  repetition?: Repetition[];
}

export interface Repetition {
  time: string;
  level: string;
  status: string;
}
