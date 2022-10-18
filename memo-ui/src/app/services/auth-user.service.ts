import { EventEmitter, Injectable, OnDestroy, Output } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, Subject, Subscription } from 'rxjs';
import { User } from '../models/user.model';
import { ApiService } from './api.service';

@Injectable({
  providedIn: 'root'
})
export class AuthUserService implements OnDestroy {
  user?: User
  userSubscription?: Subscription
  @Output() loggedUser = new Subject<User>;

  isAuthenicated(): Observable<boolean> {
    return new Observable<boolean>(observer => {
      const token = localStorage.getItem('token')
      if (token) {
        if (!this.user) {
          this.apiService.currentUser().subscribe((res: any) => {
            if (res.data) {
              this.user = res.data
              this.loggedUser.next(res.data)
              observer.next(true)
              observer.complete();
            } else {
              this.router.navigate(['/auth', "signin"])
              localStorage.setItem('token', '')
              observer.next(false);
              observer.complete();
            }
            return true;
          }, err => {
            this.router.navigate(['/auth', "signin"])
            observer.next(false);
            observer.complete();
            localStorage.setItem('token', '')
          })
        } else {
          observer.next(true)
          observer.complete();
        }
        return;
      }
      this.router.navigate(['/auth', "signin"])
      observer.next(false);
      observer.complete();
      return;
    })
  }
  constructor(private apiService: ApiService, private router: Router) {
    this.userSubscription = this.loggedUser.subscribe(user => {
      this.user = user;
    })
  }

  ngOnDestroy(): void {
    if (this.userSubscription) this.userSubscription.unsubscribe();
  }
}
