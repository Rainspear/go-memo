export interface User {
  email: string;
  name: string;
  avatar?: string;
  created_at: string;
}

export interface ParamsPostUser {
  email: string;
  name: string;
  password: string;
}