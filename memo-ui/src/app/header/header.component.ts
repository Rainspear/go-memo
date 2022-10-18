import { Component, EventEmitter, OnDestroy, OnInit, Output } from '@angular/core';
import { Observable, Subscription } from 'rxjs';
import { User } from '../models/user.model';
import { ApiService } from '../services/api.service';
import { AuthUserService } from '../services/auth-user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit, OnDestroy {
  collapsed: boolean = true;
  user?: User;
  userSubscription?: Subscription;
  @Output() featureSelected = new EventEmitter<string>();
  constructor(private authService: AuthUserService, private apiService: ApiService) {
    this.userSubscription = this.authService.loggedUser.subscribe(user => {
      this.user = user
    })
  }

  // onClickSelectNav(route: string) {
  // this.featureSelected.emit(route);
  // }

  onClickLogout() {
    this.apiService.logOutUser().subscribe(user => {})
  }

  ngOnInit(): void {
    if (!this.user)
      this.apiService.currentUser().subscribe((res: any) => {
        if (res.data) {
          this.user = res.data;
          this.authService.loggedUser.next(res.data)
        }
      })
  }

  ngOnDestroy(): void {
    if (this.userSubscription) this.userSubscription.unsubscribe();
  }
}
