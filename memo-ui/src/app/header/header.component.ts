import { Component, EventEmitter, OnDestroy, OnInit, Output } from '@angular/core';
import { Subscription } from 'rxjs';
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
  constructor(private authService: AuthUserService) {
    this.userSubscription = this.authService.loggedUser.subscribe(user => {
      this.user = user
    })
  }

  onClickLogout() {
    this.authService.logOutAndClearToken();
  }

  ngOnInit(): void {
    this.authService.isAuthenicated().subscribe()
  }

  ngOnDestroy(): void {
    if (this.userSubscription) this.userSubscription.unsubscribe();
  }
}
