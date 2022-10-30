export interface Topic {
  id: string;
  title: string;
  description: string;
  repetition?: Repetition[];
}

export interface Repetition {
  time: number;
  level: string;
  status: string;
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
