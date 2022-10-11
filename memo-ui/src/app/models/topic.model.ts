export interface Topic {
  id: string;
  title: string;
  description: string;
  repetition?: Repetition[];
}

export interface Repetition {
  time: string;
  level: string;
  status: string;
}
