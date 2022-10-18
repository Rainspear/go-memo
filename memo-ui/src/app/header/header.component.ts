import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { User } from '../models/user.model';
import { ApiService } from '../services/api.service';
import { AuthUserService } from '../services/auth-user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  collapsed: boolean = true;
  user?: User
  @Output() featureSelected = new EventEmitter<string>();
  constructor(private authService : AuthUserService, private apiService: ApiService) {
    this.authService.logged.subscribe(user => {
      this.user = user
    })
  }

  onClickSelectNav(route: string) {
    this.featureSelected.emit(route);
  }

  ngOnInit(): void {
    this.apiService.currentUser().subscribe((res: any) => {
      console.log("res", res)
      if (res.data) {
        console.log(res.data)
        this.authService.onGetUser(res.data)
        return;
      }
      return;
    }, error => {
      // this.error = error?.error?.error || error.message
    })
  }

}
