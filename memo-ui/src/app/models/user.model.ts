export interface User {
  email: string;
  name: string;
  avatar?: string;
  created_at: string;
}

export interface ParamsCreateUser {
  email: string;
  name: string;
  password: string;
}

export interface ParamsLoginUser {
  email: string;
  password: string;
}