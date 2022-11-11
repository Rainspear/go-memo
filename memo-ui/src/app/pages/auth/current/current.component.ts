import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { faCaretLeft, faCaretRight, faLock, faPen } from '@fortawesome/free-solid-svg-icons';
import { UpdateUser, User } from 'src/app/models/user.model';
import { ApiService } from 'src/app/services/api.service';
import { AuthUserService } from 'src/app/services/auth-user.service';

@Component({
  selector: 'app-current',
  templateUrl: './current.component.html',
  styleUrls: ['./current.component.scss']
})
export class CurrentComponent implements OnInit {
  faCaretLeft = faCaretLeft;
  faCaretRight = faCaretRight;
  faLock = faLock;
  faPen = faPen;

  user?: UpdateUser;

  constructor(private location: Location, private authService: AuthUserService, private apiService: ApiService) {
    // for enter to url directly, make sure has user data
    if (!this.user) {
      this.authService.loggedUser.subscribe(user => {
        this.user = { ...user };
        this.user.created_date = new Date(user.created_date * 1000)
      })

    }
    console.log("this.user", this.user)
  }

  ngOnInit(): void {
    if (this.authService.user) {
    this.user = this.authService.user;
      this.user.created_date = new Date(this.authService.user.created_date * 1000)
    }
    console.log("this.user", this.user)

  }

  onSave(): void {
    this.apiService.updateUser
  }

  onNavigate(): void {
    this.location.back();
  }

}
