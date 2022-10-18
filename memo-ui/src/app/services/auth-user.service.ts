import { EventEmitter, Injectable, Output } from '@angular/core';
import { User } from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class AuthUserService {
  user?: User
  @Output() logged = new EventEmitter<User>();

  onGetUser (user: User) {
    this.user = user
    this.logged.emit(user)
  }

  constructor() { }
}
