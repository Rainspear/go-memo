export interface User {
  email: string;
  name: string;
  avatar?: string;
  created_at: number;
}

export interface UpdateUser {
  email?: string;
  name?: string;
  avatar?: string;
  password?: string;
  created_at: number;
}

export interface ParamsCreateUser {
  email: string;
  name: string;
  password: string;
}

export interface ParamsUpdateUser {
  email?: string;
  name?: string;
  password?: string;
}

export interface ParamsLoginUser {
  email: string;
  password: string;
}